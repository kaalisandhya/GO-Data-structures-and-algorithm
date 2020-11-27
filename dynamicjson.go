1.package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Envelope struct {
	Type string
	Msg  interface{}
}

type Sound struct {
	Description string
	Authority   string
}

type Cowbell struct {
	More bool
}

func main() {
	s := Envelope{
		Type: "sound",
		Msg: Sound{
			Description: "dynamite",
			Authority:   "the Bruce Dickinson",
		},
	}
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	c := Envelope{
		Type: "cowbell",
		Msg: Cowbell{
			More: true,
		},
	}
	buf, err = json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
}
{"Type":"sound","Msg":{"Description":"dynamite","Authority":"the Bruce Dickinson"}}
{"Type":"cowbell","Msg":{"More":true}}

Program exited.
2.package main
import (
   "fmt"
   "encoding/json"
)
type example struct {
    someStringFieldsName string   `json:"someStringFieldsName"`
    someIntFieldsName int `json:"someIntFieldsName"`
}
func main() {
  data, _:= json.MarshalIndent(&example{}, "", " ")
  fmt.Println(data)
}
[123 125]

Program exited.