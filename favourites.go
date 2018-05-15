package socialblade

import "errors"

// Favourites gets a list of a users favourites (User auths only)
func (c *Client) Favourites() ([]Favourite, error) {
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
