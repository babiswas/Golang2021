package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string    `json:"name"`
	Id    int64     `json:"id"`
	Hobby [5]string `json:"hobby"`
}

func main() {
	p := Person{"Bapan", 12, [5]string{"football", "volley", "cricket", "judo", "basketball"}}
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	var per Person
	data1 := json.RawMessage(string(data))
	bytes, _ := data1.MarshalJSON()
	err := json.Unmarshal(bytes, &per)
	if err != nil {
		fmt.Println("Error occured")
		return
	}
	fmt.Println(per)
}
