package botService

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ActionResult struct {
	Text string
}

// TreatAction is manager of actions
func TreatAction(a *Action) *ActionResult {
	ar := &ActionResult{}
	switch a.actionType {
	case Invalid:
		InvalidActionTreat(a)
	case Search:
		// Call word search api
		resp, err := http.Get("https://glosbe.com/gapi/translate?from=eng&dest=kor&format=json&pretty=true&phrase=dragon")
		if err != nil {
			ar.ServerError()
			return ar
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("%+v", string(body))
		ar.Text = "search"
	}
	return ar
}

func InvalidActionTreat(a *Action) {

}

func (ar *ActionResult) ServerError() {
	ar.Text = "先生、今、体の調子が悪いの"
}
