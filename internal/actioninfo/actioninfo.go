package actioninfo

import (
	"fmt"
	"log"
	// "github.com/Yandex-Practicum/tracker/internal/daysteps"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, step := range dataset {
		err := dp.Parse(step)
		if err != nil {
			log.Println("error parsing input value")
			continue
		}
		str, err := dp.ActionInfo()
		if err != nil {
			log.Println("error create activiti info")
			continue
		}
		fmt.Println(str)
	}
}
