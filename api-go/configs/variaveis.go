package configs

import (
	"os"
)

var (
	AMBIENTE string = func() string {
		if s := os.Getenv("AMBIENTE"); s != "" {
			return s
		}
		return "dev"
	}()

	PORTA string = func() string {
		if s := os.Getenv("PORTA"); s != "" {
			return s
		}
		return ":4000"
	}()

	JWTSECRET string = func() string {
		if secret := os.Getenv("JWTSECRET"); secret != "" {
			return secret
		}
		if AMBIENTE == "prod" {
			panic("secret do jwt nao definido")
		}

		return "minha-senha-secreta"
	}()
)
