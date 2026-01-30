package test

import (
	"my-life/model"
	"my-life/service"
	"testing"
)
func TestWorkMoney(t *testing.T){
   init:=& model.Init{}
   init.Age=27
   init.HourlyRate=20
   init.TargetAge=70
   init.Worktime=8
   work:=service.NewWork(init)
   
   money:= work.Work()
   t.Log("\n",money)
}