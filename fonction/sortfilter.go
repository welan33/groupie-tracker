package groupie

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func Number_of_members() {
	sort.Slice(Struct.TabD, func(i, j int) bool {
		return len(Struct.TabD[i].Members) < len(Struct.TabD[j].Members)
	})
}

func Creation_Date() {
	sort.Slice(Struct.TabD, func(i, j int) bool {
		return Struct.TabD[i].CreationDate < Struct.TabD[j].CreationDate
	})
}

func First_Album() {
	sort.Slice(Struct.TabD, func(i, j int) bool {
		date1 := strings.Split(Struct.TabD[i].FirstAlbum, "-")
		date2 := strings.Split(Struct.TabD[j].FirstAlbum, "-")
		if date1[2] < date2[2] {
			return true
		} else if date1[2] > date2[2] {
			return false
		} else {
			if date1[1] < date2[1] {
				return true
			} else if date1[1] > date2[1] {
				return false
			} else {
				if date1[0] < date2[0] {
					return true
				} else {
					return false
				}
			}
		}
	})
}

func RepareTab() {
	Struct.TabD = make([]Artist, len(Struct.TabA))
	copy(Struct.TabD, Struct.TabA)
}

func Caching() {
	Struct.Cache = make([]Artist, len(Struct.TabD))
	copy(Struct.Cache, Struct.TabD)
}

func NbMembres(w http.ResponseWriter, r *http.Request) {
	RepareTab()
	memberlist := []int{}
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusFound)
	case "POST":
		r.ParseForm()
		for i := 0; i < 9; i++ {
			if r.Form.Get("members"+strconv.Itoa(i)) == "on" {
				memberlist = append(memberlist, i)
			}
		}
		if len(memberlist) != 0 {
			Number_of_members()
			Struct.TabD = []Artist{}
			for _, num := range memberlist {
				for _, art := range Struct.TabA {
					if len(art.Members) == num {
						Struct.TabD = append(Struct.TabD, art)
					}
				}
			}
		}
		date := r.Form.Get("filter_startingyear")
		if date != "2022" {
			Creation_Date()
			for i := range Struct.TabD {
				if strconv.Itoa(Struct.TabD[i].CreationDate) > date {
					Struct.TabD = Struct.TabD[:i]
					Struct.Temps = date
					break
				}
			}
		} else {
			Struct.Temps = "2022"
		}
		if r.Form.Get("firstalbum") == "firstalbum" {
			First_Album()
		}
		SelectLoc := r.Form.Get("locations")
		if len(SelectLoc) != 0 {
			Caching()
			Struct.TabD = []Artist{}
			for _, art := range Struct.Cache {
				for _, pos := range art.Locations {
					pos = strings.Split(pos, "-")[1]
					if pos == SelectLoc {
						Struct.TabD = append(Struct.TabD, art)
						break
					}
				}
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
