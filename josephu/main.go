package main
import (
	"fmt"
)

type Boy struct {
	No int
	Next *Boy
}

func AddBoy(num int) *Boy {

	first := &Boy{}
	curBoy := &Boy{}
	
	if num < 1{
		fmt.Println("num的值不对")
		return first
	}

	for i :=1 ; i <= num; i++ {
		boy := &Boy{
			No:i,
		}

		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = boy
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}

	return first

}

func ShowBoy(first *Boy){
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩")
		return
	}

	curBoy := first

	for {
		fmt.Println("小孩编号为：",curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

func PlayGame(first *Boy, startNo int, countNum int){

	if first.Next == nil {
		fmt.Println("空的链表，。。。")
		return
	}

	tail := first

	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	for {
		 for i:=1;i <= countNum-1;i++ {
			first = first.Next
			tail = tail.Next
		 }

		 fmt.Println("小孩编号为",first.No,"出圈")
		 first = first.Next
		 tail.Next = first

		 if tail == first {
			break
		 }

	}
	fmt.Println("小孩编号为",first.No,"出圈")


}




func main(){

	first := AddBoy(5)
	ShowBoy(first)
	PlayGame(first,2,3)


}



