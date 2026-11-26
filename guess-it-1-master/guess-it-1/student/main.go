package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var history []float64

	for scanner.Scan() {
		line := scanner.Text()

		// Преобразуем строку в число с плавающей точкой
		number, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue // Если не удалось преобразовать, продолжаем с другим числом
		}

		// Добавляем число в историю
		history = append(history, number)

		// Если в истории больше 50 чисел, удаляем старые
		if len(history) > 50 {
			history = history[1:]
		}

		// Если чисел в истории меньше двух, выводим диапазон вокруг числа
		n := float64(len(history))
		if n < 2 {
			fmt.Printf("%.0f %.0f\n", number-20, number+20)
			continue
		}

		// Вычисление среднего значения
		sum := 0.0
		for _, x := range history {
			sum += x
		}
		mean := sum / n

		// Вычисление стандартного отклонения
		varianceSum := 0.0
		for _, x := range history {
			varianceSum += math.Pow(x-mean, 2)
		}
		variance := varianceSum / n
		stdDev := math.Sqrt(variance)

		// Рассчитываем маржу как 1.5 * стандартное отклонение
		margin := stdDev * 1.5

		// Вычисляем нижнюю и верхнюю границу диапазона
		lower := mean - margin
		upper := mean + margin

		// Выводим диапазон с округлением до целого числа
		fmt.Printf("%.0f %.0f\n", lower, upper)
	}
}
