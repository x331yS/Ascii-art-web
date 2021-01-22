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
//	Out    string
//
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
					//d.Out = output
				//}
				//if r.FormValue("process") == "download" {
				//	a := strings.NewReader(d.Out)
				//	w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
				//	w.Header().Set("Content-Length", strconv.Itoa(len(d.Out)))
				//	io.Copy(w, a)
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
		if i > 127 {
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

// Data is a struct that will be sent as a respond
//type Data struct {
//	Output    string
//	ErrorNum  int
//	ErrorText string
//}
//
//var temp *template.Template
//func main() {
//	http.HandleFunc("/", serverHandler)
//	// Creating a handler for handling static files
//	FileServer := http.FileServer(http.Dir("docs"))
//	http.Handle("/docs/", http.StripPrefix("/docs/", FileServer))
//	fmt.Println("Server is listening to port #8080 ... ")
//	http.ListenAndServe(":8080", nil)
//}
//
//func serverHandler(res http.ResponseWriter, req *http.Request) {
//	d := Data{}
//	temp = template.Must(template.ParseGlob("docs/htmlTemplates/*.html"))
//
//	if req.URL.Path != "/" {
//		d.ErrorNum = 404
//		d.ErrorText = "Page Not Found"
//		errorHandler(res, req, &d) // 404 ERROR
//		return
//	}
//
//	if req.Method == "GET" {
//		temp.ExecuteTemplate(res, "index.html", d)
//
//	} else if req.Method == "POST" {
//		// Gathering information to be processed
//		text := req.FormValue("input")
//		font := req.FormValue("font")
//
//		out, err := ascii.AsciiOutput(text, font)
//		if err {
//			d.ErrorNum = 500
//			d.ErrorText = "Internal Server Error"
//			errorHandler(res, req, &d)
//			return
//		}
//		d.Output = out
//
//		if req.FormValue("process") == "show" {
//			temp.ExecuteTemplate(res, "index.html", d)
//		} else if req.FormValue("process") == "download" {
//			a := strings.NewReader(d.Output)
//			res.Header().Set("Content-Disposition", "attachment; filename=file.txt")
//			res.Header().Set("Content-Length", strconv.Itoa(len(d.Output)))
//			io.Copy(res, a)
//		} else {
//			d.ErrorNum = 400
//			d.ErrorText = "Bad Request"
//			errorHandler(res, req, &d)
//			return
//		}
//	} else {
//		d.ErrorNum = 400
//		d.ErrorText = "Bad Request"
//		errorHandler(res, req, &d)
//		return
//	}
//}
//
//func errorHandler(res http.ResponseWriter, req *http.Request, d *Data) {
//	res.WriteHeader(d.ErrorNum)
//	temp.ExecuteTemplate(res, "error.html", d)
//}