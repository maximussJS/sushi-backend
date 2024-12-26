package router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"runtime"
	"sushi-backend/utils"
	"time"
)

func (router *Router) limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := router.ipRateLimiter.GetLimiter(utils.GetClientIpFromContext(r.Context()))

		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (router *Router) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		resp := router.authService.Verify(utils.GetClientIpFromContext(r.Context()), token)

		if resp.Status != http.StatusOK {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(resp.Status)
			utils.PanicIfError(json.NewEncoder(w).Encode(&resp.Data))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (router *Router) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		clientIP := utils.GetClientIpFromContext(r.Context())

		router.logger.Log(fmt.Sprintf(
			"[%s] %s %s %s %.0fm%.0fs%dms%dns %s",
			start.Format(time.RFC3339),
			r.Method,
			r.RequestURI,
			r.Proto,
			duration.Minutes(),
			duration.Seconds(),
			duration.Milliseconds(),
			duration.Microseconds(),
			clientIP,
		))
	})
}

func (router *Router) iPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
			clientIP = ip
		}

		next.ServeHTTP(w, r.WithContext(utils.GetContextWithClientIp(r.Context(), clientIP)))
	})
}

func (router *Router) Recover(next http.Handler) http.Handler {
	stackSize := router.config.ErrorStackTraceSizeInKb() << 10
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := make([]byte, stackSize)
				length := runtime.Stack(stack, true)
				stack = stack[:length]
				router.logger.Error(fmt.Sprintf("[PANIC RECOVER] %v %s\n", err, stack[:length]))

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "Internal Server error"}`))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (router *Router) setCacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		next.ServeHTTP(w, r)
	})
}

func (router *Router) noCache(handler http.HandlerFunc) http.HandlerFunc {
	return router.setCacheControl(handler).ServeHTTP
}

func (router *Router) isAdmin(handler http.HandlerFunc) http.HandlerFunc {
	return router.authMiddleware(handler).ServeHTTP
}

func (router *Router) addDefaultMiddlewares(r *mux.Router) {
	r.Use(router.Recover)
	r.Use(router.iPMiddleware)
	r.Use(router.limitMiddleware)
	r.Use(router.logMiddleware)
}
