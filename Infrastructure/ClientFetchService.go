package Infrastructure

import (
	"encoding/json"
	"github.com/apmath-web/expenses/Domain"
	"github.com/apmath-web/expenses/Domain/models"
	"github.com/apmath-web/expenses/Infrastructure/Mapper"
	"github.com/apmath-web/expenses/Infrastructure/applicationModels"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type url struct {
	url string
}

func (u *url) GenURL() {
	host := os.Getenv("CLIENT_HOST")
	port := os.Getenv("CLIENT_PORT")
	version := os.Getenv("VERSION")
	u.url = "http://" + host + ":" + port + "/" + version + "/"
}

var instantiated *url
var once sync.Once

func GetURL() *url {
	once.Do(func() {
		instantiated.GenURL()
		instantiated = &url{}
	})
	return instantiated
}

type clientFetchService struct{}

func (clfs *clientFetchService) Fetch(id int) (Domain.PersonDomainModelInterface, error) {
	resp, err := http.Get(GetURL().url + strconv.Itoa(id))
	if err != nil {
		var pdm = new(models.PersonDomainModel)
		return pdm, err
	}
	person := new(applicationModels.PersonApplicationModel)
	if resp.StatusCode == http.StatusOK {
		dec := json.NewDecoder(resp.Body)
		for {
			if err := dec.Decode(&person); err == io.EOF {
				break
			} else if err != nil {
				var pdm = new(models.PersonDomainModel)
				return pdm, err
			}
		}
	} else {
		var pdm = new(models.PersonDomainModel)
		return nil, err
	}
	var pdm = Mapper.PersonApplicationMapper(*person)
	return pdm, nil
}
