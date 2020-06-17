package main

import (
	"fmt"
	"net"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
	fmt.Fprintf(w, "And I am ..... Ironman\n")
	ipAddress, _, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Fprintf(w, "IP address : %s", ipAddress)
}

func main() {
	fmt.Println("Server is running ....")
	fmt.Println("Press Ctrl+C to close the server. ")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
	//srv := &http.Server{Addr: ":8080"}
	//go func() {
	//	srv.ListenAndServe()
	//}()
	//time.Sleep(10 * time.Second)
	//srv.Shutdown(context.TODO())
	//s.Shutdown(context.Background())
}
