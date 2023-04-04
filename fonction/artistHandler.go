package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
)

type Coord struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
	Date      string
}

var CoordTab []Coord

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	CoordTab = []Coord{}
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusFound)
		print("ok")
	case "POST":
		r.ParseForm()
		id, _ := strconv.Atoi(r.Form.Get("artist"))

		for i, artist := range Struct.TabD {
			if id == artist.Id {
				id = i
				break
			}
		}
		tmpl := template.Must(template.ParseFiles("./static/info.html"))
		data := Artist{
			Id:           Struct.TabD[id].Id,
			Image:        Struct.TabD[id].Image,
			Name:         Struct.TabD[id].Name,
			Members:      Struct.TabD[id].Members,
			CreationDate: Struct.TabD[id].CreationDate,
			FirstAlbum:   Struct.TabD[id].FirstAlbum,
			Locations:    Struct.TabD[id].Locations,
			Date:         Struct.TabD[id].Date,
			DateL:        Struct.TabD[id].DateL,
			CoordTab:     CoordTab,
		}
		//reminder := 0

		for i, v := range data.Locations {
			var co []Coord
			req, errb := http.Get("https://nominatim.openstreetmap.org/search/" + v + "?format=json&addressdetails=1&limit=1&polygon_svg=1")
			if errb != nil {
				fmt.Println(errb)
			}
			a, erb := ioutil.ReadAll(req.Body)
			if errb != nil {
				fmt.Println(erb)
			}
			json.Unmarshal(a, &co)
			co[0].Date = data.Date[i]
			/*if len(data.Date) > 1 && reminder == 0 {
			      co[0].Date = data.Date[reminder] //out of range
			      reminder += len(data.Date) - 1
			      //probleme avec le i, il n'y pas forcément le meme nombre de dates et de locations (ex xxxtentacion 3 pour 8) et où i = 3
			      } else if len(data.Date) > 1 && reminder > 0 {
			      co[0].Date = data.Date[reminder] //out of range
			      reminder += len(data.Date) - 1
			  } else {
			      co[0].Date = data.Date[reminder-1]
			      reminder++
			  }*/
			CoordTab = append(CoordTab, co...)
		}
		data.CoordTab = CoordTab
		// fmt.Println(data)
		err := tmpl.Execute(w, data)
		if err != nil {
			fmt.Fprintln(w, err)
		}
	}
}
