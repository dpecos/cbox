package repository

import (
	"fmt"

	"github.com/dplabs/cbox/src/models"

	"github.com/dplabs/cbox/src/tools/tty"
	"github.com/spf13/viper"
)

const (
	cloudJWTDev = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3PqZEDzJ2E8le5aFs8Tw
Um0tcUrc+614d9fseI6pmVOTKcNWTgktNX9rTz/B4JTCws3/8erqMVkwuz1vhH6S
iY+BUyn24g44/szZtVc0RgZVqLnZ87nsWvL2C+M1L4AiIgAwyElOFY5MCuknXMxD
oYOmobEYbJry4+ZkraUUzCGgWIDoYK8j/JzG63mw6QUZPO9fKSgUPDyjh6NAOq2c
+4WcdG2Ss0mseYeUXjUW0S2IZTXkYcfqJyXjDgCNSUGUzA5+NwKyjl5Sijr55ULD
wUMqJYfjFEGN4HlIqA80PHpQqjiWtAekZNRbfu0yhUW1s1ZUQchw1R7LF28aq5Bo
8wIDAQAB
-----END PUBLIC KEY-----`

	cloudJWTProd = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwmN+CQ1iigKbVIIWkeXa
pvVPpbspqI6w5qLcjsGh17mVMJB0FCRbRbC0pg0/TqP3qVWFJz11oyIXv3iJSjLu
vngA+nDXGGIlHwNfWFcW9wuHJxL/a+KH3+ehW3L1waLDCPvHWGFWJCW1uEkDIFS4
Syk1HNh2S8+WqXbXDEfY3iwDmt9JnG6bjNRhyIb7KzjnuY9reo+4Ej41LotQkkFs
IEXZlqHkR0EMCucCUxrGTklGoe6Ao+ZUla6cZRfn5mT0Bf7RqFBxoohJGE9Chp4V
eY5mkphAEDjPR6abKNHUejl4wh83Stg9AW3hEI0xU52gg4tkPEKOHhq1qYO0Alfz
cQIDAQAB
-----END PUBLIC KEY-----`

	cloudServerURL = "https://api.%s.cbox.dplabs.io"
)

var (
	cloudSettingsUserID    string
	cloudSettingsUserLogin string
	cloudSettingsUserName  string
	cloudSettingsJWT       string
)

func (repo *Repository) LoadCloudSettings(env string) (string, string, string, string, string, string) {

	url := fmt.Sprintf(cloudServerURL, env)
	key := cloudJWTProd

	if env != "prod" {
		fmt.Println()
		fmt.Println(tty.ColorBgRed("  !!! You are using a DEV version of cbox !!!   "))
		fmt.Println()

		key = cloudJWTDev
	}

	if env != "" && env != "prod" {
		env = "_" + env
	}
	if env == "prod" {
		env = ""
	}
	cloudSettingsUserID = fmt.Sprintf("cloud%s.auth.user.id", env)
	cloudSettingsUserLogin = fmt.Sprintf("cloud%s.auth.user.login", env)
	cloudSettingsUserName = fmt.Sprintf("cloud%s.auth.user.name", env)
	cloudSettingsJWT = fmt.Sprintf("cloud%s.auth.jwt", env)

	if !viper.IsSet(cloudSettingsUserID) || !viper.IsSet(cloudSettingsUserLogin) || !viper.IsSet(cloudSettingsUserName) || !viper.IsSet(cloudSettingsJWT) {
		return url, key, "", "", "", ""
	}

	userID := viper.GetString(cloudSettingsUserID)
	userLogin := viper.GetString(cloudSettingsUserLogin)
	userName := viper.GetString(cloudSettingsUserName)
	jwt := viper.GetString(cloudSettingsJWT)

	return url, key, userID, userLogin, userName, jwt

}

func (repo *Repository) StoreCloudSettings(cloud *models.Cloud) {
	viper.Set(cloudSettingsUserID, cloud.UserID)
	viper.Set(cloudSettingsUserLogin, cloud.Login)
	viper.Set(cloudSettingsUserName, cloud.Name)
	viper.Set(cloudSettingsJWT, cloud.Token)
}

func (repo *Repository) DeleteCloudSettings() {
	viper.Set(cloudSettingsUserID, "")
	viper.Set(cloudSettingsUserLogin, "")
	viper.Set(cloudSettingsUserName, "")
	viper.Set(cloudSettingsJWT, "")
}
