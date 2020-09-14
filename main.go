package main

import (
	"fmt"
	"html/template"
	"log"
	"mime/multipart"
	"net/http"
	"regexp"
	"strings"

	"vtt-to-docx/unioffice/document"

	"github.com/asticode/go-astisub"
	"github.com/unidoc/unioffice/color"
)

var txtRegx = regexp.MustCompile("^(<(.*)>)?(.*)$")

func createParaRun(doc *document.Document, s string) document.Run {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(s)
	return run
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	doc, err := convertVTTToDocX(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.WriteHeader(http.StatusOK)
	doc.Save(w)
	return
}

// func (_cfc *Document) AddHeader() Header {
// 	_fbe := _fgg.NewHdr()
// 	_cfc._fbc = append(_cfc._fbc, _fbe)
// 	_ege := _cf.Sprintf("\u0068\u0065\u0061d\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", len(_cfc._fbc))
// 	_cfc._efe.AddRelationship(_ege, _c.HeaderType)
// 	_cfc.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_ege, "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0068\u0065\u0061\u0064e\u0072\u002b\u0078\u006d\u006c")
// 	_cfc._ff = append(_cfc._ff, _aeb.NewRelationships())
// 	return Header{_cfc, _fbe}
// }

func main() {
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/", Home)
	log.Fatalln(http.ListenAndServe(":9090", nil))

}

func convertVTTToDocX(f multipart.File) (*document.Document, error) {
	doc := document.New()

	// s1, _ := astisub.OpenFile("example-in.vtt")
	// fmt.Println(s1)
	s1, err := astisub.ReadFromWebVTT(f)
	if err != nil {
		return nil, err
	}

	hdr := doc.AddHeader()
	hdr.AddParagraph().AddRun().AddText("Header")

	for _, i := range s1.Items {
		fmt.Println(i.StartAt)
		fmt.Println(i.String())

		para := doc.AddParagraph()

		run := para.AddRun()
		run.Properties().SetItalic(true)
		run.Properties().SetSize(10)
		run.Properties().SetColor(color.DimGray)
		run.AddText(fmt.Sprintf("@ %s (%s)", i.StartAt.String(), i.EndAt-i.StartAt))
		run.AddBreak()

		t := txtRegx.FindStringSubmatch(i.String())
		_, _, author, text := t[0], t[1], t[2], t[3]
		run = para.AddRun()
		run.Properties().SetBold(true)
		run.AddText(strings.ToUpper(author))
		if author != "" {
			run.AddText(": ")
		}

		run = para.AddRun()
		run.AddText(text)
		run.AddBreak()
		run.AddBreak()
	}

	ftr := doc.AddFooter()
	para := ftr.AddParagraph()
	run := para.AddRun()
	run.AddText("Some subtitle goes here")

	// doc.SaveToFile("simple.docx")

	return doc, nil
}
