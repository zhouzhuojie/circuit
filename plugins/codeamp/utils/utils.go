package utils

import (
	"context"
	"errors"
	"net/http"

	"github.com/codeamp/transistor"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaims struct {
	UserId      string   `json:"userId"`
	Permissions []string `json:"permissions"`
	TokenError  string   `json:"tokenError"`
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtClaims := JWTClaims{}
		authString := r.Header.Get("Authorization")

		if len(authString) < 8 {
			jwtClaims.TokenError = "invalid access token"
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "jwt", jwtClaims)))
			return
		}

		tokenString := authString[7:len(authString)]
		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("plugins.codeamp.jwt_secret")), nil
		})
		if err != nil {
			jwtClaims.TokenError = err.Error()
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "jwt", jwtClaims)))
			return
		}

		if token.Valid {
			if claims, ok := token.Claims.(*JWTClaims); ok {
				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "jwt", *claims)))
				return
			} else {
				jwtClaims.TokenError = "an unexpected error has occurred"
			}
		} else {
			jwtClaims.TokenError = "invalid access token"
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "jwt", jwtClaims)))
	})
}

func CheckAuth(ctx context.Context, scopes []string) error {
	jwtClaims := ctx.Value("jwt").(JWTClaims)

	if jwtClaims.UserId == "" {
		return errors.New(jwtClaims.TokenError)
	}

	if len(scopes) == 0 {
		return nil
	} else {
		for _, scope := range scopes {
			if transistor.SliceContains(scope, jwtClaims.Permissions) {
				return nil
			}
		}
		return errors.New("you dont have permission to access this resource")
	}

	return nil
}
