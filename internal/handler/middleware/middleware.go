package middleware

import (
	"context"
	"gilsaputro/dating-apps/internal/handler/utilhttp"
	"gilsaputro/dating-apps/internal/store/user"
	"gilsaputro/dating-apps/pkg/token"
	"net/http"
	"strings"
)

// Middleware struct is list dependecies to run Middleware func
type Middleware struct {
	tokenMethod token.TokenMethod
	userStore   user.UserStoreMethod
}

// NewMiddleware is func to create Middleware Struct
func NewMiddleware(tokenMethod token.TokenMethod, userStore user.UserStoreMethod) Middleware {
	return Middleware{
		tokenMethod: tokenMethod,
		userStore:   userStore,
	}
}

// RequestBody is struct for parameter middleware
type RequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// MiddlewareVerifyToken is func to validate before execute the handler
func (m *Middleware) MiddlewareVerifyToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header value from the request
		authHeader := r.Header.Get("Authorization")

		// Check if the Authorization header is empty or does not start with "Bearer "
		if (authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ")) && r.URL.Path != "/register" {
			data := []byte(`{"code":401,"message":"unauthorized"}`)
			utilhttp.WriteResponse(w, data, http.StatusUnauthorized)
			return
		}

		// Extract the token from the Authorization header
		auth := strings.TrimPrefix(authHeader, "Bearer ")

		tokenBody, err := m.tokenMethod.ValidateToken(auth)
		if err != nil {
			data := []byte(`{"code":401,"message":"unauthorized"}`)
			utilhttp.WriteResponse(w, data, http.StatusUnauthorized)
			return
		}

		// Parse variable into context
		r = r.WithContext(context.WithValue(r.Context(), "id", tokenBody.UserID))
		next.ServeHTTP(w, r)
	}
}

// MiddlewareCheckVerifiedStatus is func to validate before execute the handler
func (m *Middleware) MiddlewareCheckVerifiedStatus(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value("id").(int)
		if !ok {
			data := []byte(`{"code":500,"message":"Internal Server Error"}`)
			utilhttp.WriteResponse(w, data, http.StatusInternalServerError)
			return
		}

		userInfo, err := m.userStore.GetUserInfoByID(userID)
		if err != nil {
			data := []byte(`{"code":500,"message":"Internal Server Error"}`)
			utilhttp.WriteResponse(w, data, http.StatusInternalServerError)
			return
		}

		// Parse variable into context
		r = r.WithContext(context.WithValue(r.Context(), "isverified", userInfo.IsVerified))
		next.ServeHTTP(w, r)
	}
}
