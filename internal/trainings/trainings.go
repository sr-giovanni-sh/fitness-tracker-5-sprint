package trainings

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"

	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	sliceData := strings.Split(datastring, ",") // разделяем строку, на слайс строк
	if len(sliceData) != 3 {

		return fmt.Errorf("data string does not contain 3 items")
	}

	t.Steps, err = strconv.Atoi(sliceData[0]) // форматируем значение шагов в число и сразу добавляем в структуру
	if err != nil {
		return fmt.Errorf("steps is not an integer")
	}
	if t.Steps <= 0 {
		return fmt.Errorf("steps cannot be negative")
	}

	t.TrainingType = sliceData[1] // сразу добавляем пришедший тип тренировки

	t.Duration, err = time.ParseDuration(sliceData[2]) // аналогично шагам
	if err != nil {
		return fmt.Errorf("duration is not an integer")
	}
	if t.Duration <= 0 {
		return fmt.Errorf("duration cannot be negative")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	dis := spentenergy.Distance(t.Steps, t.Height)
	spd := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	var calories float64
	var err error

	if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}

	if err != nil {
		return "", err
	}

	resultString := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		float64(t.Duration)/float64(time.Hour),
		dis,
		spd,
		calories,
	)
	return resultString, nil
}
