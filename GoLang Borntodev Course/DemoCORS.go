package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
func findID(id int) (*Course, int) {
	for i, course := range courselist {
		if course.CourseID == id {
			return &course, i
		}
	}
	return nil, 0
}
func courseHand(w http.ResponseWriter, r *http.Request) {
	urlpathsegment := strings.Split(r.URL.Path, "course/")
	id, err := strconv.Atoi(urlpathsegment[len(urlpathsegment)-1])

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	course, listitemindex := findID(id)
	if course == nil {
		http.Error(w, fmt.Sprintf("No course with id %d", id), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		coursejsonn, err := json.Marshal(course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(coursejsonn)
	case http.MethodPut:
		var updatecourse Course
		bytebody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bytebody, &updatecourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatecourse.CourseID != id {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		course = &updatecourse
		courselist[listitemindex] = *course
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
func coursesHand(w http.ResponseWriter, r *http.Request) {
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

func enableCourseMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Accept,Content-type,Content-Length,Authorization,X-CSRF-Token")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	courseitemhandler := http.HandlerFunc(courseHand)
	courselisthandler := http.HandlerFunc(coursesHand)
	http.Handle("/course/", enableCourseMiddleware(courseitemhandler))
	http.Handle("/course", enableCourseMiddleware(courselisthandler))
	http.ListenAndServe(":1234", nil)

}
