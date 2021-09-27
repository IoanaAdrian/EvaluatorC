package main

import (
	"./functionalities"
	ut "./utilities"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	panic(http.ListenAndServe(":8080", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.String()) == 1 {
		page, err := ioutil.ReadFile("./static/index.html")
		ut.HandleErr(err)
		_, err = fmt.Fprintf(w, string(page))
		ut.HandleErr(err)
	} else if strings.HasPrefix(r.URL.String(), "/solution") {
		ut.HandleErr(r.ParseForm())
		name := r.FormValue("username")
		exercise := r.FormValue("exercise")
		source := r.FormValue("source")
		functionalities.Compile(name, exercise, source, w)
		/*fmt.Println(name)
		fmt.Println(exercise)
		fmt.Println(source)*/
	} else if strings.HasPrefix(r.URL.String(),"/history"){
		functionalities.DisplayHistory(w)
	}
}
