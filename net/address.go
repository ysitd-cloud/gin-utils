package net // import "code.ysitd.cloud/gin/utils/net"

import (
	"fmt"

	"code.ysitd.cloud/gin/utils/env"
)

func GetAddress() string {
	port := env.GetEnvWithDefault("LISTEN_PORT", "8080")
	address := env.GetEnvWithDefault("LISTEN_ADDRESS", "")

	return fmt.Sprintf("%s:%s", address, port)
}
