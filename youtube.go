package socialblade

import "errors"

// StatsYouTube returns stats for a specific YouTube Channel
func (c *Client) StatsYouTube(identifier string) (*YouTubeStatsData, error) {
	var out *YouTubeStatsData

	var qs = struct {
		Query    string `url:"query"`
		Username string `url:"username"`
	}{Query: "statistics", Username: identifier}

	err := c.request("GET", "youtube/statistics", "v2", qs, nil, &out)
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
