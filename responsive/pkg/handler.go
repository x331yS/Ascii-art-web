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
					//a := strings.NewReader(ex.Out)
					//w.Header().Set("Content-Disposition", "attachment; filename=output.txt")
					//w.Header().Set("Content-Length", strconv.Itoa(len(ex.Out)))
					//io.Copy(w, a)
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

//func Handler(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/" {
//		StatusNotFound(w ,r)
//	}
//	switch r.Method {
//
//	case "GET":
//		var t, _ = template.ParseFiles("templates/index.html")
//		t.Execute(w, nil)
//
//	case "POST":
//
//		var text = r.FormValue("text")
//		var banner = r.FormValue("banner")
//
//		if ValidAscii(text) == true {
//
//			page = Page{
//				In:   text,
//				Out: Ascii(text, banner),
//			}
//
//			var t, _ = template.ParseFiles("templates/index.html")
//			t.Execute(w, page)
//
//			var output = strings.NewReader(page.Out)
//			w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
//			w.Header().Set("Content-Length", strconv.Itoa(len(page.Out)))
//			io.Copy(w, output)
//		}
//		if ValidAscii(text) == false {
//			BadRequest(w, r)
//			return
//
//		} else {
//			InternalServerError(w, r)
//		}
//	}
//}