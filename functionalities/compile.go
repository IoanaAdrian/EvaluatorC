package functionalities

import (
	ut "../utilities"
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func Compile(name string, exercise string, source string, w http.ResponseWriter) {
	path := "E:/ADI/EvaluatorC/sources/" + name + "-" + exercise
	err := os.WriteFile(path+".cpp", []byte(source), 0666)
	ut.HandleErr(err)
	cmd := exec.Command("g++", path+".cpp", "-o", path+".exe")
	//cmd := exec.Command("g++" , "C:\\Users\\Adrian\\Desktop\\test.cpp", "-o","C:\\Users\\Adrian\\Desktop\\a.exe")
	out, err := cmd.CombinedOutput()
	if err != nil {
		_, err = fmt.Fprintf(w, "<h1 style=\"font-size:2vw;color:red\">Compilation error")
	} else {
		testsPath := "E:/ADI/EvaluatorC/tests/" + exercise
		files, _ := ioutil.ReadDir(testsPath)
		correctTestNumber := 0

		//cmd = exec.Command("more",testsPath+"/1.in")

		for i := 1; i <= len(files)/2; i++ {
			//fmt.Println("cmd", "/c", "'"+path+".exe", "<", testsPath+"/"+strconv.Itoa(i)+".in"+"'")
			//cmd = exec.Command("cmd", "/c", "'",path+".exe", "<", testsPath+"/"+strconv.Itoa(i)+".in","'")
			cmd = exec.Command(path + ".exe")
			file, _ := os.Open(testsPath + "/" + strconv.Itoa(i) + ".in")
			cmd.Stdin = bufio.NewReader(file)
			out, err = cmd.CombinedOutput()
			ut.HandleErr(err)
			outFile, err := os.Open(testsPath + "/" + strconv.Itoa(i) + ".out")
			scanner := bufio.NewScanner(outFile)
			scanner.Split(bufio.ScanRunes)
			var correctOutput = ""
			for scanner.Scan() {
				correctOutput += scanner.Text()

			}
			ut.HandleErr(err)
			n := len(correctOutput)
			m := len(out) - 1
			ok := 0
			if n == m {
				for i := 0; i < n; i++ {
					if correctOutput[i] != out[i] {
						ok = 1
					}
				}
			} else {
				ok = 1
			}
			if ok == 0 {
				correctTestNumber++
				_, err = fmt.Fprintf(w, "<p style=\"font-size:2vw;color:green\">Testul "+strconv.Itoa(i)+" a trecut.</p>")
			} else {
				_, err = fmt.Fprintf(w, "<p style=\"font-size:2vw;color:red\">Testul "+strconv.Itoa(i)+" nu a trecut.</p>")
			}

			//fmt.Print('\n')

			//fmt.Println(bytes.Equal(out,correctOutput))
			/*if strings.TrimRight(string(correctOutput),"\n") == strings.TrimRight(string(out),"\n") {
				correctTestNumber++
				fmt.Println("muie steaua")
			}*/
			//fmt.Println(string(correctOutput))
			fmt.Println(string(out))
			fmt.Println()
		}
		points := float64(float64(correctTestNumber)/float64(len(files)/2)) * 100.0
		_, err = fmt.Fprintf(w, "<p style=\"font-size:2vw;color:blue\">Punctaj final "+strconv.Itoa(int(points))+" puncte. </p>")
		//fmt.Println(float64(float64(correctTestNumber) / float64(len(files) / 2)) * 100.0)
		err = os.Remove(path + ".exe")
		ut.HandleErr(err)
		err = os.Rename(path+".cpp", "E:/ADI/EvaluatorC/sources/"+strconv.Itoa(int(points))+name+"-"+exercise+".cpp")
		AddToFile(name + " a obtinut " + strconv.Itoa(int(points)) + " puncte la problema " + exercise + "\n")
	}
}
