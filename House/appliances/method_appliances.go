package appliances

import (
	"fmt"
)

func GetInputForAppliance() (Appliance, error) {
	var appliance Appliance
	var err [7]error
	fmt.Println("Введите название бытового прибора:")
	_, err[0] = fmt.Scanln(&appliance.Name)

	fmt.Println("Введите бренд бытового прибора:")
	_, err[1] = fmt.Scanln(&appliance.Brand)

	fmt.Println("Введите мощность бытового прибора:")
	_, err[2] = fmt.Scanln(&appliance.Power)

	fmt.Println("Введите высоту бытового прибора:")
	_, err[3] = fmt.Scanln(&appliance.Height)

	fmt.Println("Введите длину бытового прибора:")
	_, err[4] = fmt.Scanln(&appliance.Length)

	fmt.Println("Введите ширину бытового прибора:")
	_, err[5] = fmt.Scanln(&appliance.Width)

	fmt.Println("Включен ли бытовой прибор? (true/false):")
	_, err[6] = fmt.Scanln(&appliance.IsOn)

	for _, value := range err {
		if value != nil {
			return appliance, value
		}
	}
	return appliance, nil
}

// Функция для вывода информации о бытовых приборах

func PrintAppliance(a []Appliance) {
	for _, appliance := range a {
		fmt.Printf("%s\n Бренд: %s\n Размер: %f x %f x %f\n см Мощность: %i\n Прибор включён? %b", appliance.Name,
			appliance.Brand, appliance.Width, appliance.Length,
			appliance.Width, appliance.Power, appliance.IsOn)
	}

}
