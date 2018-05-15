package socialblade

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testUser string

func init() {
	key = os.Getenv("SOCIALBLADE_KEY")

	email = os.Getenv("SOCIALBLADE_EMAIL")
	token = os.Getenv("SOCIALBLADE_TOKEN")

	testUser = "socialblade"
}

func TestFavourites(t *testing.T) {
	// Key Test
	sb, err := Auth(key)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		_, err := sb.GetFavourites()
		assert.NotNil(t, err)
		t.Log(" Key Test:", err)
	}

	// User Test
	sb, err = AuthAsUser(email, token)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		favourites, err := sb.GetFavourites()
		assert.Nil(t, err)
		t.Log("User Test: Favourites", len(favourites))
	}
}

func TestFavourite(t *testing.T) {
	// Key Test
	sb, err := Auth(key)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		err := sb.Favourite(SectionTwitter, testUser)
		assert.NotNil(t, err)
		t.Log(" Key Test:", err)
	}

	// User Test
	sb, err = AuthAsUser(email, token)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		err := sb.Favourite(SectionTwitter, testUser)
		assert.Nil(t, err)
		t.Log("User Test: Favourited", testUser, "on", SectionTwitter)
	}
}

func TestCheckFavourited(t *testing.T) {
	// Key Test
	sb, err := Auth(key)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		ok := sb.CheckFavourited(SectionTwitter, testUser)
		assert.False(t, ok)
		t.Log(" Key Test:", ok)
	}

	// User Test
	sb, err = AuthAsUser(email, token)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		ok := sb.CheckFavourited(SectionTwitter, testUser)
		assert.True(t, ok)
		t.Log("User Test:", testUser, "is in our", SectionTwitter, "favourites list")
	}
}

func TestUnfavourite(t *testing.T) {
	// Key Test
	sb, err := Auth(key)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		err := sb.Unfavourite(SectionTwitter, testUser)
		assert.NotNil(t, err)
		t.Log(" Key Test:", err)
	}

	// User Test
	sb, err = AuthAsUser(email, token)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		err := sb.Unfavourite(SectionTwitter, testUser)
		assert.Nil(t, err)
		t.Log("User Test: Unfavourited", testUser, "on", SectionTwitter)
	}
}
