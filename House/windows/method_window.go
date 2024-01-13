package windows

import "fmt"

func StatusOpenCloseWindow(w *Window, command bool) {
	switch command {
	case command == true || w.OpenWindow == true:
		println("Окно открыто")
	case false:
		println("Окно закрыто")
	default:
		println("ОШИБКА! ОКНО СЛОМАЛОСЬ!")
	}
}

func PrintWindow(w Window) {
	fmt.Print("\n\n -------ОКНА-------\n")
	if w.Number == 0 {
		print("\n\tОкон нет!\n\n")
		return
	}
	openClose := "открыто"
	if w.OpenWindow {
		openClose = "закрыто"
	}
	fmt.Printf("\n\tДлина окна %f\n\tШирина окна %f\n\tОкно открыто %s \n", w.Length, w.Width, openClose)
}
