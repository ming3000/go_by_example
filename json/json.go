package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"os"
)

type Book struct {
	Name string `json:"the_book_name"`
	Authors []string `json:"the_author_list"`
}

type User struct {
	Name string
	Age int
	phone string
}

func line() {
	fmt.Println(strings.Repeat("-", 30))
}

func marshals() {
	jBool, _ := json.Marshal(true)
	fmt.Println(string(jBool))
	jInt, _ := json.Marshal(3)
	fmt.Println(string(jInt))
	jStr, _ := json.Marshal("golang")
	fmt.Println(string(jStr))
	line()

	jSlice, _ := json.Marshal([]string{"python", "golang", "node.js"})
	fmt.Println(string(jSlice))
	jMap, _ := json.Marshal(map[string]int{"apple": 5, "lettuce": 7})
	fmt.Println(string(jMap))
	line()

	objBook := Book{
		Name: "Node.js Design Patterns",
		Authors: []string{"Mario Casciaro", "Luciano Mammino", "冯康", " 孙光金", " 丁全"},
	}
	retBook, _ := json.Marshal(objBook)
	fmt.Println(string(retBook))

	objUser := &User{
		Name: "kafka",
		Age: 30,
		// won't be marshaled
		phone: "15623093333",
	}
	retUser, _ := json.Marshal(objUser)
	fmt.Println(string(retUser))
}

func userUnMarshals() {
	by := []byte(`{"Name": "json", "Age": 50}`)
	var obj map[string]interface{}
	if err := json.Unmarshal(by, &obj); err != nil {
		panic(err)
	}
	fmt.Println(obj)
	name := obj["Name"].(string)
	// must be float64, can not be int...
	age := obj["Age"].(float64)
	fmt.Printf("result, Name is %s, Age is %f \n", name, age)

	line()

	userStr := `{"Name": "wwj", "Age": 31}`
	userObj := User{}
	err := json.Unmarshal([]byte(userStr), &userObj)
	if err != nil {
		panic(err)
	}
	fmt.Println(userObj)
	fmt.Printf("user name is:%s, age is:%d", userObj.Name, userObj.Age)
}

func bookUnMarshal() {
	by := []byte(`{"name": "go in action", "authors": ["William Kennedy", "Brian Ketelsen", "Erik St.Martin"]}`)
	var obj map[string]interface{}
	if err := json.Unmarshal(by, &obj); err != nil {
		panic(err)
	}
	fmt.Println(obj)

	name := obj["name"].(string)
	fmt.Println("name:", name)

	tempStr := obj["authors"].([]interface{})
	author1 := tempStr[0].(string)
	fmt.Println("first author:", author1)
	author2 := tempStr[1].(string)
	fmt.Println("second author:", author2)

	line()

	// TODO, watch out of this! both two are wrong!
	//bookStr := `{"name": "go in action", "authors": ["William Kennedy", "Brian Ketelsen", "Erik St.Martin"]}`
	//bookStr := `{"Name": "go in action", "Authors": ["William Kennedy", "Brian Ketelsen", "Erik St.Martin"]}`
	bookStr := `{"the_book_name": "go in action", "the_author_list": ["William Kennedy", "Brian Ketelsen", "Erik St.Martin"]}`
	bookObj := Book{}
	if err := json.Unmarshal([]byte(bookStr), &bookObj); err != nil {
		panic(err)
	}
	fmt.Println(bookObj)
	fmt.Println("name:", bookObj.Name)
	fmt.Println("authors:", bookObj.Authors)
}

func jsonCode() {
	enc := json.NewEncoder(os.Stdout)
	data := map[string]int{"apple": 5, "orange": 6}
	enc.Encode(data)

	user := User{Name:"wwj", Age:666}
	enc.Encode(user)
}

func main() {
	//marshals()
	//userUnMarshals()
	//bookUnMarshal()
	jsonCode()
}
