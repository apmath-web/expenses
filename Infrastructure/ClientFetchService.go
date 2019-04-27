package Infrastructure

import (
	"github.com/apmath-web/expenses/Infrastructure/applicationModels"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type clientFetchService struct{}

func (clfs *clientFetchService) Fetch(id int) *applicationModels.PersonApplicationModel {
	resp, err := http.Get("http://0.0.0.0:8080/v1/:" + strconv.Itoa(id))
	if err != nil {
		return nil
	}

	person := new(applicationModels.PersonApplicationModel)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		person.UnmarshalJSON(bodyBytes)
	}
	return person
}
