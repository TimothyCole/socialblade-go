package socialblade

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

func (c *Client) request(method, uri, version string, q interface{}, b []byte, r interface{}) error {
	var qs = ""
	if q != nil {
		query, err := query.Values(q)
		if err == nil {
			qs = "?" + query.Encode()
		}
	}

	req, err := http.NewRequest(method, "https://api.socialblade.com/"+version+"/"+uri+qs, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	if c.email != "" && c.token != "" {
		req.Header.Add("Authorization", c.email+"::"+c.token)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(body), r)
	if err != nil {
		return err
	}

	return nil
}
