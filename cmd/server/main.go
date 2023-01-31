package main

// todo : checl after redis validation to hash or un hash the file
//todo: hash and un hash File and const to connection info
import (
	"net/http"

	handeler "github.com/alijkdkar/AI-File-DB-Go-Core/internal/handlers"
)

func main() {
	defaultnet := http.DefaultServeMux

	defaultnet.HandleFunc("/hash", handeler.GetHashFile)
	defaultnet.HandleFunc("/unhash", handeler.GetUnHashFile)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
