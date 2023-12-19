package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// 2) создали мапу, в которой будут храниться отпарсированные страницы
var tc = make(map[string]*template.Template)

// 1) создали функцию с временной приставкой Test
// 21) изменили название, убрав Test
func RenderTemplate(w http.ResponseWriter, t string) {
	// 3) объявили переменную которой, по-видимому, будем присваивать отпарсированную страницу
	var tmpl *template.Template

	// 4) объявили переменную err, я не знаю что с ней будем делать
	var err error

	// check to see if we already have the template in our cache
	// 5) техника "comma ok", в качестве ключа используется параметр t, который был передан с вызовом функции
	// проверяем есть ли такой ключ в мапе или нет
	// если ключ есть, значит страницу уже парсировали и она в кеше, если нет, то нужно отпарсировать и добавить в кеш
	_, inMap := tc[t]
	// 6) добавляем инструкцию if, что бы определить поток выполнения кода на основании того, есть ключ в мапе или нет
	if !inMap {
		// need to create template
		// 18) не понимаю зачем тут переменная err и почему нельзя просто вызывать функцию;
		// почти сразу же понял, это нужно для того, что бы получить ошибку, если она будет;
		err = createTemplateCache(t)
		// 19) обрабатываем ошибку, если она будет
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}

	// 7) достаем их кеша отпарсированную страницу и присваиваем ее копию переменной tmpl
	tmpl = tc[t]

	// 8) я не понимаю этого действия
	err = tmpl.Execute(w, nil)
	// 20) обрабатываем ошибку, если она будет здесь
	if err != nil {
		log.Println(err)
	}
}

// 9) определяем функцию, действие которой, на момент ее создания, я не понимаю в полной мере
func createTemplateCache(t string) error {
	// 10) объявляем переменную с типом slice of strings, куда будем добавлять, наверное путь, необходимый для рендеринга страницы
	templates := []string{
		// 11) функция Sprintf возвращает все содержимое как одну строку; напоминаю, что t - это то, куда обратился пользователь, /home, /about и т.д.
		fmt.Sprintf("./templates/%s", t),
		// 12) я понимаю, что инфо из предыдущего пункта и этот функт должны передаваться вместе в будущем для рендеринга страницы,
		// но не понимаю, почему они содержатся в слайсе как разные элементы, что будет дальше?
		"./templates/base.layout.tmpl",
	}

	// parse the template
	// 13) происходит парсирование из того что было указано в предыдущем пункте, результат присваивается tmpl и err на случай, если будет ошибка
	tmpl, err := template.ParseFiles(templates...)
	// 14) если ошибка есть (т.е. не равна nil), то возвращаем ошибку
	if err != nil {
		return err
	}
	
	// add template to cache (map)
	// 15) добавляем темплейт в мапу
	tc[t] = tmpl

	// 16) т.к. в мапу был добавлен темплейт, то ошибки нет, значит возвращаем nil
	return nil

	// 17) функция не возвращает мапу или какие-то значения, в функции в мапу был добавлен темплейт, мапа объявлена на уровне пакета, т.е.
	// весь необходимый функционал выполнен, отсюда следует что функции помимо своего дейфствия может лишь вернуть ошибку если она будет
}
