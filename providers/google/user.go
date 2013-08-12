package google

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
)

const (
	googleKeyID        string = "id"
	googleKeyName      string = "name"
	googleKeyEmail     string = "email"
	googleKeyAvatarUrl string = "picture"
)

type User struct {
	data objects.Map
}

func NewUser(data objects.Map, creds *common.Credentials, provider common.Provider) *User {
	user := &User{data}

	creds.Set(common.CredentialsKeyID, data[googleKeyID])

	// set provider credentials
	user.data[common.UserKeyProviderCredentials] = map[string]*common.Credentials{
		provider.Name(): creds,
	}

	return user
}

// Email gets the users email address.
func (u *User) Email() string {
	return u.Data().GetStringOrEmpty(googleKeyEmail)
}

// Name gets the users full name.
func (u *User) Name() string {
	return u.Data().GetStringOrEmpty(googleKeyName)

}

// Nickname gets the users nickname or username.
func (u *User) Nickname() string {
	return ""

}

// AvatarURL gets the URL of an image representing the user.
func (u *User) AvatarURL() string {
	return u.Data().GetStringOrEmpty(googleKeyAvatarUrl)
}

// ProviderCredentials gets a map of Credentials (by provider name).
func (u *User) ProviderCredentials() map[string]*common.Credentials {
	return u.Data().Get(common.UserKeyProviderCredentials).(map[string]*common.Credentials)
}

// IDForProvider gets the ID value for the specified provider name for
// this user from the ProviderCredentials data.
func (u *User) IDForProvider(provider string) string {
	return u.ProviderCredentials()[provider].GetStringOrEmpty(common.UserKeyID)
}

// ID gets this user's globally unique ID.
func (u *User) ID() string {
	return u.Data().GetStringOrEmpty(common.UserKeyID)
}

// GetValue gets any User field by name.
func (u *User) Data() objects.Map {
	return u.data
}

func (u *User) PublicData(options map[string]interface{}) (publicData interface{}, err error) {
	return u.data, nil
}
