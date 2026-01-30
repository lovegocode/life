package service



func(w*WorkMoney) dealAge(unit string)int{
   Hope:=w.init.TargetAge-w.init.Age
    locat:="Asia/Shanghai"
      
   switch unit {
   case "years":
      return Hope
   case "months":
      now:=Months(locat,Hope)
     return now
   case "days":
      now:=Days(locat,Hope)
      return now
   }
   return -1
}