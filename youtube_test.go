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
