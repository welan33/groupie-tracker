package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ArtistsStruct struct {
	TabA  []Artist
	TabD  []Artist
	Cache []Artist
	Loc   []string
	Temps string
}

type Artist struct {
	Id           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"locations"`
	Date         []string            `json:"dates"`
	DateL        map[string][]string `json:"datesLocations"`
	CoordTab     []Coord
}

var Struct ArtistsStruct
var ArtistsTab []Artist

func APIRequest() {

	Apilist := []string{"artists", "locations", "dates", "relation"}

	for _, API := range Apilist {
		request, err := http.Get("https://groupietrackers.herokuapp.com/api/" + API)
		if err != nil {
			fmt.Println(err)
		}
		content, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println(err)
		}
		if API != Apilist[0] {
			content = content[9 : len(content)-2]
		}
		json.Unmarshal(content, &ArtistsTab)
	}

	Struct = ArtistsStruct{
		TabA: ArtistsTab,
	}
	Struct.Temps = "2022"
}

func SelectCountry() {
	verif := false
	for i := range Struct.TabA {
		for _, loc := range Struct.TabA[i].Locations {
			loc = strings.Split(loc, "-")[1]
			verif = false
			for _, val := range Struct.Loc {
				if val == loc {
					verif = true
				}
			}
			if !verif {
				Struct.Loc = append(Struct.Loc, loc)
			}
		}
	}
}
