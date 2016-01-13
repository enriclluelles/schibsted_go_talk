package pokemon

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

//start pokemon_type OMIT
type Evolution struct {
	To string `json:"to"`
}

type Pokemon struct {
	Id         json.Number `json:"national_id"`
	Name       string      `json:"name"`
	Evolutions []Evolution `json:"evolutions"`
}

//end pokemon_type OMIT

//start pokemon_get OMIT
func GetPokemons(f func(p *Pokemon), ids ...int) {
	url := "http://pokeapi.co/api/v1/pokemon/%d"
	for _, pokeid := range ids {
		go func(i int) {
			resp, err := http.Get(fmt.Sprintf(url, i))
			if err != nil {
				log.Println(err)
				return
			}

			var p Pokemon
			dec := json.NewDecoder(resp.Body)
			dec.UseNumber()

			if err := dec.Decode(&p); err == nil {
				f(&p)
			} else {
				log.Println(err)
			}
		}(pokeid)
	}
}

//end pokemon_get OMIT

//end pokemon_get OMIT

//start pokemon_print OMIT
func (p *Pokemon) Print() {
	ev := make([]string, 0)
	for _, v := range p.Evolutions {
		ev = append(ev, v.To)
	}
	fmt.Printf("Id: %s\n", p.Id)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Name: %s\n", p.Name)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Evolutions: %s\n", strings.Join(ev, ", "))
	fmt.Println()
}

//end pokemon_print OMIT
