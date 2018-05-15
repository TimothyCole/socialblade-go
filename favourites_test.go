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
}

func TestFavourites(t *testing.T) {
	// Key Test
	sb, err := Auth(key)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		_, err := sb.Favourites()
		assert.NotNil(t, err)
	}

	// User Test
	sb, err = AuthAsUser(email, token)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		favourites, err := sb.Favourites()
		assert.Nil(t, err)
		t.Log("Favourites", len(favourites))
	}
}
