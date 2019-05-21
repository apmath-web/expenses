package Infrastructure

import (
	"encoding/json"
	"errors"
	"github.com/apmath-web/expenses/Domain"
	"github.com/apmath-web/expenses/Infrastructure/Mapper"
	"github.com/apmath-web/expenses/Infrastructure/applicationModels"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type clientFetchService struct {
	url string
}

func (clfs *clientFetchService) GenURL() {
	host := os.Getenv("CLIENT_HOST")
	port := os.Getenv("CLIENT_PORT")
	version := os.Getenv("VERSION")
	clfs.url = "http://" + host + ":" + port + "/" + version + "/"
}

var instantiated *clientFetchService
var once sync.Once

func GenClientFetchService() Domain.ClientFetchInterface {
	once.Do(func() {
		instantiated = &clientFetchService{}
		instantiated.GenURL()
	})
	return instantiated
}

func (clfs *clientFetchService) Fetch(id int) (Domain.PersonDomainModelInterface, error) {
	resp, err := http.Get(clfs.url + strconv.Itoa(id))
	if resp == nil {
		return nil, errors.New("clients service not available")
	}
	if err != nil {
		return nil, err
	}
	person := new(applicationModels.PersonApplicationModel)

	if resp.StatusCode == http.StatusOK {
		dec := json.NewDecoder(resp.Body)
		err := dec.Decode(&person)
		if err != nil {
			return nil, err
		}
	}
	if resp.StatusCode == http.StatusBadRequest {
		return nil, errors.New("bad request")
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("client not found")
	}
	pdm := Mapper.PersonApplicationMapper(*person)
	return pdm, nil
}
