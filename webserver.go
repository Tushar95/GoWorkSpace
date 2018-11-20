package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AutoGenerated struct {
	SquadName  string   `json:"squadName"`
	HomeTown   string   `json:"homeTown"`
	Formed     int      `json:"formed"`
	SecretBase string   `json:"secretBase"`
	Active     bool     `json:"active"`
	Members    []Member `json:"members"`
}

type Member struct {
	Name           string   `json:"name"`
	Age            int      `json:"age"`
	SecretIdentity string   `json:"secretIdentity"`
	Powers         []string `json:"powers"`
}

type myOwnJSON struct {
	Name   string   `json:"name"`
	Skills []Skills `json:skills`
}

type Skills struct {
	SkillName  string   `json:"skillname"`
	Skillscore []int `json:"skillscore"`
}

func creatingMyJSON() {

	myJSON := myOwnJSON{"Tushar", [ {"Java",{0,1,2}}, {"Java",{0,1,2}}  ]}
}

func indexpage(w http.ResponseWriter, r *http.Request) {

	jsonString := []byte(`{"squadName":"Super hero squad","homeTown":"Metro City","formed":2016,"secretBase":"Super tower","active":true,"members":[{"name":"Molecule Man","age":29,"secretIdentity":"Dan Jukes","powers":["Radiation resistance","Turning tiny","Radiation blast"]},{"name":"Madame Uppercut","age":39,"secretIdentity":"Jane Wilson","powers":["Million tonne punch","Damage resistance","Superhuman reflexes"]},{"name":"Eternal Flame","age":1000000,"secretIdentity":"Unknown","powers":["Immortality","Heat Immunity","Inferno","Teleportation","Interdimensional travel"]}]}`)
	var check AutoGenerated
	json.Unmarshal(jsonString, &check)
	fmt.Printf("check this out %s", check.Members[0].Name)
	w.Write([]byte("pong"))
	w.Write(jsonString)
	JSONHandler(jsonString)
}

func JSONHandler(input []byte) {

	fmt.Println(string(input))
}

func main() {

	http.HandleFunc("/", indexpage)

	http.ListenAndServe(":8000", nil)
}
