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

	steps, err := strconv.Atoi(dataSlices[0])
	if err != nil {
		return 0, "", 0, err
	}

	duration, err := time.ParseDuration(dataSlices[2])
	if err != nil {
		return 0, "", 0, err
	}

	return steps, dataSlices[1], duration, nil
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
	if duration < 0 {
		return 0
	}

	distanceKm, err := distance(steps, height)
	if err != nil {
		return 0
	}

	avarageSpeed := distanceKm / duration.Hours()

	return avarageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, typeOfActivity, duration, err := parseTraining(trainings)
	if err != nil {
		log.Println("Ошибка при парсинге данных:", err)
		return "", err
	}

	switch typeOfActivity {
	case "Бег":
		distanceR, err := distance(steps, height)
		if err != nil {
			return "", err
		}
		avarageSpeedR := meanSpeed(steps, height, duration)
		if err != nil {
			return "", err
		}
		caloriesR := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			typeOfActivity, duration, distanceR, avarageSpeedR, caloriesR)

		return result, err

	case "Ходьба":
		distanceW, err := distance(steps, height)
		if err != nil {
			return "", err
		}
		avarageSpeedW := meanSpeed(steps, height, duration)
		if err != nil {
			return "", err
		}
		caloriesW := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return errors.New("неизвестный тип активности")
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		typeOfActivity, duration, distanceW, avarageSpeedW, caloriesW)

	return result, err

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

	avarageSpeed, err := meanSpeed(steps, height, duration)
	if err != nil {
		return 0, err
	}

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

	avarageSpeed, err := meanSpeed(steps, height, duration)
	if err != nil {
		return 0, err
	}

	calories := (weight * float64(avarageSpeed) * float64(duration.Minutes()) / minInH)

	return calories * walkingCaloriesCoefficient, nil
}
