package main

import (
    "errors"
    "fmt"
    "log"
    "os"
)

func getUser() (s string, e error) {
    username := os.Getenv("USER")
    if username == "" {
        err := errors.New("unable to set username")
        return "", err
    }
    return username, nil
}

func checkError(s string, err error) {
    if err != nil {
        log.Fatalf("at %s: %v", s, err)
        os.Exit(1)
    }
}

func main() {
    username, err := getUser()
    checkError("setting username", err)

    fmt.Printf("Hello %s, I hope that you're having a great day\n", username)
}
