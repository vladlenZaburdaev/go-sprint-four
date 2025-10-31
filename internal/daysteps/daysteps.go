package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	dataSlices := strings.Split(data, ",")

	if len(dataSlices) != 2 {
		return 0, 0, fmt.Errorf("ожидалось два элемента, получено %d", len(dataSlices))
	}

	steps, err := strconv.Atoi(dataSlices[0])
	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов не может быть отрицательным")
	}

	duration, err := time.ParseDuration(dataSlices[1])
	if err != nil {
		return 0, 0, err
	}

	if duration < 1 {
		return 0, 0, err
	}

	return steps, duration, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Printf("Произошла ошибка: %v", err)
		return ""
	}

	if steps <= 0 {
		log.Printf("Некорректное количество шагов: %d", steps)
	}

	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Printf("Ошибка расчета калорий: %v", err)
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)

	return result

}
