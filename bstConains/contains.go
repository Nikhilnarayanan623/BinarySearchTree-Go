package main

import (
	"binarySearchTree/bst"
	"fmt"
)

func main() {

	bst1 := bst.BST{}

	bst1.AddMultipleValues()

	var value int
	fmt.Print("Enter a value to check the vlaue conatins or not \nEnter vlaue: ")
	fmt.Scanf("%d", &value)

	fmt.Println(contains(&bst1, value))

}

func contains(b *bst.BST, value int) bool {

	currentNode := b.Root

	for currentNode != nil {
		if value < (*currentNode).Data { // if value less than cuurent node's value then move to left
			currentNode = currentNode.Left
		} else if value > currentNode.Data { //if value greater than cuurent node's value then move to right
			currentNode = currentNode.Right
		} else { // value equal to the node's value means conatains so return true
			return true
		}
	}
	//completing the loop means same value not found in the loop
	return false
}
