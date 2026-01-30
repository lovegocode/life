package service

import (
	"my-life/model"
	
)
type WorkMoney struct{
	init *model.Init
}
func NewWork( i*model.Init)*WorkMoney{
	work:=&WorkMoney{
		init:i,
	}
	return work
}
func (w*WorkMoney)Work()int{
   years:=w.dealAge("years")
   days:=w.dealAge("days")
   time:=DealRest(years,days)
   daymoney:=w.init.Worktime*w.init.HourlyRate
   leftOverMoney:=daymoney*time
   return leftOverMoney
}