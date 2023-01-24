package main

// todo : checl after redis validation to hash or un hash the file
//todo: hash and un hash File and const to connection info
import (
	"net/http"

	"github.com/alijkdkar/AI-File-DB-Go-Core/pkg/utility"
)

func main() {
	defaultnet := http.DefaultServeMux

	defaultnet.HandleFunc("/ping", hashFile)

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func hashFile(w http.ResponseWriter, r *http.Request) {

	fileID := r.URL.Query().Get("fileID")
	clint, err := utility.GetInstance()

	if err != nil {
		panic("redis is down")
	}
	clint.Get(fileID)

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{'noting':'nothung'"))
		w.WriteHeader(http.StatusAccepted)

	} else if r.Method == "Post" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{'noting':'nothung'"))
		w.WriteHeader(http.StatusAccepted)
	}

}
