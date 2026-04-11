package main

import (
    "fmt"
    //"html/template"
    "net/http"
    "github.com/gorilla/sessions"
)

type User struct {
    Name string
    Age  int
}

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func main() {
    // Fix 1: Register AFTER function definitions
    fmt.Println("Server running on http://localhost:8080")
    http.HandleFunc("/", homeHandler)
    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    Name, _ := session.Values["Name"].(string)

	if 
    
  fmt.Fprintln(w, Name,"Welcome to the Home page")

}
