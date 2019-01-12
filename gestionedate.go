package gestionedate

import (
"time"
)

func OttieniNumeroDiGiorniInUnMese(anno,mese int) int{
	d:=time.Date(anno,1,0,0,0,0,0,time.UTC)
	if(mese<12){
		m:=time.Month(mese+1)
		d=time.Date(anno,m,0,0,0,0,0,time.UTC)
	}else{
		a:=anno+1
		d=time.Date(a,1,0,0,0,0,0,time.UTC)
	}
	return d.Day()
}

func OttieniGiorniDellaSettimana(anno,mese,giorno int) int{
	d:=time.Date(anno,time.Month(mese),giorno,20,0,0,0,time.UTC)
	return int(d.Weekday())
}
