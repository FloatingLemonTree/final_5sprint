package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	words := strings.Split(datastring, ",")

	if len(words) != 3 {
		return fmt.Errorf("длина слайса строк должна быть равна трём")
	}

	steps, err := strconv.Atoi(words[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования шагов: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(words[2])
	if err != nil {
		return fmt.Errorf("ошибка преобразования времени: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть больше нуля")
	}

	t.Steps = steps
	t.TrainingType = words[1]
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	var (
		calories float64
		err      error
	)

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", fmt.Errorf("ошибка расчёта калорий: %w", err)
	}

	return fmt.Sprintf("Тип тренировки: %v\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, avgSpeed, calories), nil
}
