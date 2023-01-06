package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Course struct {
	CourseID   int     `json:"id"`
	CourseName string  `json:"name"`
	Price      float64 `json:"price"`
	Instrutor  string  `json:"instrutor"`
}

var courselist []Course

func init() {
	coursejson := `[
		{
			"id":1,
			"name":"violin",
			"price":500,
			"instrutor":"Thone"
		},
		{
			"id":2,
			"name":"sesorrr",
			"price":1500,
			"instrutor":"Esanmukuount"
		},
		{
			"id":3,
			"name":"jinggabel",
			"price":700,
			"instrutor":"THTOhu"
		}
	]`
	err := json.Unmarshal([]byte(coursejson), &courselist)
	if err != nil {
		log.Fatal(err)
	}
}
func getNextid() int {
	highest := -1
	for _, course := range courselist {
		if highest < course.CourseID {
			highest = course.CourseID
		}
	}
	return highest + 1
}
func courseHand(w http.ResponseWriter, r *http.Request) {
	coursejsonnnnn, err := json.Marshal(courselist)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-Type", "application/json")
		w.Write(coursejsonnnnn)
	case http.MethodPost:
		var newcourse Course
		bodybyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodybyte, &newcourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newcourse.CourseID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newcourse.CourseID = getNextid()
		courselist = append(courselist, newcourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}
func main() {
	http.HandleFunc("/course", courseHand)
	http.ListenAndServe(":8080", nil)

}
