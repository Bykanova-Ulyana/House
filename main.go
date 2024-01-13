package main

import (
	"Home/House/rooms"
)

func main() {
	var home rooms.Home
	home = home.CreatHome()
	rooms.PrintHome(home)
	/*for _, room := range Home.Rooms {
		for _, family := range room.Family {
			room = append(family.Family, family.GetInputForFamilyMember())
		}

	}*/
	println("Готово!")
}
