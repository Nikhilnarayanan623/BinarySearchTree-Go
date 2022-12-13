package main

import (
	"binarySearchTree/bst"
	"fmt"
)

func main() {

	fmt.Println("Program to Print values in the BST in various order")

	bst1 := bst.BST{}

	bst1.AddMultipleValues()

	bst1.Inorder()

	bst1.PreOrder()

	bst1.PostOder()
}
