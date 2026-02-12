package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration > 0 {
		if steps > 0 && weight > 0 && height > 0 {
			speed := MeanSpeed(steps, height, duration)
			calories := (weight * speed * duration.Minutes()) / minInH
			corrCalories := calories * walkingCaloriesCoefficient // корректируем полученные выше калории
			return corrCalories, nil
		}
		return 0, fmt.Errorf("need value > 0")
	}
	return 0, fmt.Errorf("need value duration > 0")
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration > 0 {
		if steps > 0 && weight > 0 && height > 0 {
			speed := MeanSpeed(steps, height, duration)
			calories := (weight * speed * duration.Minutes()) / minInH // получаем калории
			return calories, nil
		}
		return 0, fmt.Errorf("need value > 0")
	}
	return 0, fmt.Errorf("value duration need > 0")
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration > 0 && steps > 0 && height > 0 {
		dis := Distance(steps, height)
		speed := dis / duration.Hours() // получаем среднюю скорость в час
		return speed
	}
	return 0
}

func Distance(steps int, height float64) float64 {
	if steps > 0 && height > 0 {
		stepLenght := height * stepLengthCoefficient        // расчитываем длину шага
		distanceKm := (float64(steps) * stepLenght) / mInKm // сразу получаем расстояние в км
		return distanceKm
	}
	return 0
}
