package test

import (
	"my-life/model"
	"my-life/service"
	"testing"
)
func TestRest(t *testing.T){
   init:=&model.Init{}
   init.Age=27
   init.TargetAge=65
   days:=service.DealAge(init,"days")
   years:=service.DealAge(init,"years")
   t.Log(days,years)
    rest:=service.DealRest(years,days)
	t.Log("rest-time",rest)
}