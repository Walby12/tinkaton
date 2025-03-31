package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func start_app() {
  file := "password.txt"
  buffer := bufio.NewReader(os.Stdin)
  fmt.Println("Welcome to the tinkaton password manager")

  for {
    fmt.Print("USER >> ")
    command, err := buffer.ReadString('\n')
    check(err)
    command = strings.TrimSpace(command)
    switch command {
    case "help":
      usage()
    case "hello":
      fmt.Println("Hello :) glad you trust me!")
    case "newkey":
      fmt.Print("What is the app you would like to link the password to? ")
      command, err = buffer.ReadString('\n')
      check(err)
      app_name := strings.TrimSpace(command)
      err = os.WriteFile(file, []byte(app_name), 0644)
      check(err)
    default:
      fmt.Println("ERROR: unrecognized command:", command)
    }
  }
}

func usage() {
  fmt.Println("\nUSAGE")
  fmt.Println("help\tPrints this")
  fmt.Println("newkey\tCreates an new password\n")
}

func main() {
  start_app()
}
