package main

import (
	"binarySearchTree/bst"
	"fmt"
)

func main() {
	fmt.Println("Program to find the closest value near given number in BST")

	bst1 := bst.BST{}

	bst1.AddMultipleValues()

	bst1.Inorder()

	var checkVal int

	fmt.Print("Enter a number to check closest value: ")
	fmt.Scanf("%d", &checkVal)

	findClostest(&bst1, checkVal)
}

func findClostest(b *bst.BST, value int) {

	currentNode := b.Root

	if currentNode == nil {
		fmt.Println("Threr is no value in the BST")
		return
	}
	var closeVal = (*currentNode).Data

	for currentNode != nil {

		if value < (*currentNode).Data {

			if value > closeVal && (*currentNode).Data-value < value-closeVal {

				closeVal = (*currentNode).Data
			} else if (*currentNode).Data-value < closeVal-value {

				closeVal = (*currentNode).Data
			}

			// if currentNode.Left == nil {
			// 	fmt.Println(checkClose(value, closeVal, (*currentNode).Data))
			// 	return
			// }
			//closeVal = checkClose(value, closeVal, (*currentNode).Data)

			currentNode = currentNode.Left

		} else if value > (*currentNode).Data {

			if value > closeVal && value-(*currentNode).Data < value-closeVal {

				closeVal = (*currentNode).Data
			} else if value-(*currentNode).Data < closeVal-value {

				closeVal = (*currentNode).Data
			}
			// if currentNode.Right == nil {
			// 	fmt.Println(checkClose(value, closeVal, (*currentNode).Data))
			// 	return
			// }
			// closeVal = checkClose(value, closeVal, (*currentNode).Data)

			currentNode = currentNode.Right

		} else {
			break
		}

	}

	fmt.Println("Closest value in BST is ", closeVal)
}

/*
func checkClose(data int, value1 int, value2 int) int {

	var res1, res2 int

	if value1 < data {
		res1 = data - value1
	} else {
		res1 = value1 - data
	}

	if value2 < data {
		res2 = data - value2
	} else {
		res2 = value2 - data
	}

	if res1 < res2 {
		return value1
	} else {
		return value2
	}
}
*/
