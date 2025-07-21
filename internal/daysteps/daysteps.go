package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("error: wrong string format %s", datastring)
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("error steps parsing: %w", err)
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("error time parsing: %w", err)
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps, distance, calories,
	)
	return result, nil
}
