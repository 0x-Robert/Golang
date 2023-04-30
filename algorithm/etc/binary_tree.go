package main

import (
	"fmt"
	"io"
	"os"
)

//바이너리 트리 구조체 
type BinaryTree struct {
    root *BinaryNode //이진트리의 루트노드를 필드로 가짐 
}

//바이너리 노드 구조체
type BinaryNode struct {
    left  *BinaryNode //왼쪽 자식 노드 
    right *BinaryNode //오른쪽 자식 노드 
    data  int64 //노드에 저장할 데이터 
}
 
 
//포인터 트리 구조체를 인자로 받는다
//BinaryTree 구조체에 insert 메서드를 정의한다.
func (t *BinaryTree) insert(data int64) *BinaryTree {
    if t.root == nil { 
        t.root = &BinaryNode{data: data, left: nil, right: nil}
    } else {
        t.root.insert(data)
    }
    return t
}
 
//포인터 바이너리노드를 인자로 받는다. 
//바이너리 노드 구조체에도 insert 메서드를 정의한다. 
func (n *BinaryNode) insert(data int64) {
    if n == nil {
        return
    } else if data <= n.data {
        if n.left == nil {
            n.left = &BinaryNode{data: data, left: nil, right: nil}
        } else {
            n.left.insert(data)
        }
    } else {
        if n.right == nil {
            n.right = &BinaryNode{data: data, left: nil, right: nil}
        } else {
            n.right.insert(data)
        }
    }   
}
//입력받은 노드를 루트로 하는 이진트리를 출력한다. 

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
    if node == nil {
        return
    }
 
    for i := 0; i < ns; i++ {
        fmt.Fprint(w, " ")
    }
    fmt.Fprintf(w, "%c:%v\n", ch, node.data)
    print(w, node.left, ns+2, 'L')
    print(w, node.right, ns+2, 'R')
}
 
func main() {
    tree := &BinaryTree{}
    tree.insert(100).
        insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
        insert(50).
		insert(60).
		insert(55).
        insert(85).
		insert(15).
		insert(5).
        insert(-10)
    print(os.Stdout, tree.root, 0, 'M')
}