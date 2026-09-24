package main

import (
	cmd "ecommerce/Cmd"
)

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello HandleFunc")
// }

//	func AboutHandler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "I am a software engineer")
//	}

func main() {
	cmd.Serve()
}
