package socialblade

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	key = os.Getenv("SOCIALBLADE_KEY")

	email = os.Getenv("SOCIALBLADE_EMAIL")
	token = os.Getenv("SOCIALBLADE_TOKEN")

	testUser = "socialblade"
}

func TestStatsYouTube(t *testing.T) {
	// Key Test
	sb, err := Auth(key)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		_, err := sb.StatsYouTube(testUser)
		assert.NotNil(t, err)
		t.Log(" Key Test:", err)
	}

	// User Test
	sb, err = AuthAsUser(email, token)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		stats, err := sb.StatsYouTube(testUser)
		assert.Nil(t, err)
		assert.Equal(t, stats.ID.ChannelID, "UCl6vWwMCjufI8OPtOInHf0g")
		t.Log("User Test:", stats.Data.DisplayName, "has", stats.Data.Subs, "subscribers")
	}
}

func TestTopYouTube(t *testing.T) {
	var clients []*Client

	// Auth as Third-Party
	sb, err := Auth(key)
	clients = append(clients, sb)

	// Auth as First-Party
	sb, err = AuthAsUser(email, token)
	clients = append(clients, sb)

	// Run tests on Auth types
	for _, sb := range clients {
		assert.Nil(t, err)
		if assert.NotNil(t, sb) {
			top, err := sb.TopYouTube("subscribers")
			assert.NotNil(t, top)
			assert.Nil(t, err)
			assert.Equal(t, len(top.Result) >= 100, true)
		}
	}
}
