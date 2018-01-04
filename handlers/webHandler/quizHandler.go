package webHandler

import (
	"fmt"
	"net/http"
	"omt-project/services/quizService"
	"omt-project/services/userService"
	"omt-project/templateEngine"
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
		data := templateEngine.NewData()
		data.Add("userId", userId)
		data.Add("words", words)
		data.Add("questionNums", len(words))

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

		err = qc.UpdateWordsInfo(userId)
		if err != nil {
			e.Logger.Errorf("Update Words Info failed %+v", err)
			return c.Render(http.StatusBadGateway, "notFound.html", "not Found")
		}
		fmt.Printf("qc: %+v", qc.Words)

		// TODO: Push result to the user
		return nil
	}
}
