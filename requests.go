package socialblade

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

func (c *Client) request(method, uri, version string, q, b, r interface{}) error {
	var qs = ""
	if q != nil {
		q, err := query.Values(q)
		if err == nil {
			qs = "?" + q.Encode()
		}
	}

	var postBody = ""
	if b != nil {
		b, err := query.Values(b)
		if err == nil {
			postBody = b.Encode()
		}
	}

	req, err := http.NewRequest(method, "https://api.socialblade.com/"+version+"/"+uri+qs, bytes.NewBuffer([]byte(postBody)))
	if err != nil {
		return err
	}

	if c.email != "" && c.token != "" {
		req.Header.Add("Authorization", c.email+"::"+c.token)
	}

	if postBody != "" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
