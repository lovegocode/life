package service

import "time"
//获取当前时间
func NowTime(location string)any{
	now:=time.Now()
	locat,_:=time.LoadLocation(location)
	locatTime:=now.In(locat)
	return locatTime
}