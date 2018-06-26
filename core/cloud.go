package core

import (
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/viper"
)

func CloudLogin(jwt string) (string, error) {
	user, err := tools.VerifyJWT(jwt)
	if err != nil {
		return "", err
	}
	viper.Set("cloud.auth.user", user)
	viper.Set("cloud.auth.jwt", jwt)
	return user, nil
}
