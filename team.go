package socialblade

// Team gets a list of the Social Blade Team
func (c *Client) Team() (*TeamData, error) {
	var out *TeamData

	err := c.request("GET", "team", "v2", nil, nil, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
