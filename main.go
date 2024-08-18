package main

import (
	"log"
	"net/http"
	queue "github.com/mubashirmalik/sqs-go/queue"
)

var sqlQueue queue.Queue

// Define a home handler function which writes a byte slice containing "Hello from SQS-GO Server!" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	sqlQueue.Print()
	w.Write([]byte("Hello from SQS-GO Server!"))
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	// Send a message to the SQS queue

}

func receiveMessage(w http.ResponseWriter, r *http.Request) {
	// Receive a message from the SQS queue
}

func main() {
	mux := http.NewServeMux()

	sqlQueue = queue.CreateQueue()
	
	mux.HandleFunc("/", home)
	mux.HandleFunc("/send-message", sendMessage)
	mux.HandleFunc("/receive-message", receiveMessage)

	log.Println(("Starting server on :4000"))

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
