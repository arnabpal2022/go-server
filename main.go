package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request	){

	// This code snippet is checking if there is an error while parsing the form data from the HTTP request.
	if err:= r.ParseForm(); err!= nil {
		fmt.Println(w, "ParseForm() err: %v", err)
		return;
	}

	fmt.Fprintf(w, "POST Request Successful");
	name := r.FormValue("name");
	address := r.FormValue("address");
	fmt.Fprintf(w, "Name=%s\n", name);
	fmt.Fprintf(w, "Address=%s\n", address);
}

func helloHandler(w http.ResponseWriter, r *http.Request){

	// This code snippet is checking if the requested URL path is not equal to "/hello". If the condition
	// is true, it returns a 404 Not Found error response to the client using `http.Error` function with
	// the message "404 not found" and the HTTP status code `http.StatusNotFound`. This is a way to handle
	// requests to paths other than "/hello" and inform the client that the requested resource is not
	// found on the server.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound);
		return;
	}

	// The code snippet `if r.Method!= "GET"{ http.Error(w, "Method is not Supported",
	// http.StatusNotFound) return }` is checking if the HTTP request method is not a GET method. If the
	// condition is true, it means that the client is using a different HTTP method other than GET to
	// access the "/hello" path.
	if r.Method!= "GET"{
		http.Error(w, "Method is not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!!")
}

func main(){
	
	
	//This Code Snippet is creating a File Server Handler from Static Folder which shows any static files in that directory
	fileServer := http.FileServer(http.Dir("./static"));

	http.Handle("/", fileServer);
	http.HandleFunc("/form", formHandler);
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Server in Port 8080\n");
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
