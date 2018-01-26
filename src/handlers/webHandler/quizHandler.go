package webHandler

import (
	"fmt"
	"net/http"
	"net/url"
	"omt-project/config"
	"omt-project/src/services/quizService"
	"omt-project/src/services/userService"
	"omt-project/src/templateEngine"
	"path"
	"strconv"

	"github.com/labstack/echo"
)

func GetQuizHandlerGenerator(e *echo.Echo) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")
		// Check user exists
		_, err := userService.IsUserExist(userId)
		if err != nil {
			e.Logger.Errorf("User %s does not exist", userId)
			return c.Render(http.StatusNotFound, "notFound.html", "not Found")
		}

		// Make quiz
		qc, err := quizService.NewQuizController(userId)
		if err != nil {
			e.Logger.Errorf("Cannot make quiz for id: %s", userId)
			return c.Render(http.StatusNotFound, "notFound.html", "not Found")
		}
		// TODO: n could be changed by something
		n := 3
		words, err := qc.CreateMeaningWordQuiz(n)
		if err != nil {
			e.Logger.Errorf("%v\n", err)
			return c.Render(http.StatusNotFound, "notFound.html", "not Found")
		}

		// Make data to make dynamic template
		cfg := config.Setting()

		data := templateEngine.NewData()
		data.Add("userId", userId)
		data.Add("words", words)
		data.Add("questionNums", len(words))
		data.Add("appEnv", cfg.AppEnv)

		return c.Render(http.StatusOK, "quiz.html", data)
	}
}

func PostQuizHandlerGenerator(e *echo.Echo) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.FormValue("user-id")
		questionNums, err := strconv.Atoi(c.FormValue("question-nums"))
		if err != nil {
			e.Logger.Errorf("%v\n", err)
			return c.Render(http.StatusBadRequest, "notFound.html", "not Found")
		}
		userAns := make([]string, 0, 5)
		goodAns := make([]string, 0, 5)
		for i := 0; i < questionNums; i++ {
			ua := c.FormValue(fmt.Sprintf("user-ans-%d", i))
			ga := c.FormValue(fmt.Sprintf("good-ans-%d", i))
			userAns = append(userAns, ua)
			goodAns = append(goodAns, ga)
		}

		// Check correctness and apply results to awsS3
		// Change word priority
		qc, err := quizService.NewQuizController(userId)
		if err != nil {
			e.Logger.Errorf("Make new quiz controller error %+v", err)
			e.Logger.Errorf("%v\n", err)
			return c.Render(http.StatusBadGateway, "notFound.html", "not Found")
		}
		srs := quizService.GetScoreResults(userAns, goodAns)
		qc.ApplyMeaningWordQuizResult(srs)

		// Update Words on S3
		err = qc.UpdateWordsInfo(userId)
		if err != nil {
			e.Logger.Errorf("Update Words Info failed %+v", err)
			return c.Render(http.StatusBadGateway, "notFound.html", "not Found")
		}
		fmt.Printf("qc: %+v", qc.Words)

		// For redirect url quiz results
		redirectUrl, err := createQuizResultUrl(srs, userId)
		if err != nil {
			e.Logger.Errorf("Url setting is not valid: %+v", err)
			return c.Render(http.StatusBadGateway, "notFound.html", "not Found")
		}

		return c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
	}
}

func GetQuizResultHandlerGenerator(e *echo.Echo) echo.HandlerFunc {
	return func(c echo.Context) error {
		qp := c.QueryParams()
		userId := qp.Get("userId")
		words := qp["words"]
		correctness := qp["correctness"]
		n := len(words)

		boolCorrectness := make([]bool, 0, n)
		for _, cp := range correctness {
			if cp == "1" {
				boolCorrectness = append(boolCorrectness, true)
			} else {
				boolCorrectness = append(boolCorrectness, false)
			}
		}

		srs := make(quizService.ScoreResults, 0, n)
		correct := 0
		for i := 0; i < n; i++ {
			sr := &quizService.ScoreResult{
				Word:        words[i],
				Correctness: boolCorrectness[i],
			}
			srs = append(srs, sr)
			if boolCorrectness[i] {
				correct++
			}
		}

		data := templateEngine.NewData()
		data.Add("userId", userId)
		data.Add("srs", srs)
		data.Add("totalAnswer", n)
		data.Add("correctAnswer", correct)

		return c.Render(http.StatusOK, "quizResult.html", data)
	}
}

func createQuizResultUrl(srs quizService.ScoreResults, userId string) (string, error) {
	// add basic url
	cfg := config.Setting()
	baseUrl, err := url.Parse(cfg.Host)
	if err != nil {
		return "", err
	}
	baseUrl.Path = path.Join(baseUrl.Path, "/quiz/result")
	// add query parameter
	v := url.Values{}
	v.Set("userId", userId)
	for _, sr := range srs {
		v.Add("words", sr.Word)
		if sr.Correctness == true {
			v.Add("correctness", "1")
		} else {
			v.Add("correctness", "0")
		}
	}
	baseUrl.RawQuery = v.Encode()
	return baseUrl.String(), nil
}
