// _range_ перебирает элементы в различных структурах
// данных. Давайте посмотрим, как использовать
// `range` с некоторыми из структур данных, которые
// мы уже изучили.

package main

import "fmt"

func main() {

	// В данном примере мы используем `range` для
	// подсчета суммы чисел в срезе.
	// Для массива синтаксис будет такой же.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` для массивов и срезов возвращает индекс
	// и значение для каждого элемента. Если нам не
	// требуется индекс, мы можем использовать оператор
	// `_` для игнорирования. Иногда нам действительно
	// необходимы индексы.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` для карт перебирает пары ключ/значение.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` может перебирать только ключи в карте
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` для строк перебирает кодовые точки Unicode.
	// Первое значение - это начальный байтовый индекс
	// руны, а второе - сама руна.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}