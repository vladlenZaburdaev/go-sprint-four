package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	dataSlices := strings.Split(data, ",")

	if len(dataSlices) != 3 {
		return 0, "", 0, fmt.Errorf("ожидалось три элемента, получено %d", len(dataSlices))
	}

	stepsStr := strings.TrimSpace(dataSlices[0])
	if stepsStr == "" {
		return 0, "", 0, errors.New("количество шагов не указано")
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, "", 0, err
	}

	if steps == 0 {
		return 0, "", 0, errors.New("количество шагов должно быть положительным")
	}

	if steps < 0 {
		return 0, "", 0, errors.New("количество шагов должно быть положительным")
	}

	typeOfActivity := strings.TrimSpace(dataSlices[1])
	if typeOfActivity == "" {
		return 0, "", 0, errors.New("тип активности не указан")
	}

	if typeOfActivity != "Бег" && typeOfActivity != "Ходьба" {
		return 0, "", 0, errors.New("неизвестный вид тренировки")
	}

	durationStr := strings.TrimSpace(dataSlices[2])
	durationStr = strings.ReplaceAll(durationStr, "h0", "h")

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, "", 0, err
	}

	if duration == 0 {
		return 0, "", 0, errors.New("продолжительность должна быть положительной")
	}

	if duration < 0 {
		return 0, "", 0, errors.New("продолжительность должна быть положительной")
	}

	return steps, typeOfActivity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / float64(mInKm)

	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	distanceKm := distance(steps, height)

	avarageSpeed := distanceKm / duration.Hours()

	return avarageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, typeOfActivity, duration, err := parseTraining(data)
	if err != nil {
		log.Println("Ошибка при парсинге данных:", err)
		return "", err
	}

	switch typeOfActivity {
	case "Бег":
		distanceR := distance(steps, height)

		avarageSpeedR := meanSpeed(steps, height, duration)

		caloriesR, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeOfActivity, duration.Hours(), distanceR, avarageSpeedR, caloriesR)

		return result, err

	case "Ходьба":
		distanceW := distance(steps, height)

		avarageSpeedW := meanSpeed(steps, height, duration)

		caloriesW, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeOfActivity, duration.Hours(), distanceW, avarageSpeedW, caloriesW)

		return result, err

	default:
		return "", errors.New("неизвестный тип активности")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 0 {
		return 0, errors.New("количество шагов должно быть больше или равно 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше 0")
	}

	avarageSpeed := meanSpeed(steps, height, duration)

	calories := (weight * float64(avarageSpeed) * float64(duration.Minutes()) / minInH)

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 0 {
		return 0, errors.New("количество шагов должно быть больше или равно 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше 0")
	}

	avarageSpeed := meanSpeed(steps, height, duration)

	calories := (weight * float64(avarageSpeed) * float64(duration.Minutes()) / minInH)

	return calories * walkingCaloriesCoefficient, nil
}
