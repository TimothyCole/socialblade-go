package socialblade

import "errors"

// StatsTwitch returns stats for a specific Twitch Channel
func (c *Client) StatsTwitch(identifier string) (*TwitchStatsData, error) {
	var out *TwitchStatsData

	var qs = struct {
		Query    string `url:"query"`
		Username string `url:"username"`
	}{Query: "statistics", Username: identifier}

	err := c.request("GET", "twitch/statistics", "v2", qs, nil, &out)
	if err != nil {
		return nil, err
	}

	switch r := out.Status.Error.(type) {
	case string:
		return nil, errors.New(r)
	case bool:
		if r {
			return nil, errors.New(out.Status.Message)
		}
		return out, nil
	default:
		return out, nil
	}
}
