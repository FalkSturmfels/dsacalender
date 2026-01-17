package date

import (
	"fmt"

	"github.com/FalkSturmfels/dsacalender/data"
)

type Date struct {
	Day   data.Day
	Month data.Month
	Year  int
}

var CurrentDate Date

func SetCurrentDate(dayId int, year int) {
	month := data.GetMonthByDay(dayId)
	day := month.GetDay(dayId)
	CurrentDate = Date{Day: day, Month: month, Year: year}
}

func PlusDays(offset int) Date {
	newDayId := CurrentDate.Day.Id + offset
	year := CurrentDate.Year

	if newDayId > 365 {
		newDayId = (newDayId % 366) + 1 // ensure that newDayId is between 1 and 365
		year++
	}

	return getNewDay(newDayId, year)
}

func MinusDays(offset int) Date {
	newDayId := CurrentDate.Day.Id - offset
	year := CurrentDate.Year

	if newDayId < 1 {
		newDayId = (newDayId % 366) + 1 // ensure that newDayId is between 1 and 365
		year--
	}
	return getNewDay(newDayId, year)
}

func getNewDay(newDayId int, year int) Date {
	month := data.GetMonth(CurrentDate.Month.Id)
	var newDay data.Day
	var newMonth data.Month

	if month.ContainsDay(newDayId) {
		newDay = month.GetDay(newDayId)
		newMonth = month
	} else {
		for _, mon := range data.MonthMap {
			if mon.ContainsDay(newDayId) {
				newDay = mon.GetDay(newDayId)
				newMonth = mon
				break
			}
		}
	}

	return Date{Day: newDay, Month: newMonth, Year: year}
}

func enrichtCurrentDate() {
	CurrentDate.Month = data.GetMonth(CurrentDate.Month.Id)
	day := CurrentDate.Month.GetDay(CurrentDate.Day.Id)
	if day.Name == "" {
		fmt.Println("Error: CurrentDate corrupted!")
	}
	CurrentDate.Day = day
}

func (d Date) ToString() string {
	return fmt.Sprintf("%s %s %d (Id: %d)", d.Day.Name, d.Month.Month, d.Year, d.Day.Id)
}

func init() {
	cd, err := LoadCurrentDate()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	CurrentDate = cd
	enrichtCurrentDate()
}
