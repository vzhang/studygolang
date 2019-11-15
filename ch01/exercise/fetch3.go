// Modify fetch to also print the HTTP status code, found in resp.Status.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		var url string = arg
		if !strings.HasPrefix(arg, "http://") {
			url = "http://" + arg
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d\n", resp.StatusCode)
		fmt.Printf("%s", string((b)))
	}
}
