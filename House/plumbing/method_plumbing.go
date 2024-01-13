package plumbing

import "fmt"

// Функция для ввода данных о сантехнике

func GetInputForPlumbing() Plumbing {
	var plumbing Plumbing
	fmt.Println("Введите nип устройства (например, раковина, унитаз, душ и т.д.):")
	fmt.Scanln(&plumbing.Type)
	fmt.Println("Введите бренд:")
	fmt.Scanln(&plumbing.Brand)
	fmt.Println("Какой цвет у устройства:")
	fmt.Scanln(&plumbing.Color)
	fmt.Println("Введите материал:")
	fmt.Scanln(&plumbing.Material)
	fmt.Println("Вы установили устройство(да - true, нет - false: ")
	fmt.Scanln(&plumbing.Installed)
	return plumbing
}

// Функция для вывода информации о посуде

func PrintCrockery(plumbing []Plumbing) {
	if plumbing == nil {
		return
	}
	fmt.Print("\n\n -------САНТЕХНИКА-------\n")
	for _, pl := range plumbing {
		install := "установлено"
		if pl.Installed {
			install = "не установлено"
		}
		fmt.Printf("%s бренда %s имеет %s цвет, материал: %s, установлено: %s", pl.Type, pl.Brand, pl.Color, pl.Material, install)
	}
}

// Метод для установки устройства

func (device *Plumbing) Install() {
	device.Installed = true
	fmt.Println("Устройство успешно установлено!")
}

// Метод для удаления устройства

func (device *Plumbing) Uninstall() {
	device.Installed = false
	fmt.Println("Устройство успешно удалено!")
}
