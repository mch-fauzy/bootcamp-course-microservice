package middleware

import (
	"bootcamp-course-microservice/infras"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	HeaderAuthorization = "Authorization"
)

type Authentication struct {
	DB     *infras.Conn
	Secret []byte
}

func ProvideAuthentication(db *infras.Conn, secret []byte) *Authentication {
	return &Authentication{
		DB:     db,
		Secret: secret,
	}
}

func IsTeacherRole(tokenString string) (bool, error) {
	// Make an HTTP GET request to the bootcamp-auth microservice to validate the JWT
	authServiceURL := "http://localhost:8080/v1/validate-auth"
	req, err := http.NewRequest("GET", authServiceURL, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "Bearer "+tokenString)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var responseData map[string]string
		err := json.NewDecoder(resp.Body).Decode(&responseData)
		if err != nil {
			return false, err
		}

		role, ok := responseData["role"]
		if !ok {
			return false, nil
		}

		return role == "teacher", nil
	}

	return false, nil
}

func (a *Authentication) VerifyJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Parse and validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if token.Method.Alg() != jwt.SigningMethodHS256.Name {
				return nil, jwt.ErrSignatureInvalid
			}
			return a.Secret, nil
		})
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			http.Error(w, "Token invalid", http.StatusUnauthorized)
			return
		}

		// Extract user information from token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "No user information from JWT", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			http.Error(w, "JWT user_id is not found", http.StatusUnauthorized)
			return
		}

		username, ok := claims["username"].(string)
		if !ok {
			http.Error(w, "JWT usernam is not found", http.StatusUnauthorized)
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			http.Error(w, "JWT role is not found", http.StatusUnauthorized)
			return
		}

		// Check if the user's role is "teacher"
		isTeacher, err := IsTeacherRole(tokenString)
		if err != nil {
			http.Error(w, "Error validating role", http.StatusInternalServerError)
			return
		}

		if !isTeacher {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Add user information to the request context
		ctx := context.WithValue(r.Context(), "user_id", userID)
		ctx = context.WithValue(ctx, "username", username)
		ctx = context.WithValue(ctx, "role", role)
		ctx = context.WithValue(ctx, "token", tokenString)
		r = r.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
