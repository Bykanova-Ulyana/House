package family

// Структура, описывающая членов семьи

type FamilyMember struct {
	Name       string // Имя
	Surname    string // Фамилия
	Patronymic string // Отчество
	Relation   string // Статус в семье
	Age        int    // Возраст
	Birthday   string // Дата рождения
	Gender     string // Пол, true - ж, false - м
	Allergy    string // Наличие аллергии
}
