package client

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ubozov/grpc-example/proto"

	"google.golang.org/grpc"
)

// Start starts the grpc client
func Start() error {
	conn, err := grpc.Dial(":4040", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not connect to backend: %v", err)
	}

	defer conn.Close()

	client := proto.NewAddServiceClient(conn)

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		queryValues := r.URL.Query()
		a, err := strconv.ParseUint(queryValues.Get("a"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error:" + fmt.Sprint(err.Error()) + "\n"))
		}

		b, err := strconv.ParseUint(queryValues.Get("b"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error:" + fmt.Sprint(err.Error()) + "\n"))
		}

		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := client.Add(context.Background(), req); err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("result:" + fmt.Sprint(response.Result) + "\n"))

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error:" + fmt.Sprint(err.Error()) + "\n"))
		}

	})

	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		queryValues := r.URL.Query()
		a, err := strconv.ParseUint(queryValues.Get("a"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error:" + fmt.Sprint(err.Error()) + "\n"))
		}

		b, err := strconv.ParseUint(queryValues.Get("b"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error:" + fmt.Sprint(err.Error()) + "\n"))
		}

		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := client.Multiply(context.Background(), req); err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("result:" + fmt.Sprint(response.Result) + "\n"))

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error:" + fmt.Sprint(err.Error()) + "\n"))
		}

	})

	http.ListenAndServe(":3001", nil)
	return nil
}
