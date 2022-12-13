package main

import (
	"fmt"
)

func main() {

	vbst := node{data: 20}

	/*
		valid bst
			 20
			/  \
		   10   30
		   /\
		  5  15


	*/

	vbst.left = &node{data: 10}
	vbst.right = &node{data: 30}

	vbst.left.left = &node{data: 5}
	vbst.left.right = &node{data: 15}

	fmt.Println("vbst", validBST(&vbst, nil, nil))

	/*
		non valid bst

				 10
				/  \
		 	  5 	20
		 	 /
		 	15
	*/
	nbst := node{data: 10}
	nbst.left = &node{data: 5}
	nbst.right = &node{data: 20}
	nbst.left.left = &node{data: 15}

	fmt.Println("nbst ", validBST(&nbst, nil, nil))

}

type node struct {
	data  int
	left  *node
	right *node
}

func validBST(root *node, parent *node, gParent *node) bool {

	if root == nil {
		return true
	}

	if parent != nil && (*root).data <= (*parent).data {
		return false
	}

	if gParent != nil && (*root).data >= (gParent).data {
		return false
	}

	valid1 := validBST(root.left, parent, root)

	valid2 := validBST(root.right, root, gParent)

	return valid1 && valid2
}
