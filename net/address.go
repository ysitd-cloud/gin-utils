package net

import (
	"fmt"

	"github.com/ysitd-cloud/gin-utils/env"
)

func GetAddress() string {
	port := env.GetEnvWithDefault("LISTEN_PORT", "8080")
	address := env.GetEnvWithDefault("LISTEN_ADDRESS", "")

	return fmt.Sprintf("%s:%s", address, port)
}
