package service

import "time"

//获取当前时间
func Years(location string)time.Time{
	now:=time.Now()
	locat,_:=time.LoadLocation(location)
	locatTime:=now.In(locat)
	return locatTime
}
//计算月份
func Months(location string,hope int)int{
	time:=Years(location)
	 months:=int(time.Month())
	 month:=(hope-1)*12 +(12-months)
	return month
}
func Days(location string,hope int)int{
	time:=Years(location)
	furter:=time.AddDate(hope,0,0)
   days:= int(furter.Sub(time).Hours()/24)
   return days
}