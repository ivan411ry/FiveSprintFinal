package daysteps

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("error: wrong string format %s", datastring)
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("error steps parsing: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("error: expected steps > 0, got %d", steps)
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("error time parsing: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("error: expected duration > 0, got %s", parts [1])
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
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	)
	return result, nil
}
