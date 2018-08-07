package main
import "fmt"

type Client struct{
	client_id int
	max_daily_load float32
	max_weekly_load float32
	max_num_load_daily int
	num_loads_currentday int
	last_load_date Date
	loadamount_currentday float32
	loadamount_currentweek float32
}

func (self *Client) init(client_ID int){
	self.client_id = client_ID
	self.max_daily_load = 5000
	self.max_weekly_load = 20000
	self.max_num_load_daily = 3
	self.num_loads_currentday = 0
	self.last_load_date = Date{}
	self.loadamount_currentday = 0
	self.loadamount_currentweek = 0
}

func (self *Client) initiate_new_day_load(load_date string) {
	date := Date{}
	date.set_from_string(load_date)
	if(self.last_load_date.is_different_date(date)){
		self.loadamount_currentday = 0
		self.num_loads_currentday = 0
		if(self.last_load_date.is_different_week(date)){
			self.loadamount_currentweek = 0
		}
		self.last_load_date = date
	}	
}

func (self *Client) check_daily_load_exceeded(load_amount float32) bool{
	if(load_amount > self.max_daily_load){
		return true
	}
	if(self.max_daily_load - self.loadamount_currentday < load_amount){
		return true
	}
	if(self.num_loads_currentday >= self.max_num_load_daily){
		return true
	}
	return false
}

func (self *Client) update_daily_load(load_amount float32){
	self.loadamount_currentday += load_amount
	self.num_loads_currentday += 1
}

func (self *Client) check_weekly_load_exceeded(load_amount float32) bool{
	if(load_amount > self.max_weekly_load){
		return true
	}
	if (self.max_weekly_load - self.loadamount_currentweek < load_amount){
		return true
	}
	return false
}

func (self *Client) update_weekly_load(load_amount float32){
	self.loadamount_currentweek += load_amount
}

func (self *Client) load(load_amount float32, load_date string) bool{
	self.initiate_new_day_load(load_date)
	A := self.check_daily_load_exceeded(load_amount)
	B := self.check_weekly_load_exceeded(load_amount)

	if(!A && !B){
		self.update_daily_load(load_amount)
		self.update_weekly_load(load_amount)
		return true	
	}
	return false
}



func main(){
	client := Client{}
	client.init(100)
	client.initiate_new_day_load("2018-08-06")
	fmt.Println(client.loadamount_currentday)
	fmt.Println(client.last_load_date)


	fmt.Println(client.check_daily_load_exceeded(6000))
	fmt.Println(client.check_daily_load_exceeded(5000))


	client.update_daily_load(2000)
	fmt.Println(client.loadamount_currentday)
	fmt.Println(client.num_loads_currentday)


	fmt.Println(client.check_weekly_load_exceeded(21000))
	fmt.Println(client.check_weekly_load_exceeded(19000))

	fmt.Println(client.load(5000, "2018-08-01"))
	fmt.Println(client.load(5000, "2018-08-02"))
	fmt.Println(client.load(5000, "2018-08-03"))
	fmt.Println(client.load(5000, "2018-08-04"))
	fmt.Println(client.load(5000, "2018-08-05"))
	fmt.Println(client.load(5000, "2018-08-06"))


}



