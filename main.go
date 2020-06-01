package main

import (
    "log"
    "net/http"
    "os"
    "strconv"
)

func redirect(w http.ResponseWriter, r *http.Request) {
    status_code, err := strconv.Atoi(os.Getenv("STATUS_CODE"))
    if err != nil {
          log.Fatal(err)
    }
    http.Redirect(w, r, os.Getenv("REDIRECT_URL"), status_code)
}

func main() {
    http.HandleFunc("/", redirect)
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

