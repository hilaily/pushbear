package pushbear

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BEAR_URL = "https://pushbear.ftqq.com/sub"
)

func httpGet(params string) ([]byte, error) {
	url := fmt.Sprintf("%s?%s", BEAR_URL, params)
	var res []byte
	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	return body, nil
}
