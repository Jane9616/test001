package main

import (
	"fmt"
	"net/http"
	"time"

	"test0812/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// 子路由
	s := r.PathPrefix("/users").Subrouter()
	s.HandleFunc("", handlers.GetUsers).Methods("GET")
	s.HandleFunc("/{id}", handlers.GetUser).Methods("GET")
	s.HandleFunc("", handlers.CreateUser).Methods("POST")
	s.HandleFunc("/{id}", handlers.UpdateUser).Methods("PUT")

	r.Use(middleware)

	// r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	// r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	// r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	// r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")



	fmt.Println("server is running on port 8800")
	// http.ListenAndServe(":8080", r)
	if err := http.ListenAndServe(":8800", r); err != nil {
		fmt.Println("Error running server", err)
	}
}

// 中间件
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 在处理请求之前记录日志
		start := time.Now()
		fmt.Printf("started %s %s\n", r.Method, r.URL.Path)

		// 调用下一个处理函数
		next.ServeHTTP(w, r)

		// 在处理请求之后记录日志
		fmt.Printf("completed %s %s in %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}
