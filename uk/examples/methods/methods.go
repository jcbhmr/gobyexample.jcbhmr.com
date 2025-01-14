// Go підтримує _методи_ визначені на типі структура.

package main

import "fmt"

// Ось - ми визначимо наш тип.
type rect struct {
	width, height int
}

// Функція `area` (площа) являє собою метод типу `rect`,
// або як кажуть з отримувачем типу *`rect`*.
func (r *rect) area() int {
	return r.width * r.height
}

// Методи можуть визначатись для отримувачів типу
// вказівник або значення. Ось приклад метода з
// отримувачем типу значення. Методи з таким типом не
// змінюють значення стуктури.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Тут ми викличемо 2 методи визначені для нашої структури.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go автоматично визначає чи потрібне конвертування між
	// значеннями та вказівниками для виклику методів. Можливо
	// вам заманеться використати отримувач-вказівник, щоб уникнути
	// копіювання даних при виклику методу або для зміни даних
	// в структурі.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
