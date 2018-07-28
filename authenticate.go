package socialblade

import "errors"

// Auth allows you to authenticate with a third-party api key
func Auth(key string) (*Client, error) {
	var client = &Client{
		user:  false,
		email: "",
		token: "",
		key:   key,
	}

	var out tpAuth
	err := client.request("GET", "third-party", "v2", nil, nil, &out)
	if err != nil {
		return nil, err
	}

	if out.Status.Response != 200 {
		return nil, errors.New(out.Status.Error)
	}

	return client, nil
}

// AuthAsUser allows you to authenticate with an email and account token
func AuthAsUser(email, token string) (*Client, error) {
	var client = &Client{
		user:  true,
		email: email,
		token: token,
		key:   "",
	}

	var out userAuth
	err := client.request("GET", "bridge", "v2", nil, nil, &out)
	if err != nil {
		return nil, err
	}

	if out.Status.Error {
		return nil, errors.New(out.Status.Message)
	}

	return client, nil
}

// AuthAsAnon allows you to create a client without needing creds to access unlocked endpoints
func AuthAsAnon() *Client {
	var client = &Client{
		user:  true,
		email: "",
		token: "",
		key:   "",
	}

	return client
}
