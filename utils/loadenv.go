package utils

import (
	"os"
	"strings"
)

func LoadEnv(env string) {
	s := strings.Split(env, "\r\n")

	for i, v := range s {
		if i < 7 {
			vS := strings.Split(v, "=")

			os.Setenv(vS[0], vS[1])
		}

	}
}
