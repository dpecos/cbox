package core

import (
	"fmt"

	"github.com/dpecos/cbox/tools"
)

var (
	Env     = "dev"
	Version = "development"
	Build   = "-"
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

	cloudURL = "https://api.%s.cbox.dplabs.io"
)

func CloudURL() string {
	return fmt.Sprintf(cloudURL, Env)
}

func init() {
	if Env == "dev" {
		tools.CloudJWTKey = cloudJWTDev
	} else {
		tools.CloudJWTKey = cloudJWTProd
	}
}
