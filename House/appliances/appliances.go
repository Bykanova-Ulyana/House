package appliances

// Структура, описывающая бытовые приборы

type Appliance struct {
	Name   string  // Название
	Brand  string  // Бренд
	Power  int     // Напряжение
	Width  float32 // Ширина
	Length float32 // Длина
	Height float32 // Высота
	IsOn   bool    // Включен или выключен
}
