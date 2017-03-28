package main

import (
	"fmt"
	"net/http"
	"os"

	//	"strconv"
	//	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"

	//	"image"
	//	_ "image/gif"
	//	"image/jpeg"
	//	_ "image/png"
)

////------------ Объявление типов и глобальных переменных

var (
	hd             string
	user           string
	nameConfigFile string // имя конфигурационного файла
)

var (
	tekuser     string // текущий пользователь который задает условия на срабатывания
	pathcfg     string // адрес где находятся папки пользователей, если пустая строка, то текущая папка
	pathcfguser string
)

type page struct {
	Title  string
	Msg    string
	Msg2   string
	TekUsr string
}

type DataInstagramStyle struct {
	UrlImage   string // урл картинки к
	TextMesage string
}

type ArrayDataInstagramStyle []DataInstagramStyle

var (
	defaultUrlImage    string = "http://vashgolos.net/photo/life/64084_0.jpg"
	dataInstagramStyle ArrayDataInstagramStyle
)

// функция проверки имени пользователя
func authFunc(username, password string) bool {
	return (auth.SecureCompare(username, "admin") && auth.SecureCompare(password, "qwe123!!"))
}

// обработчик начальной страницы
func indexHandler(user auth.User, rr render.Render, w http.ResponseWriter, r *http.Request) {
	rr.HTML(200, "index", &page{Title: "Ask like Instagram-style", Msg: "Начальная страница", TekUsr: "Текущий пользователь: " + string(user)})
}

// просмотр конфига
func ViewConfigFileHandler(user auth.User, rr render.Render, w http.ResponseWriter, r *http.Request) {

	//	rr.HTML(200, "view", &page{Title: "Ask like Instagram-style", Msg: "Начальная страница", TekUsr: "Текущий пользователь: " + string(user)})
	rr.HTML(200, "view", dataInstagramStyle)
}

func main() {

	dataInstagramStyle = make(ArrayDataInstagramStyle, 0)

	dataInstagramStyle = append(dataInstagramStyle, DataInstagramStyle{UrlImage: defaultUrlImage, TextMesage: "комментарий к фото"})
	dataInstagramStyle = append(dataInstagramStyle, DataInstagramStyle{UrlImage: "http://cdn.fishki.net/upload/post/201510/27/1714378/3_1_4.jpg", TextMesage: "комментарий к фото2"})

	fmt.Println("")

	m := martini.Classic()

	if pathcfg == "" {
		pathcfguser = ""
	} else {
		pathcfguser = pathcfg + string(os.PathSeparator)
	}

	m.Use(render.Renderer(render.Options{
		Directory:  "templates", // Specify what path to load the templates from.
		Layout:     "layout",    // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}}))

	m.Use(auth.BasicFunc(authFunc))

	m.Get("/view", ViewConfigFileHandler)
	m.Get("/", indexHandler)
	m.RunOnAddr(":9999")

}
