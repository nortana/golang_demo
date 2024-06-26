package main
import (
	"fmt"
	"os"
	"errors"
)

type Queue struct {
	maxSize int
	array [5]int
	front int
	rear int
}

func (this *Queue) AddQueue(val int) (err error) {

	if this.rear == this.maxSize - 1 {
		return errors.New("queue full")
	}

	this.rear++

	this.array[this.rear] = val
	return
}


func (this *Queue) GetQueue() (val int, err error) {
	
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err

}


func (this *Queue) ShowQueue(){
	fmt.Println("显示当前的情况是：")
	for i := this.front + 1; i <=this.rear; i++ {
		fmt.Printf("array[%d]=%d\t",i,this.array[i])
	}
	fmt.Println()
}


func main(){
	queue := &Queue{
		maxSize:5,
		front:-1,
		rear:-1,
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
				err := queue.AddQueue(val)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("输入成功")
				}
				
			case "get":
				val, err := queue.GetQueue()
				if err != nil {
					fmt.Println(err.Error)
				} else {
					fmt.Println("从队列中取出一个数="，val)
				}
			case "show":
				queue.ShowQueue()
			case "exit":
				os.Exit(0)

		}
	}
}







