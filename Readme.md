## Задания

- Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

- Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

- Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

- Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

- Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

- Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.


- Реализовать все возможные способы остановки выполнения горутины.


- Реализовать конкурентную запись данных в map.


- Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.


- Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.


- Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.


- Реализовать пересечение двух неупорядоченных множеств.


- Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.


- Поменять местами два числа без создания временной переменной.


- Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.


- К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}

func main() {
someFunc()
}


- Реализовать быструю сортировку массива (quicksort) встроенными методами языка.


- Реализовать бинарный поиск встроенными методами языка.


- Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.


- Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.


- Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».


- Реализовать паттерн «адаптер» на любом примере.


- Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.


- Удалить i-ый элемент из слайса.


- Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.


- Реализовать собственную функцию sleep.

- Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true 
abCdefAaf — false 
aabcd — false
