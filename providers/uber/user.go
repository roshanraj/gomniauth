package uber

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
	"strconv"
)

const (
	uberKeyID        string = "uuid"
	uberKeyName      string = "first_name"
	uberKeyNickname  string = "email"
	uberKeyEmail     string = "email"
	uberKeyAvatarUrl string = "picture"
)

type User struct {
	data objx.Map
}

// NewUser builds a new User object for Uber.
func NewUser(data objx.Map, creds *common.Credentials, provider common.Provider) *User {
	user := &User{data}

	creds.Set(common.CredentialsKeyID, data[uberKeyID])
	// set provider credentials
	user.data[common.UserKeyProviderCredentials] = map[string]*common.Credentials{
		provider.Name(): creds,
	}

	return user
}

// Name gets the users full name.
func (u *User) Name() string {
	return u.Data().Get(uberKeyName).Str()

}

// Nickname gets the users nickname or username.
func (u *User) Nickname() string {
	return u.Data().Get(uberKeyNickname).Str()

}

// Uber API doesn't return email
func (u *User) Email() string {
	return ""
}

// AvatarURL gets the URL of an image representing the user.
func (u *User) AvatarURL() string {
	return u.Data().Get(uberKeyAvatarUrl).Str()
}

// ProviderCredentials gets a map of Credentials (by provider name).
func (u *User) ProviderCredentials() map[string]*common.Credentials {
	return u.Data().Get(common.UserKeyProviderCredentials).Data().(map[string]*common.Credentials)
}

// IDForProvider gets the ID value for the specified provider name for
// this user from the ProviderCredentials data.
func (u *User) IDForProvider(provider string) string {
	id := u.ProviderCredentials()[provider].Get(common.CredentialsKeyID).Data()
	switch id.(type) {
	case string:
		return id.(string)
	case float64:
		return strconv.FormatFloat(id.(float64), 'f', 0, 64)
	}
	return ""
}

// AuthCode gets this user's globally unique ID (generated by the host program)
func (u *User) AuthCode() string {
	return u.Data().Get(common.UserKeyAuthCode).Str()
}

// GetValue gets any User field by name.
func (u *User) Data() objx.Map {
	return u.data
}

func (u *User) PublicData(options map[string]interface{}) (publicData interface{}, err error) {
	return u.data, nil
}
