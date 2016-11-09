// test project main.go
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := r.URL.EscapedPath()[1:] // Все элементы массива, начиная с индекса 1
		url, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "http://"+string(url), http.StatusFound) // 302, трехсотые коды ответа
	})
	http.ListenAndServe(":8085", nil)

}
