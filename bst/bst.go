package bst

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

// add multiple values from user
func (t *BST) AddMultipleValues() {
	var length, value int

	fmt.Print("\nEnter how many values want to enter int BST: ")
	fmt.Scan(&length)

	for i := 1; i <= length; i++ {
		fmt.Print("Enter value: ")
		fmt.Scan(&value)
		t.AddValue(value)
	}
}

// inset values in bst
func (t *BST) AddValue(data int) {
	//check if the root is nill if it is nill set the root
	if t.Root == nil {
		t.Root = &Node{Data: data}
		return
	}

	currentNode := t.Root

	for {
		if data < currentNode.Data {
			//if left is nill assign value and return
			if currentNode.Left == nil {
				currentNode.Left = &Node{Data: data}
				return
			}
			currentNode = currentNode.Left
		} else { // if right is nill add return
			if currentNode.Right == nil {
				currentNode.Right = &Node{Data: data}
				return
			}
			currentNode = currentNode.Right
		}
	}
}

// print all values in the binary search tree in ascending order wa
func (t *BST) PrintValuesLRR() {
	if t.Root == nil {
		fmt.Println("There is no values in the BST")
		return
	}

	fmt.Println("Values in the BST ascending order")
	printValuesLRRHelper(t.Root)
	fmt.Println()
}

func printValuesLRRHelper(node *Node) {

	if node == nil {
		return
	}

	printValuesLRRHelper(node.Left)

	fmt.Print((*node).Data, " ")

	printValuesLRRHelper(node.Right)
}

// print all values in the binary search tree in descending order wa
func (t *BST) PrintValuesRRl() {
	if t.Root == nil {
		fmt.Println("There is no values in the BST")
		return
	}

	fmt.Println("Values in the BST descending order")
	printValuesRRLHelper(t.Root)
	fmt.Println()
}

func printValuesRRLHelper(node *Node) {

	if node == nil {
		return
	}

	printValuesRRLHelper(node.Right)

	fmt.Print((*node).Data, " ")

	printValuesRRLHelper(node.Left)
}
