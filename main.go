package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

var users = map[string]*User{}

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

// 4
func jsonContentTypeMiddleware(next http.Handler) http.Handler {

	// 들어오는 요청의 Response Header에 Content-Type을 추가
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		// 전달 받은 http.Handler를 호출한다.
		next.ServeHTTP(rw, r)
	})
}

func main() {

	// NewRelic
	// app, err := newrelic.NewApplication(
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("GO_API_SAMPLE"),
		newrelic.ConfigLicense(""),
	)
	if err != nil {
		log.Fatal(err)
	}

	// func() {
	// 	txn := app.StartTransaction("myTask", nil, nil)
	// 	defer txn.End()

	// 	time.Sleep(time.Second)
	// }()

	// 1
	mux := http.NewServeMux()

	// 2
	userHandler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {

		txn := newrelic.FromContext(r.Context())
		txn.NoticeError(errors.New("my error message"))

		switch r.Method {
		case http.MethodGet: // 조회
			json.NewEncoder(wr).Encode(users)
		case http.MethodPost: // 등록
			var user User
			json.NewDecoder(r.Body).Decode(&user)

			users[user.Email] = &user

			json.NewEncoder(wr).Encode(user)
		}
	})

	// func myHandler(w http.ResponseWriter, r *http.Request) {
	// 	txn := newrelic.FromContext(r.Context())
	// 	txn.NoticeError(errors.New("my error message"))
	// }

	mux.HandleFunc(newrelic.WrapHandleFunc(app, "/", userHandler))

	// 3
	mux.Handle("/users", jsonContentTypeMiddleware(userHandler))
	http.ListenAndServe(":9000", mux)
}
