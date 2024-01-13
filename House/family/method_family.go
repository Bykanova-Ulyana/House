package family

import (
	"fmt"
	"strings"
	"time"
)

// Функция для ввода данных о членах семьи
func GetInputForFamilyMember() FamilyMember {
	var familyMember FamilyMember
	fmt.Println("Введите ФИО члена семьи: ")
	fmt.Scanln(&familyMember.Surname, &familyMember.Name, &familyMember.Patronymic)
	fmt.Println("Введите отношение к члену семьи: ")
	fmt.Scanln(&familyMember.Relation)
	fmt.Println("Введите дату рождения (дд.мм.гггг): ")
	fmt.Scanln(&familyMember.Birthday)
	ageCalculation(familyMember.Birthday)
	fmt.Println("Введите пол (м - мужской, ж - женский): ")
	fmt.Scanln(&familyMember.Gender)
	fmt.Println("Есть ли у члена семьи аллергия(-и) (да - какие, нет): ")
	fmt.Scanln(&familyMember.Allergy)
	return familyMember
}

// Функция для вывода информации о семье

func PrintFamilyMember(family []FamilyMember) {
	for _, fm := range family {
		fmt.Printf(" %s\nФИО: %s %s %s\n Дата рождения: %s\nВозраст: %d\nПол: %s\n", fm.Relation,
			fm.Surname, fm.Name, fm.Patronymic, fm.Birthday, fm.Age, fm.Gender)
	}

}

func ageCalculation(birthday string) int {

	// Преобразуем строку в формат времени
	layout := "02.01.2006"
	data, err := time.Parse(layout, birthday)
	if err != nil {
		fmt.Println("Неправильный формат даты.")
	}

	// Получаем текущую дату
	now := time.Now()

	// Вычисляем разницу между текущей датой и датой рождения
	years := now.Year() - data.Year()

	// Проверяем, достиг ли человек день рождения в этом году
	if now.YearDay() < data.YearDay() {
		years--
	}

	return years
}

func AllergyCheck(member FamilyMember, allergicFactor string) {
	if strings.Contains(member.Allergy, allergicFactor) {
		fmt.Println("%s имеет аллергию на %s", member.Name, allergicFactor)
		return
	}
	fmt.Println("%s не имеет аллергию на %s", member.Name, allergicFactor)

}
