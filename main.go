package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // defaults to port 8081 if no env port set
	}

	srv := http.Server{
		Handler:      m,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// this blocks forever, until the server
	// has an unrecoverable error
	fmt.Println("server started on ", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
</body>
</html>
`
	w.Write([]byte(page))
}
