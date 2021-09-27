package functionalities

import (
	ut "../utilities"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func DisplayHistory(w http.ResponseWriter) {
	file, err := os.Open("./history.txt")
	ut.HandleErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		_, err = fmt.Fprintf(w, "<p style=\"font-size:2vw;color:#148F77 \">"+scanner.Text()+"</p>")
	}
	ut.HandleErr(err)
}

func AddToFile(s string) {
	file, err := os.OpenFile("./history.txt", os.O_APPEND|os.O_WRONLY, 0666)
	ut.HandleErr(err)
	if _, err := file.WriteString(s); err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	ut.HandleErr(err)
}
