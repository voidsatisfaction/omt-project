package webHandler

import (
	"net/http"
	"omt-project/services/quizService"
	"omt-project/services/userService"
	"omt-project/templateEngine"

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
			e.Logger.Errorf("%v", err)
			return c.Render(http.StatusNotFound, "notFound.html", "not Found")
		}

		// Make data to make dynamic template
		data := templateEngine.NewData()
		data.Add("words", words)

		return c.Render(http.StatusOK, "quiz.html", data)
	}
}
