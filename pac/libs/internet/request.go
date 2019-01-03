package internet

import (
	pac "go-learning-codebin/pac/data"
	"go-learning-codebin/pac/libs/json"
	"io/ioutil"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

//HistoryFromURL .
func HistoryFromURL(url string) (pac.History, error) {

	r, err := myClient.Get(url)

	if err != nil {
		return pac.History{}, err
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return pac.History{}, err
	}

	return json.HistoryFromJSON(body)

}
