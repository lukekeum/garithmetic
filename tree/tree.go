package tree

import (
	"fmt"

	"github.com/lukekeum/garithmetic/store"
)

type Tree struct {
	Node *Node
}

type Node struct {
	TokenType store.TokenType
	value     int
	left      *Node
	right     *Node
}

func (t *Tree) Insert(token store.TokenType, value int) *Tree {
	if t.Node == nil {
		t.Node = &Node{TokenType: token, value: value}
	} else {
		t.Node.Insert(token, value)
	}

	return t
}

func PrintNode(node *Node) {

	if node == nil {
		return
	}

	PrintNode(node.left)
	if node.TokenType == store.CONST {
		fmt.Println(fmt.Sprintf("%s %d", store.GetTokenKeyFromInteger(node.TokenType), node.value))
	} else {
		fmt.Println(fmt.Sprintf("%s %s", store.GetTokenKeyFromInteger(node.TokenType), fmt.Sprintf("%c", node.value)))
	}
	PrintNode(node.right)
}

func (n *Node) Insert(token store.TokenType, value int) *Node {
	if n.left == nil {
		n.left = &Node{TokenType: token, value: value}
	} else {
		n.right = &Node{TokenType: token, value: value}
	}

	return n
}
