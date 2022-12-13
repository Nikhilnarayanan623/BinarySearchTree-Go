package main

import (
	"binarySearchTree/bst"
	"fmt"
)

func main() {
	fmt.Println("Program to remove value from bst")

	bst1 := bst.BST{}

	bst1.AddMultipleValues()
	bst1.Inorder()

	var value int
	fmt.Print("Enter value to remove from BST: ")
	fmt.Scanf("%d", &value)

	removeValue(&bst1, value)
	fmt.Println("Values after removed")
	bst1.Inorder()

}

func removeValue(b *bst.BST, value int) {

	cuurentNode := b.Root

	if cuurentNode == nil {
		fmt.Println("There is no value in BST")
		return
	}

	for cuurentNode != nil {

		if value < cuurentNode.Data {
			//value found at current node left then assign here changed value
			if cuurentNode.Left != nil && value == cuurentNode.Left.Data {
				cuurentNode.Left = removeValueHelper(cuurentNode.Left)
				return
			}
			cuurentNode = cuurentNode.Left
		} else if value > cuurentNode.Data {
			//value found at current node right then assign here changed value
			if cuurentNode.Right != nil && value == cuurentNode.Right.Data {

				cuurentNode.Right = removeValueHelper(cuurentNode.Right)
				return
			}
			cuurentNode = cuurentNode.Right
		} else { // if root have value
			fmt.Println("At root")
			b.Root = removeValueHelper(b.Root)
			return
		}
	}
}

// right then left approach
func removeValueHelper(child *bst.Node) *bst.Node {
	if child.Right == nil { //right nill then simply return left
		return child.Left
	} else if child.Right.Left == nil { //if child.right.left is nil
		child.Right.Left = child.Left //connect child.left to its right.left
		return child.Right
	}

	temp := child.Right.Left
	innnerParent := child.Right //inner parent to conncet last nodes right to its parent left

	for temp.Left != nil {
		innnerParent = temp
		temp = temp.Left
	}

	innnerParent.Left = temp.Right
	child.Data = temp.Data

	return child

}
