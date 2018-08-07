package main
import "time"
// import "fmt"
import "strconv"

type Date struct{
	year int
	month int
	day int
}

func (self *Date) set_from_string(date string) {

	self.year, _ = strconv.Atoi(date[0:4])
	self.month, _ = strconv.Atoi(date[5:7])
	self.day, _ = strconv.Atoi(date[8:10])
}

func (self *Date) is_different_date(other Date) bool {
	A := self.year == other.year
	B := self.month == other.month
	C := self.day == other.day

	if (A && B && C){
		return false	
	}
	return true
}

func (self *Date) is_different_week(other Date) bool {
	day := time.Date(self.year, time.Month(self.month), self.day, 0, 0, 0, 0, time.UTC)
	other_day := time.Date(other.year, time.Month(other.month), other.day, 0, 0, 0, 0, time.UTC)
	delta := other_day.Sub(day).Hours() / 24
	if (delta >= 7){
		return true
	}
	other_week_day := other_day.Weekday()-1
	week_day := day.Weekday()-1

	if (other_week_day == -1){
		other_week_day = 6
	}
	if (week_day == -1){
		week_day = 6
	}
	if (delta < 7 &&  other_week_day < week_day){
		return true
	}
	return false
}


// func main(){
// 	date := Date{}
// 	date.set_from_string("2018-08-06")

// 	fmt.Println(date.year)
// 	fmt.Println(date.month)
// 	fmt.Println(date.day)

// 	date2 := Date{2018, 8, 6}
// 	date3 := Date{2018, 8, 7}

// 	fmt.Println(date.is_different_date(date2))
// 	fmt.Println(date.is_different_date(date3))

// 	date4 := Date{2018, 8, 5}
// 	date5 := Date{2018, 8, 6}

// 	fmt.Println(date4.is_different_week(date5))
// }

