package rooms

import (
	"fmt"
	"golang/home/House/appliances"
	"golang/home/House/crockery"
	"golang/home/House/family"
	"golang/home/House/furniture"
	"golang/home/House/plumbing"
	"golang/home/House/windows"
)

// Функция, которая проверяет, можно ли установить предмет на полу и влезет ли он по высоте

func CanFitOnFloor(furniture []furniture.Furniture, room Room, newFurniture furniture.Furniture) bool {
	// Проверяем, достаточно ли места на полу для новой мебели
	if newFurniture.Length <= room.Length && newFurniture.Width <= room.Width {
		// Проверяем, влезет ли новая мебель по высоте
		heightLimit := room.Height - newFurniture.Height
		for _, item := range furniture {
			if item.Height > heightLimit {
				return false
			}
		}
		return true
	}
	return false
}

// считаем площадь комнаты

func (r Room) AreaCalculate() float32 {
	return r.Width * r.Length
}

func PrintHome(h Home) {
	for _, room := range h.Rooms {
		fmt.Printf("%s\n размеры (м):%f x %f x %f, площадь комнаты: %f \nВ комнате есть: \n\n",
			room.Type, room.Length, room.Width, room.Height, room.AreaCalculate())
		appliances.PrintAppliance(room.Appliances)
		furniture.PrintFurniture(room.Furniture)
		crockery.PrintCrockery(room.Crockery)
		family.PrintPet(room.Pets)
		plumbing.PrintCrockery(room.Plumbings)
		family.PrintFamilyMember(room.Family)
		windows.PrintWindow(room.Window)
	}
}
