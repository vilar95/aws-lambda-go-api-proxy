package main

import (
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Pong")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "                      		\n" +
		"      @@@@  @@@@  @   		\n" +
		"      @     @  @  @   		\n" +
		"      @ @@  @  @  @   		\n" +
		"      @  @  @  @      		\n" +
		"      @@@@  @@@@  @   		\n" +
		"         ___          		\n" +
		"        ʕ◔ᴥ◔ʔ        		\n" +
		"       /|   |\\       		\n" +
		"      / |   | \\      		\n" +
		"        |   |         		\n" +
		"        |___|         		\n" +
		"       /     \\       		\n" +
		"      /       \\      		\n" +
		"   +==================+ 	\n" +
		"   | Hello, World! 🚀 | 	\n" +
		"   | Prepare-se para  | 	\n" +
		"   | o Go Way! 🛣️      | 	 \n" +
		"   +==================+  	 \n" +
		"                         ",)
	})

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}