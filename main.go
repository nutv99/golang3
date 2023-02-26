package main

import (
	"encoding/json"
	"fmt"

	sunday999 "github.com/nutv99/golang2"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{Name: "Alice", Age: 30}

	sunday999.Main999()
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

//require github.com/nutv99/golang2 v0.0.0-20230226112911-d14fd4e05398
/*
echo "# golang2" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/nutv99/golang2.git
git push -u origin main
*/

/*
git remote add origin https://github.com/nutv99/golang2.git
git branch -M main
git push -u origin main
*/
