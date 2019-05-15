package Infrastructure

import (
	"encoding/json"
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

func (u *clientFetchService) GenURL() {
	host := os.Getenv("CLIENT_HOST")
	port := os.Getenv("CLIENT_PORT")
	version := os.Getenv("VERSION")
	u.url = "http://" + host + ":" + port + "/" + version + "/"
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
	} else {
		return nil, err
	}
	var pdm = Mapper.PersonApplicationMapper(*person)
	return pdm, nil
}
