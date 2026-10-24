package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	words := strings.Split(datastring, ",")

	if len(words) != 2 {
		return fmt.Errorf("длина слайса строк должна быть равна двум")
	}

	steps, err := strconv.Atoi(words[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования шагов: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(words[1])
	if err != nil {
		return fmt.Errorf("ошибка преобразования времени: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть больше нуля")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчёта калорий: %w", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
