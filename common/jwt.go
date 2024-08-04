package common

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dongsu8142/blog-common/database"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"

type UserStore interface {
	GetUserByID(id int) (*database.User, error)
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := GetTokenFromRequest(r)

		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert userID to int: %v", err)
			permissionDenied(w)
			return
		}

		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(3600*24*7)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	return token.SignedString(secret)
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(EnvString("SECRET", "q1w2e3r4")), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	WriteError(w, http.StatusForbidden, "permission denied")
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}
