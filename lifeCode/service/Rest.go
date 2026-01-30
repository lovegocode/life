package service 
func DealRest(years int,days int)int{
	restday:=13*years+(52*years)
	leftOver:=days-restday
	return leftOver
}