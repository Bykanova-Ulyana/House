package rooms

import (
	"Home/House/appliances"
	"Home/House/crockery"
	"Home/House/family"
	"Home/House/furniture"
	"Home/House/windows"
)

func (home Home) CreatHome() Home {
	home = Home{Rooms: []Room{{

		// ---------------ПРИХОЖАЯ------------------

		Type:   "Прихожая",
		Width:  3.15,
		Length: 3.47,
		Height: 2,
		Appliances: []appliances.Appliance{{ // БЫТОВЫЕ ПРИБОРЫ
			Name:   "Пылесос",
			Brand:  "Hyundai",
			Power:  220,
			Width:  0.3,
			Length: 0.29,
			Height: 0.49,
			IsOn:   false,
		}},
		Furniture: []furniture.Furniture{{ // МЕБЕЛЬ
			Type:     "Прихожая",
			Number:   1,
			Height:   0.19,
			Width:    0.4,
			Length:   1,
			Color:    "белый",
			Material: "ЛДСП",
		}, {
			Type:     "Зеркало",
			Number:   1,
			Height:   0.9,
			Width:    0.6,
			Length:   0.05,
			Color:    "белый",
			Material: "ЛДСП",
		}, {
			Type:     "Когтеточка",
			Number:   1,
			Height:   1.5,
			Width:    0.3,
			Length:   0.5,
			Color:    "бежевый",
			Material: "дерево",
		}},
		Pets: []family.Pets{{ // ЖИВОТНЫЕ
			TypeAnimal:     "кошка",
			Breed:          "нет",
			Name:           "Мурка",
			Colour:         "серый",
			Weight:         3.3,
			SexAnimal:      true,
			AllergicFactor: "шерсть",
		}},
	}, {

		//------------------СПАЛЬНЯ----------------------

		Type:   "Спальня",
		Width:  3.18,
		Length: 5.94,
		Height: 2,
		Appliances: []appliances.Appliance{{ // МЕБЕЛЬ
			Name:   "Телевизор",
			Brand:  "Samsung",
			Power:  69,
			Width:  0.071,
			Length: 0.72,
			Height: 0.42,
			IsOn:   true,
		}, {
			Name:   "Торшер",
			Brand:  "Ambrella",
			Power:  55,
			Width:  0.84,
			Length: 1.7,
			Height: 0.84,
			IsOn:   true,
		}, {
			Name:   "Вентилятор",
			Brand:  "Aceline",
			Power:  55,
			Width:  0.44,
			Length: 1.3,
			Height: 0.39,
			IsOn:   false,
		}},
		Family: []family.FamilyMember{{
			Name:       "Анна",
			Surname:    "Быканова",
			Patronymic: "Николаевна",
			Relation:   "мать",
			Birthday:   "21.03.1978",
			Gender:     "ж",
			Allergy:    "нет",
		}, {
			Name:       "Федор",
			Surname:    "Быканов",
			Patronymic: "Иванович",
			Relation:   "отец",
			Birthday:   "26.04.1975",
			Gender:     "м",
			Allergy:    "нет",
		}},
		Crockery: []crockery.Crockery{{
			Type:      "Чайная пара",
			Purpose:   "сервировка стола",
			Number:    6,
			Material:  "фарфор",
			Engraving: false,
			Capacity:  250,
		}},
		Furniture: []furniture.Furniture{{
			Type:     "Диван",
			Number:   1,
			Height:   0.9,
			Width:    1.55,
			Length:   2.41,
			Color:    "тёмно-серый",
			Material: "ППУ, ДСП",
		}, {
			Type:     "Шкаф",
			Number:   1,
			Height:   1.6,
			Width:    0.5,
			Length:   2.3,
			Color:    "Венге",
			Material: "ЛДСП, зеркало",
		}, {
			Type:     "Сервант",
			Number:   2,
			Height:   2,
			Width:    6.9,
			Length:   3.6,
			Color:    "Венге",
			Material: "ЛДСП, МДФ, стекло",
		}},
	}, {

		//----------------------КУХНЯ---------------------
		Type:   "Кухня",
		Width:  2.1,
		Length: 3.15,
		Height: 2,
		Window: windows.Window{
			Number:     1,
			Length:     1,
			Width:      1.2,
			OpenWindow: false,
		},
	}, {

		//-------------------ВАННАЯ---------------------

		Type:   "Ванная",
		Width:  1.57,
		Length: 2.07,
		Height: 2,
	}}}

	return home
}
