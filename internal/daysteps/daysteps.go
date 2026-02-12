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

func (ds *DaySteps) Parse(datastring string) (err error) {
	sliceData := strings.Split(datastring, ",") // разделяем строку, на слайс строк
	if len(sliceData) != 2 {
		return fmt.Errorf("data string does not contain 3 items")
	}

	ds.Steps, err = strconv.Atoi(sliceData[0]) // форматируем значение шагов в число и сразу добавляем в структуру
	if err != nil {
		return fmt.Errorf("steps is not an integer")
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("value steps is zero")
	}

	ds.Duration, err = time.ParseDuration(sliceData[1]) // аналогично шагам
	if err != nil {
		return fmt.Errorf("duration is not an integer")
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("value duration need > 0")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	resultString := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)
	return resultString, nil
}
