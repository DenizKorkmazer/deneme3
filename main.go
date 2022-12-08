package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

func add_user(isimsoyisim string, telefon string, eposta string) bool {

	db, err := sql.Open("mysql", "root:1Mhszxisq4r@tcp(127.0.0.1:3306)/deneme1")
	if err != nil {
		panic(err)
	}
	add, err := db.Query("INSERT INTO telefondefter (isimsoyisim,telefon,eposta) VALUES (?,?,?)", (isimsoyisim), (telefon), (eposta))
	if err != nil {
		panic(err)
	}

	fmt.Println(add)
	defer db.Close()
	return true
}

func kayitekle1(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("html/index.html"))
	tmplt.Execute(w, nil)
}

func kayitekle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var isimsoyisim = r.Form["isimsoyisim"]
	var telefon = r.Form["telefon"]
	var eposta = r.Form["eposta"]
	fmt.Println(isimsoyisim, telefon, eposta)
	if add_user(isimsoyisim[0], telefon[0], eposta[0]) {
		var tmplt = template.Must(template.ParseFiles("html/index3.html"))
		tmplt.Execute(w, nil)
	} else {
		var tmplt = template.Must(template.ParseFiles("html/error.html"))
		tmplt.Execute(w, nil)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./layouts"))
	http.Handle("/layouts/*", http.StripPrefix("layouts", fs))
	http.HandleFunc("/", kayitekle1)
	http.HandleFunc("/kayitekle", kayitekle)
	http.ListenAndServe(":8080", nil)

}
