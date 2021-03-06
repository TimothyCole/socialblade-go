package socialblade

// TopYouTube get the top YouTube Users
func (c *Client) TopYouTube(filter string) (*YouTubeTopList, error) {
	var qs = struct {
		Query  string `url:"query"`
		Filter string `url:"top_type"`
	}{Query: "top", Filter: filter}

	var out *YouTubeTopList
	err := c.request("GET", "youtube/statistics", "v2", qs, nil, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
