package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/** JSONデコード用に構造体定義 */
type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("vro.json")
	if err != nil {
		log.Fatal(err)
	}

	// // JSONデコード
	// var persons []Person
	// if err := json.Unmarshal(bytes, &persons); err != nil {
	// 	log.Fatal(err)
	// }
    //
	// // デコードしたデータを表示
	// for _, p := range persons {
	// 	fmt.Printf("%d : %s\n", p.Id, p.Name)
	// }

    fmt.Fprintf(w, string(bytes[:]))
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":3218", nil)
}
