package Infrastructure

import (
	"os"
	"testing"
)

func TestClientFetchServiceURL(t *testing.T) {

	err := os.Setenv("CLIENTS_HOST", "0.0.0.0")
	if err != nil {
		t.Error("Can't set CLIENTS_HOST")
	}

	err = os.Setenv("CLIENTS_PORT", "8080")
	if err != nil {
		t.Error("Can't set CLIENTS_PORT")
	}

	err = os.Setenv("VERSION", "v1")
	if err != nil {
		t.Error("Can't set VERSION")
	}

	cfs := GenClientFetchService()
	url := cfs.GetURL()

	if url != "http://0.0.0.0:8080/v1/" {
		t.Error("Expected http://0.0.0.0:8080/v1/, got", url)
	}

}
