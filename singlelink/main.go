package main
import (
	"fmt"
)


type HeroNode struct {

	no int
	name string
	nikename string
	next *HeroNode
}


func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode

}


func ListHeroNode(){
	
}


func main(){
	hero := &HeroNode{}

	hero1 := &HeroNode{
		no:1,
		name:"宋江",
		nikename:"及时雨",
	}





}

















