package main

import (
	ascii "./AsciiArt"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

type Page struct {
	In  string
	Out string
}
//type Data struct {
//	Output    string
//	ErrorNum  int
//	ErrorText string
//}
func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	t, _ := template.ParseFiles("templates/500.html")
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//d := Data{}
	if r.URL.Path == "/" {
		switch r.Method {
		case "GET":
			t, err := template.ParseFiles("templates/index.html")
			if err != nil {
				internalServerError(w, r)
			}
			t.Execute(w, nil)
		case "POST":
			r.ParseForm()
			if !ValidAscii(r.Form.Get("input")) {
				w.WriteHeader(http.StatusBadRequest)
				t, err := template.ParseFiles("templates/400.html")
				if err != nil {
					internalServerError(w, r)
				}
				t.Execute(w, nil)
			} else {
				output, status := ascii.AsciiOutput(r.Form["input"][0], r.Form["font"][0])
				if status == 500 {
					internalServerError(w, r)
					//d.Output = output
					//if r.FormValue("process") == "download" {
					//	a := strings.NewReader(d.Output)
					//	w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
					//	w.Header().Set("Content-Length", strconv.Itoa(len(d.Output)))
					//	io.Copy(w, a)
					//}
				}else {
					ex := Page{
						In:  r.Form["input"][0],
						Out: output,
					}
					t, err := template.ParseFiles("templates/index.html")
					if err != nil {
						internalServerError(w, r)
						return
					}
					t.Execute(w, ex)
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		t, err := template.ParseFiles("templates/404.html")
		if err != nil {
			internalServerError(w, r)
			return
		}
		t.Execute(w, nil)
	}
}
func ValidAscii(s string) bool {
	for _, i := range []byte(s) {
		if !(i >= 32 && i <= 126) {
			return false
		}
	}
	return true
}

func main() {
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", Handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	fmt.Println("Listening at localhost:6969\nHttp Status :", http.StatusOK )
	openbrowser("http://localhost:6969")
	err := http.ListenAndServe(":6969", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func openbrowser(zz string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", zz).Start()
	}
	if err != nil {
		log.Fatal(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/css/img/favicon.ico")
}