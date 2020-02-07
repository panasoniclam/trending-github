package  main

import (
	"fmt"
	"sync"
	"time"
)

type (
	Dog struct {  //type difinition
		Id int
		Name string
		SaidNumber int
	}
	Cat struct {
		Id int
		Name string
		SaidNumber int
	}
	Animal interface {
		saidNumber() int
		saidName() string
	}
)

func (d *Dog) saidNumber() int {
	d.SaidNumber++
	return d.SaidNumber

}
func (c *Cat)saidNumber()int  {
	c.SaidNumber ++
	return  c.SaidNumber
}
func (d *Dog)saidName() string  {
	return  fmt.Sprintf("%d gau gau %s",d.SaidNumber,d.Name)
}
func (c *Cat)saidName()string  {
	return fmt.Sprintf("%d mew mew %s",c.SaidNumber,c.Name)
}
func sortEqual(a[]int) (count int ){
	var status []bool
	for i:=0;i<len(a);i++ {
		for j:=i;j<len(a)-1;j++{
				if (a[i] == a[j] && status[i] == false && status[j] == false && i!= j){
					status[i] = true
					status[j] = true
					count++
					break
			     }
		}
	}
	return
}
func PrintfWord(words []string,c chan string)  {
	defer close(c)
	for _,v := range words {
		time.Sleep(1*time.Second)
		c<-v
	}
}
func Printf(words []string, wg sync.WaitGroup)  {

	for _,v := range words {
		time.Sleep(time.Second)
		fmt.Println(v)
		defer wg.Done()
	}

}
type sock struct {
	number int32
	occurrences int32
}
func (pair *sock)increment() {
	pair.occurrences  ++
}
func (pair *sock)getAmountOfPair()int32  {
	return pair.occurrences/2
}
func sockMerchan(n int32,ar []int32) int32 {
	pairs := make([]*sock,0,10)
	for _, e := range ar {
		found := false
		for _, pair := range pairs {
			if pair.number == e {
				pair.increment()
				found = true
			}
		}

		if !found {
			pairs = append(pairs, &sock{
				number:      e,
				occurrences: 1,
			})
		}
	}
	var result int32
	for _, pair := range pairs {
		result += pair.getAmountOfPair()
	}
	return result
}
func main()  {
	arr := []int32{1,2,1,1,3,4,5,4}
	result := sockMerchan(2,arr)
	fmt.Println(result)
}