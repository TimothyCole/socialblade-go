package socialblade

import (
	"errors"
)

// GetFavourites gets a list of a users favourites (User auths only)
func (c *Client) GetFavourites() ([]Favourite, error) {
	if !c.user {
		return nil, errors.New("Only users can have Favourites")
	}

	var out []Favourite
	err := c.request("GET", "favorites/list", "v2", nil, nil, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Favourite gets a list of a users favourites (User auths only)
func (c *Client) Favourite(sec section, id string) error {
	return c.fav(sec, id, "add")
}

// Unfavourite gets a list of a users favourites (User auths only)
func (c *Client) Unfavourite(sec section, id string) error {
	return c.fav(sec, id, "remove")
}

// CheckFavourited gets a list of a users favourites (User auths only)
func (c *Client) CheckFavourited(sec section, id string) bool {
	err := c.fav(sec, id, "check")
	if err != nil {
		return false
	}
	return true
}

func (c *Client) fav(sec section, id, function string) error {
	if !c.user {
		return errors.New("Only users can have Favourites")
	}

	var body = struct {
		Section  string `url:"section"`
		Username string `url:"username"`
	}{Section: string(sec), Username: id}

	var out favouriteUpdate
	err := c.request("POST", "favorites/"+function, "v2", nil, body, &out)
	if err != nil {
		return err
	}

	if out.Favourite && function == "check" {
		return nil
	} else if !out.Favourite && function == "check" {
		return errors.New("Not in favourites")
	}

	switch r := out.Status.(type) {
	case bool:
		return nil
	case interface{}:
		return errors.New(r.(map[string]interface{})["error"].(string))
	default:
		return errors.New("unknown error")
	}
}
