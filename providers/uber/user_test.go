package uber

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUserInterface(t *testing.T) {

	var user common.User = new(User)

	assert.NotNil(t, user)

}

func TestNewUser(t *testing.T) {

	testProvider := new(test.TestProvider)
	testProvider.On("Name").Return("providerName")

	data := objx.MSI(
		uberKeyID, "123435467890",
		uberKeyName, "Raquel",
		uberKeyNickname, "r@raquel.nyc",
		uberKeyAvatarUrl, "http://uberpic.com/")
	creds := &common.Credentials{Map: objx.MSI(oauth2.OAuth2KeyAccessToken, "ABC12345")}

	user := NewUser(data, creds, testProvider)

	if assert.NotNil(t, user) {

		assert.Equal(t, data, user.Data())

		assert.Equal(t, "Raquel", user.Name())
		assert.Equal(t, "r@raquel.nyc", user.Nickname())
		assert.Equal(t, "http://uberpic.com/", user.AvatarURL())

		// check provider credentials
		creds := user.ProviderCredentials()[testProvider.Name()]
		if assert.NotNil(t, creds) {
			assert.Equal(t, "ABC12345", creds.Get(oauth2.OAuth2KeyAccessToken).Str())
			assert.Equal(t, "123435467890", creds.Get(common.CredentialsKeyID).Str())
		}

	}

	mock.AssertExpectationsForObjects(t, testProvider.Mock)

}

func TestIDForProvider(t *testing.T) {

	user := new(User)
	user.data = objx.MSI(
		common.UserKeyProviderCredentials,
		map[string]*common.Credentials{
			"uber":   &common.Credentials{Map: objx.MSI(common.CredentialsKeyID, "uberid")},
			"google": &common.Credentials{Map: objx.MSI(common.CredentialsKeyID, "googleid")}})

	assert.Equal(t, "uberid", user.IDForProvider("uber"))
	assert.Equal(t, "googleid", user.IDForProvider("google"))

}
