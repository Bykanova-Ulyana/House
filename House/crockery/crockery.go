package crockery

// Посуда

type Crockery struct {
	Type      string // Вид прибора
	Purpose   string // Назначение: 1. Сервировка стола 2. Кухонная утварь 3. Посуда для хранения
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
	Capacity  int    // Вместимость
}
