package date

import (
	"encoding/json"
	"os"

	"github.com/FalkSturmfels/dsacalender/data"
	"github.com/FalkSturmfels/dsacalender/pkg/filemgr"
)

type SaveDate struct {
	DayId   int `json:"dayId"`
	MonthId int `json:"monthId"`
	Year    int `json:"year"`
}

func SaveCurrentDate(cd Date) error {
	path, err := filemgr.CurrentDayPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(dateToSaveDate(cd), "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LoadCurrentDate() (Date, error) {
	path, err := filemgr.CurrentDayPath()
	if err != nil {
		return Date{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Date{}, nil
		}
		return Date{}, err
	}

	var sd SaveDate
	err = json.Unmarshal(data, &sd)
	return saveDateToDate(sd), err
}

func dateToSaveDate(date Date) SaveDate {
	return SaveDate{DayId: date.Day.Id, MonthId: date.Month.Id, Year: date.Year}
}

func saveDateToDate(sd SaveDate) Date {
	cd := Date{Day: data.Day{Id: sd.DayId}, Month: data.Month{Id: sd.MonthId}, Year: sd.Year}
	return cd
}
