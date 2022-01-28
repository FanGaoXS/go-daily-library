package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// LoggerMiddle 日志中间件
func LoggerMiddle(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path:%s process start...\n", r.URL.Path)
		defer func() {
			log.Printf("path:%s process end...\n", r.URL.Path)
		}()
		handler.ServeHTTP(w, r)
	})
}

// TimeMiddle 耗时中间件
func TimeMiddle(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Printf("path:%s time waste: %fs\n", r.URL.Path, time.Since(start).Seconds())
		}()
		time.Sleep(1 * time.Second)
		handler.ServeHTTP(w, r)
	})
}

// PanicMiddle 处理panic中间件
func PanicMiddle(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// 利用recover()捕获panic
			if err := recover(); err != nil {
				log.Println(string(debug.Stack()))
			}
		}()
		handler.ServeHTTP(w, r)
	})
}
