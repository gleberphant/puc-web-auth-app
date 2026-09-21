package configs

import (
	"os"
)

var (
	AMBIENTE string = func() string {
		s := os.Getenv("AMBIENTE")
		if s == "" {
			return "dev"
		}
		return s
	}()

	PORTA string = func() string {
		s := os.Getenv("PORT")
		if s == "" {
			s = os.Getenv("PORTA")
		}
		if s == "" {
			return ":4000"
		}
		if s[0] != ':' {
			return ":" + s
		}
		return s
	}()

	CORS_ORIGIN string = func() string {
		origin := os.Getenv("CORS_ORIGIN")
		if origin == "" {
			return "http://localhost:5173"
		}
		return origin
	}()

	JWTSECRET string = func() string {
		secret := os.Getenv("JWTSECRET")

		if secret == "" {

			if AMBIENTE == "prod" {
				panic("secret do jwt nao definido")
			}

			return "minha-senha-secreta"

		}

		return secret
	}()
)
