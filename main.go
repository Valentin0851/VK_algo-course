package main

import (
	"fmt"
	"net/http"
	"errors"
	"github.com/mitchellh/mapstructure"
)

type User struct {
	name                 string
	age                  uint16
	money                int
	avg_grades, happines float64
}

type Emploee struct {
	id     int
	name   string
	age    uint32
	salary float32
}

type storage interface {
	insert(e Emploee) error
	get(id int) (Emploee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]Emploee
}

func newMemoryStorage() *memoryStorage{
	return &memoryStorage{
		data: make(map[int]Emploee),
	}
}

func (s *memoryStorage) insert(e Emploee) error{
	s.data[e.id] = e
	return nil
}

func (s *memoryStorage) get(id int) (Emploee, error){
	value, exist := s.data[id]
	if !exist{
		return Emploee{}, errors.New("This element isnt exist!")
	}

	return value, nil
}

func (s *memoryStorage) delete(id int) error{
	delete(s.data, id)

	return nil
}



type Point struct {
	x, y int
}

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is super easy")
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	// handleRequest()
	// newmap := map[string]int{}
	// newmap["Misha"] = 3
	pointmap := map[int]Point{}
	pointmap[1] = Point{
		x: 2,
		y: 2,
	}
	// userMap := map[int]User{}
	// userMap[1] = User{
	// 	name:       "Misha",
	// 	age:        7,
	// 	money:      1000000,
	// 	avg_grades: 3.0,
	// 	happines:   0.9,
	// }
	// fmt.Printf("X coord:%d", pointmap[1].x)
	// fmt.Println()
	// fmt.Printf("Y coord:%d", pointmap[1].y)
	// fmt.Println()
	// fmt.Printf("Happines:%f", userMap[1].happines)
	// fmt.Println()
	p1 := Point{}
	mapstructure.Decode(pointmap, &p1)
	// fmt.Println(p1)
}
