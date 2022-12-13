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

// inorder
// left -> root -> right
func (t *BST) Inorder() {
	if t.Root == nil {
		fmt.Println("There is no values in the BST")
		return
	}

	fmt.Println("\nInorder")
	inorderHelper(t.Root)
	fmt.Println()
}

func inorderHelper(node *Node) {

	if node == nil {
		return
	}

	inorderHelper(node.Left)

	fmt.Print((*node).Data, " ")

	inorderHelper(node.Right)
}

// pre-oder printing
// root -> left -> right
func (t *BST) PreOrder() {
	if t.Root == nil {
		fmt.Println("There is no values in the BST")
		return
	}

	fmt.Println("\nPre-order")
	preOrderHelper(t.Root)
	fmt.Println()
}

func preOrderHelper(node *Node) {

	if node == nil {
		return
	}

	fmt.Print((*node).Data, " ")
	preOrderHelper(node.Left)

	preOrderHelper(node.Right)
}

//post order
// left -> root -> right

func (t *BST) PostOder() {
	if t.Root == nil {
		fmt.Println("There is no values in the BST")
		return
	}
	fmt.Println("\nPost-order")
	postOderHelper(t.Root)
	fmt.Println()
}

func postOderHelper(node *Node) {

	if node == nil {
		return
	}

	postOderHelper(node.Left)

	postOderHelper(node.Right)

	fmt.Print((*node).Data, " ")
}
