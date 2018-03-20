package main

import (
	"fmt"
	"os"
)

func main() {
	debug := os.Getenv("MYAPP_DEBUG")           // false
	port := os.Getenv("MYAPP_PORT")             // 8080
	user := os.Getenv("MYAPP_USER")             // Kelsey
	rate := os.Getenv("MYAPP_RATE")             // 0.5
	timeout := os.Getenv("MYAPP_TIMEOUT")       // 3m
	users := os.Getenv("MYAPP_USERS")           // rob,ken,robert
	colorcodes := os.Getenv("MYAPP_COLORCODES") // red:1,green:2,blue:3
	fmt.Printf("debug: %s\n", debug)
	fmt.Printf("port: %s\n", port)
	fmt.Printf("user: %s\n", user)
	fmt.Printf("rate: %s\n", rate)
	fmt.Printf("timeout: %s\n", timeout)
	fmt.Printf("users: %s\n", users)
	fmt.Printf("colorcodes: %s\n", colorcodes)
}
