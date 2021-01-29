package pkg

import (
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	In  string
	Out string
}


func FaviconHandler(w http.ResponseWriter, r * http.Request) {
	http.ServeFile(w, r, "assets/css/img/favicon.ico")
}


func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		switch r.Method {
		case "GET":
			t, err := template.ParseFiles("templates/index.html")
			if err != nil {
				InternalServerError(w, r)
			}
			_ = t.Execute(w, nil)
		case "POST":
			_ = r.ParseForm()
			if !ValidAscii(r.Form.Get("input")) {
				BadRequest(w, r)
			} else {
				output, status := AsciiOutput(r.Form["input"][0], r.Form["font"][0])
				if status == 500 {
					InternalServerError(w, r)
				} else {
					ex := Page{
						In:  r.Form["input"][0],
						Out: output,
					}
					t, err := template.ParseFiles("templates/index.html")
					if err != nil {
						InternalServerError(w, r)
						return
					}
					_ = t.Execute(w, ex)
					_ = os.Remove("assets/output/output.txt")
					file, err := os.OpenFile("assets/output/output.txt", os.O_CREATE, 0600)
					_, _ = file.WriteString(output)
					defer file.Close()
					if err != nil {
						panic(err)
					}
				}
			}
		}
	} else {
		StatusNotFound(w, r)

	}
}