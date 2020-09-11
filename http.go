package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type emp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

var arr = []emp{}

func get(e *emp) string {
	return (e.ID + e.Name + e.Description + e.Date)
}

/*func getAll(arr []emp) string {
	return ((arr[0].id + arr[0].name + arr[0].description + arr[0].date) + "\n" + (arr[1].id + arr[1].name + arr[1].description + arr[1].date))
}*/

func post(reqbody []byte) string {
	var t emp
	err := json.Unmarshal(reqbody, &t)
	if err != nil {
		panic(err)
	}
	arr = append(arr, t)
	return t.ID
}

func put(reqbody []byte) string {
	var t emp
	err := json.Unmarshal(reqbody, &t)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(arr); i++ {
		if t.ID == arr[i].ID {
			arr[i] = t
		}
	}
	return "data set"
}
func patch(reqbody []byte) string {
	var t emp
	err := json.Unmarshal(reqbody, &t)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(arr); i++ {
		if t.ID == arr[i].ID {
			if t.Name != arr[i].Name {
				arr[i].Name = t.Name
			} else if t.Description != arr[i].Description {
				arr[i].Description = t.Description
			} else if t.Date != arr[i].Date {
				arr[i].Date = t.Date
			}
		}
	}
	return "ok"
}

func addCookie(w http.ResponseWriter, name, value string, ttl time.Duration) {
	expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		key := "id"
		val := r.URL.Query().Get(key)
		for i := 0; i < len(arr); i++ {
			if val == arr[i].ID {
				io.WriteString(w, get(&arr[i]))
			}
		}
	} else if r.Method == "POST" {
		reqbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		io.WriteString(w, post(reqbody))
	} else if r.Method == "PUT" {
		reqbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		io.WriteString(w, put(reqbody))
		//io.WriteString(w, post())
	} else if r.Method == "PATCH" {
		reqbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		io.WriteString(w, patch(reqbody))
	} else {
		io.WriteString(w, "404 page not found")
	}
	//io.WriteString(w, "Hello world!")
	addCookie(w, "TestCookieName", "TestValue", 30*time.Minute)

}

func main() {
	e := emp{ID: "E1", Name: "english", Description: "hgxahgsxkjhs", Date: "01/01/2001"}
	arr = append(arr, e)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
