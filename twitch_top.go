package socialblade

// TopTwitch get the top YouTube Users
func (c *Client) TopTwitch(filter string) (*TwitchTopList, error) {
	var qs = struct {
		Query  string `url:"query"`
		Filter string `url:"top_type"`
	}{Query: "top", Filter: filter}

	var out *TwitchTopList
	err := c.request("GET", "twitch/statistics", "v2", qs, nil, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
