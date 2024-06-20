package main
import (
	"fmt"
	"errors"
	"os"
)

type CircleQueue struct {
	maxSize int
	array [5]int
	head int
	tail int
}

func (this *CircleQueue) Push(val int) (err error) {
	if this.IsPull() {
		return errors.New("queue full")
	}

	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return -1, errors.New("queue empty")
	}
	val = this.array[this.head]
	this.head= (this.head + 1) % this.maxSize
	return

}

func (this *CircleQueue) ListQueue()   {

	size := this.Size()
	//fmt.Println("size",size)
	if size == 0 {
		fmt.Println("队列为空")
		return
	}

	tempHead := this.head
	//fmt.Println("tempHead",tempHead)
	for i := 0; i < size; i++ {
		//fmt.Println("this.array",this.array)
		fmt.Printf("arr[%d] = %d\t",tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}

	fmt.Println()

}

func (this *CircleQueue) IsPull() bool {
	return (this.tail + 1)% this.maxSize == this.head
}


func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}


func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}



func main(){
	queue := &CircleQueue{
		maxSize:5,
		head:0,
		tail:0,
	}

	var key string
	var val int
	for {
		fmt.Println("1.输入队列")
		fmt.Println("2.获取队列")
		fmt.Println("3.显示队列")
		fmt.Println("4.退出队列")

		fmt.Scanln(&key)

		switch key {
			case "add":
				fmt.Println("请输入队列：：")
				fmt.Scanln(&val)
				err := queue.Push(val)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("输入成功")
				}
				
			case "get":
				val, err := queue.Pop()
				if err != nil {
					fmt.Println(err.Error)
				} else {
					fmt.Println("从队列中取出一个数=", val)
				}
			case "show":
				queue.ListQueue()
			case "exit":
				os.Exit(0)

		}
	}
}












