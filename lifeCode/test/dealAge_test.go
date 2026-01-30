package test

import (
	"my-life/model"
	"my-life/service"
	"testing"
)
func TestAge(t *testing.T){
   init:=&model.Init{}
   init.Age=55
   init.TargetAge=80
  age:= service.DealAge(init,"days")
  t.Log(age)
}