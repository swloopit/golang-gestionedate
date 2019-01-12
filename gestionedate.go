package gestionedate

import (

"time"
)

//func main(){
//argomenti:=os.Args
/*
	fmt.Println("Hello, playground")
	for i:=1;i<=12;i++{
	
	fmt.Println(OttieniNumeroDiGiorniInUnMese(2019,i))
	for k:=1;k<=OttieniNumeroDiGiorniInUnMese(2019,i);k++{
	fmt.Println(i)
	fmt.Println(k)
	fmt.Println(OttieniGiorniDellaSettimana(2019,i,k))
	}
	}*/
//}

func OttieniNumeroDiGiorniInUnMese(anno,mese int) int{
d:=time.Date(anno,1,0,0,0,0,0,time.UTC)
if(mese<12){
m:=time.Month(mese+1)
d=time.Date(anno,m,0,0,0,0,0,time.UTC)
}else{
a:=anno+1
d=time.Date(a,1,0,0,0,0,0,time.UTC)
}
//fmt.Println(OttieniGiorniDellaSettimana(2019,1,12))
return d.Day()
}

func OttieniGiorniDellaSettimana(anno,mese,giorno int) int{
d:=time.Date(anno,time.Month(mese),giorno,20,0,0,0,time.UTC)
return int(d.Weekday())
}
