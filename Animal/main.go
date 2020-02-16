package  main

import (
	"fmt"
	"strings"
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
	oncurrences int32
}

func (pair *sock)increment()  {
	pair.oncurrences ++
}
func (pair *sock)getAmountOfPair()int32  {
	return pair.oncurrences/2
}
func sockMerchan(n int32,arr []int32) int32 {
	pairs := make([]*sock,0,10)
	for _,i := range arr {
		found := false
		for _,j :=range pairs {
			if j.number == i {
				j.increment()
				found = true
			}
		}
		if !found {
			pairs = append(pairs,&sock{
				number:      i,
				oncurrences: 1,
			})
		}
	}
	var result int32  = 0
	for _, a := range  pairs {
		result += a.getAmountOfPair()
	}
	return result
}
func valeyCount()  {
	s := "UDDUDUU"
	var countingValey int32 = 0
	var antitude int32 = 0
	for i:=0;i<len(s);i++ {
		if string(s[i]) == "U"{
			antitude ++
			if antitude == 0 {
				countingValey ++
			}
		}else  {
			antitude --
		}
	}

	fmt.Println(countingValey)
}
func Clount(){
	s:= "000010"
	c:= make([]int32,len(s))
	var currentPossition int32  = 0
	lastCloud := int32(len(s)) -1
	var jumb int32  = 0
	for currentPossition < lastCloud {
		if currentPossition +1 ==  lastCloud {
			currentPossition ++
		}else if c[currentPossition+2]==0  {
			currentPossition = currentPossition +2
		}else {
			currentPossition ++
		}
		jumb ++
	}
	fmt.Println(jumb)
}
func main()  {
     s:= "abc"
     s1 := strings.Trim(s,"c")

     s2:= strings.TrimRight(s,"c")
     fmt.Println(s1)
     fmt.Println(s2)


}