package main
import (
	"fmt"
)

type Emp struct {
	Id int
	Name string
	Next *Emp
}

type EmpLink struct {
	Head *Emp
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *EmpLink) Insert(emp *Emp){

	cur := this.Head
	var pre *Emp = nil

	if cur == nil {
		this.Head = emp
		return
	}

	for {
		if cur != nil {
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
 
	pre.Next = emp
	emp.Next = cur
 
}

func (this *EmpLink) ShowLink(no int){
	if this.Head == nil {
		fmt.Printf("链表%d 为空\n",no)
		return
	}

	cur := this.Head

	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id = %d 名字 = %s ->",no,cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}

	}
	fmt.Println()
}


func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if  cur == nil {
			break
		}
		cur = cur.Next
	}

	return nil
}

func (this *HashTable) Insert(emp *Emp){

	linkNo := this.HashFun(emp.Id)

	this.LinkArr[linkNo].Insert(emp)

}


func (this *HashTable) HashFun(id int) int {
	return id % 7
}


func (this *HashTable) ShowAll(){
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}


func (this *HashTable) FindById(id int) *Emp {
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

func (this *Emp) ShowMe(){
	fmt.Printf("链表%d 找到该雇员 %d",this.Id %7,this.Id,)
}


func main(){

	key := ""
	id := 0
	name := ""

	var hashTable HashTable

	for {
		fmt.Println("-----------雇员菜单----------------")
		fmt.Println("-----------input 添加雇员----------------")
		fmt.Println("-----------show 显示雇员----------------")
		fmt.Println("-----------find 查找雇员----------------")
		fmt.Println("-----------exit 退出----------------")
		fmt.Println("请输入选择：")
		fmt.Scanln(&key)

		switch key {
			case "input":
				fmt.Println("输入雇员id：")
				fmt.Scanln(&id)
				fmt.Println("输入雇员name：")
				fmt.Scanln(&name)
				emp := &Emp{
					Id:id,
					Name:name,
				}
				hashTable.Insert(emp)
			case "show":
				hashTable.ShowAll()
			case "find":
				fmt.Println("输入雇员id：")
				fmt.Scanln(&id)
				emp := hashTable.FindById(id)
				if emp == nil {
					fmt.Printf("%d 该雇员不存在\n",id)
				} else {
					emp.ShowMe()
				}
			case "exit":
			default:
				fmt.Println("输入有误")
		}
	}

}









