package main

import "fmt"

func printer (c chan int){
	for{
		data:=<-c
		if data==0{
			break
		}
		fmt.Print(data)
	}
	c<-0
}

func main ()  {
	c:=make(chan int)
go printer(c)
for i:=1;i<=5000;i++{
	c<-i
}
c<-0
<-c
}
