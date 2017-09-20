// Однострочный комментарий
/* Много
строчный комментарий*/

package main

import (
	"fmt"    // пакет в стандартной библиотеке Go
	m "math" // импортировать math под локальным именем m
)

// Объявление функции main. Это специальная функция,
// служащая точкой входа для исполняемой программы.
func main() {
	// Println выводит строку в stdout
	// Это функция из пакета fmt
	fmt.Println("Hello world!")
	beyondHello()
}

// Параметры функции указываются в круглых скобках
// пустые скобки обязательны, даже если параметров нет
func beyondHello() {
	var x int // Переменные должны быть объявлены до их использования
	x = 3     // Присвоение значения переменной
	// Краткое определение := позволяет объявить переменную с автоматической
	// подстановкой типа значчения
	y := 4
	sum, prod := learnMultiple(x, y)        // Функция возвращает два значения
	fmt.Println("sum:", sum, "prod:", prod) // Простой вывод
	learnTypes()
}

// Функция, имеющая входные параметры и возвращающая несколько значений
func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y // возврат двух значений
}

func learnTypes() {
	// тип string
	s := "Learn Go!"

	s2 := `"Чистый" строковый литерал
может содержать переносы строк`

	// Символы не из ASCII
	g := 'Σ' // тип rune - алиас для типа int32, содержит символ юникода

	f := 3.141595 // float64, 64-х битное число с плавающей точкой
	c := 3 + 4i   // complex128, внутри себя содержит два float64

	// Синтаксис var  с инициализ0ациями
	var u uint = 7 // беззнаковое, но размер зависит от реализации, как и у int
	var pi float32 = 22. / 7

	// Синтаксис приведения типа с определением
	n := byte('\n') // byte - это алиас для uint8

	// Массивы имеют фиксированный размер на момент компиляции
	var a4 [4]int           // массив из 4-х int, инициализирован нулями
	a3 := [...]int{3, 1, 5} // массив из 3-х int, ручная инициализация

	// Слайсы (slices) имеют динамическую длину. И массивы, и слайсы имеют свои
	// преимущества, но слайсы используются гораздо чаще
	s3 := []int{4, 5, 9}    // в отличие от a3, тут нет троеточия
	s4 := make([]int, 4)    // выделение памяти для слайса из 4-х int (нули)
	var d2 [][]float64      // только объявление, память не выделяется
	bs := []byte("a slice") // синтаксис приведения типов

	p, q := learnMemory() // объявление p и q как указателей на int
	fmt.Println(*p, *q)   // * извлекает указатель. Печатает два int-а

	// Map, также как и словарь или хэш из некоторых других языков, является
	// ассоциативным массивом с динамически изменяемым размером.
	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	delete(m, "three") // встроенная функция, удаляет элемент из map-а

	// неиспользуемые переменные в Go считаются ошибкой
	// Нижнее подчеркивание позволяет игнорировать такие переменные
	_, _, _, _, _, _, _, _, _ = s2, g, f, u, pi, n, a3, s4, bs
	// Вывод считается использованием переменной
	fmt.Println(s, c, a4, s3, d2, m)

	learnFlowControl() // управление потоком
}

// у Go есть сборщик мусора. В нём есть указатели, но нет арифметики
// указателей. Вы можете допустить ошибку с указателем на nil, но не с
// инкрементацией указателя
func learnMemory() (p, q *int) {
	// именованные возвращаемые значения p и q являются указателями на int
	p = new(int) // встроенная функция new выделяет память
	// выделенный int проинициализирован нулём, p больше не содержит nil
	s := make([]int, 20) // Выделение единого блока памяти под 20 int-ов
	s[3] = 7             // Присвоить значение одному из них
	r := -2              // определить ещё одну локальную переменную
	return &s[3], &r     // Амперсанд(&) обозначает получение адреса переменной
}

func expensiveComputation() float64 {
	return m.Exp(10)
}

func learnFlowControl() {
	// if-ы всегда требуют наличия фигурных скобок, но не круглых
	if true {
		fmt.Println("told ya")
	}
	// форматирование кода стандартизировано утилитой "go fmt"
	if false {
		// не выполняется
	} else {
		// выполняется
	}
	// вместо нескольких if исползуется switch
	x := 42.0
	switch x {
	case 0:
	case 1:
	case 42:
		// Case-ы в Go не "проваливаются" (неявный break)
	case 43:
		//  не выполняется
	}
	// for, как и if не требует круглх скобок
	// переменные, объявленные в for и if являются локальными
	for x := 0; x < 3; x++ { // ++ - это операция
		fmt.Println("итерация", x)
	}
	// здесь x == 42

	// for - единственный цикл в Go, но у него есть альтернативные формы
	for { // бесконечный цикл
		break    // не такой уж бесконечный
		continue // не выполнится
	}
	// как и в for, := в if-е означает объявление и присвоение значения y,
	// проверка y > x происходит после
	if y := expensiveComputation(); y > x {
		x = y
	}
	// Функции являются замыканиями
	xBig := func() bool {
		return x > 10000 // Ссылается на x, объявленный выше switch
	}
	fmt.Println("xBig", xBig()) // true (т.к. мы присвоили x = e^10)
	x = 1.3e3                   // тут x == 1300
	fmt.Println("xBig", xBig()) // Теперь false

	// метки
	goto love
love:

	learnDefer()
	learnInterfaces()
}

func learnDefer() (ok bool) {
	// отложенные (deferred) выражения выполняются сразу перед тем, как функция
	// возвратит значения
	defer fmt.Println("deferred statements execute in reverse (LIFO) order.")
	defer fmt.Println("\nThis line is being printed first because")
	// defer широко используется для закрытия файлов, чтобы закрывающая файл
	// функция находилась близко к открывающей
	return true
}

// Stringer объявление как интерфейса с одним методом, String
type Stringer interface {
	String() string
}

// объявление pair как структуры с двумя полями x и y типа int
type pair struct {
	x, y int
}

// объявление метода для типа pair; теперь pair реализует интерфейс Stringer
func (p pair) String() string { // p в данном случае называют receiver-ом
	// Sprintf - ещё одна функция из пакета fmt
	// Обращение к полям p через точчку
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func learnInterfaces() {
	// синтаксис с фигурными скобками это "литерал структуры". Он возвращает
	// проинициализированную структуру, а оператор := присваивает её p
	p := pair{3, 4}
	fmt.Println(p.String()) // вызов метода String у переменной p типа pair
	var i Stringer          // объявление i как типа с интерфейсом Stringer
	i = p                   // валидно, т.к. pair реализует Stringer
	// Вызов метода String у i типа Stringer
	fmt.Println(i.String())

	// Функция в пакете fmt сами всегда вызывают метод String у объектов для
	// получения их строкового представления
	fmt.Println(p) // вывод такой-же, как и выше
	fmt.Println(i) // вывод такой-же, как и выше
}
