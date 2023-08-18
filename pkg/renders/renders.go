package renders

import (
	// "bytes"
	"bytes"
	"fmt"
	"gocourse/pkg/config"
	"gocourse/pkg/models"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)
var App *config.AppConfig
func NewTemplate(a *config.AppConfig){
	App = a
}
// renders a template
func RenderTemplate(w http.ResponseWriter, html string , td *models.TemplateData) {
	var tc map[string]*template.Template
	if App.UseCache {
 		tc = App.TemplateCache
	} else {
		tc,_ = CreateCacheTemplate()
	}	
		//create a template cache
	// tc, err := CreateCacheTemplate()
	
	// fmt.Println(tc)
	// get requested template from cache
	t, inMap := tc[html]
	if !inMap {
		log.Fatal("Could not get template")

	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)

	}

	//render the template
	_, err = buf.WriteTo(w)
	if err!= nil{
		log.Fatal(err)
	}

	// parses the  given file

	//for simple cache
	// parsed, _ := template.ParseFiles("./templates/" + html, "./templates/base.layout.html")
	// err := parsed.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("Error while parsing: ", err)
	// 	return
	// }
}

func CreateCacheTemplate() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}
	// fmt.Println(pages)
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	// fmt.Println(myCache)
	return myCache, nil

}

//Simple Cache With Map

// var tc = make(map[string]*template.Template)

// func cacheTemplate(t string) error{
// 	temps := []string{
// 		fmt.Sprintf("./templates/%s",t),
// 		"./templates/base.layout.html"}
// 	parsed, err := template.ParseFiles(temps...)
// 	if err!= nil{
// 		return err
// 	}
// 	tc[t] = parsed
// 	return nil
// }

// func RenderTemplate(w http.ResponseWriter, t string){
// 	var htm *template.Template
// 	var err error
// 	_, inMap := tc[t]
// 	if !inMap {
// 		log.Println("Exceuting Without Cache")
// 		err = cacheTemplate(t)
// 		if err!=nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		log.Println("Executing With Cache")
// 	}
// 	htm = tc[t]
// 	err = htm.Execute(w,nil)
// 	if err!= nil {
// 		log.Println(err)
// 	}
// }
