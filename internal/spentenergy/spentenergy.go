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
	if steps <= 0 {
		return 0, fmt.Errorf("error: expect steps > 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("error: expect weight > 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("error: expect height > 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("error: expect duration > 0")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	walkCalories := (weight * meanSpeed * durationInMinutes) / minInH
	walkingSpentCalories := walkCalories * walkingCaloriesCoefficient
	return walkingSpentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("error: expect steps > 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("error: expect weight > 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("error: expect height > 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("error: expect duration > 0")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	runCalories := (weight * meanSpeed * durationInMinutes) / minInH
	return runCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	hours := duration.Hours()
	meanSpeed := distance / hours
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	distance := (float64(steps) * stepLength) / mInKm
	return distance
}
