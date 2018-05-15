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
	var qs = struct {
		Key string `url:"key"`
	}{Key: client.key}

	err := client.request("GET", "third-party", "v2", qs, nil, &out)
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
	var qs = struct {
		Email string `url:"email"`
		Token string `url:"token"`
	}{
		Email: client.email,
		Token: client.token,
	}

	err := client.request("GET", "bridge", "v2", qs, nil, &out)
	if err != nil {
		return nil, err
	}

	if out.Status.Error {
		return nil, errors.New(out.Status.Message)
	}

	return client, nil
}
