package data

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
)

type Day struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Short string `json:"short"`
}

type Month struct {
	Month      string `json:"month"`
	Id         int    `json:"id"`
	Real       string `json:"real"`
	FirstDayId int    `json:"firstDay"`
	LastDayId  int    `json:"lastDay"`
	Days       []Day  `json:"days"`
}

//go:embed months.json
var monthsfile []byte

var Months []Month
var MonthMap map[int]Month

func GetMonthByDay(dayId int) Month {
	for _, month := range MonthMap {
		if month.ContainsDay(dayId) {
			return month
		}
	}
	return Month{}
}

func (m Month) GetDay(dayId int) Day {
	for _, day := range m.Days {
		if day.Id == dayId {
			return day
		}
	}
	return Day{}
}

func (m Month) ContainsDay(dayId int) bool {
	return m.FirstDayId <= dayId && dayId <= m.LastDayId
}

func (m Month) FirstDay() Day {
	return m.Days[0]
}
func (m Month) SecondDay() Day {
	return m.Days[1]
}

func (m Month) LastDay() Day {
	return m.Days[len(m.Days)-1]
}

func (m Month) WeekStartDays() (res []int) {
	for _, day := range m.Days {
		if day.Short == "Untergang" {
			res = append(res, day.Id)
		}
	}
	return
}

func (m Month) IsEquinoxMonth() bool {
	for _, day := range m.Days {
		if strings.Contains(day.Name, "Tag und Nachtgleiche") {
			return true
		}
	}
	return false
}

func GetMonth(monthId int) Month {
	return MonthMap[monthId]
}

func init() {
	if err := json.Unmarshal(monthsfile, &Months); err != nil {
		fmt.Println("Cannot load months json")
	}

	MonthMap = make(map[int]Month)
	for _, month := range Months {
		MonthMap[month.Id] = month
	}
}
