package trainings
import (
"strings"
"time"
"fmt"
"strconv"
"github.com/Yandex-Practicum/tracker/internal/personaldata"
"github.com/Yandex-Practicum/tracker/internal/spentenergy"
) 

type Training struct {
Steps int
TrainingType string
Duration time.Duration
Personal personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
parts := strings.Split(datastring, ",")
if len(parts) != 3 {
	return fmt.Errorf("error: len(parts) < 3")
}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("parsing steps error %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("error: expected steps > 0, got %d", steps)
	}
	t.Steps = steps
	t.TrainingType = parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("parsing duration error %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("error: expected duration > 0, got %v", duration)
	}
	t.Duration = duration
	return nil
}


func (t Training) ActionInfo() (string, error) {
distance := spentenergy.Distance(t.Steps, t.Personal.Height)
meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
var calories float64
var err error
switch t.TrainingType {
case "Бег":
	calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
case "Ходьба":
	calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
default:
	return "", fmt.Errorf("error: unknown training type: %s", t.TrainingType)
	}
	if err != nil {
		return "", err
	}
	durationHours := t.Duration.Hours()
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", 
        t.TrainingType, durationHours, distance, meanSpeed, calories)
		return result, nil
}
func (t Training) Print () {
	t.Personal.Print()
}