package test

import (
	"my-life/service"
	"testing"
)
func TestNowTime(t *testing.T){
   time:=service.Years("Asia/Beijing")
   t.Log("\n",time)
}