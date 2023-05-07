package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func insert(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }

    if val < root.Val {
        root.Left = insert(root.Left, val)
    } else {
        root.Right = insert(root.Right, val)
    }

    return root
}

func search(root *TreeNode, val int) bool {
    if root == nil {
        return false
    }

    if root.Val == val {
        return true
    } else if val < root.Val {
        return search(root.Left, val)
    } else {
        return search(root.Right, val)
    }
}

func delete(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return root
    }

    if val < root.Val {
        root.Left = delete(root.Left, val)
    } else if val > root.Val {
        root.Right = delete(root.Right, val)
    } else {
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        }

        minNode := root.Right
        for minNode.Left != nil {
            minNode = minNode.Left
        }

        root.Val = minNode.Val
        root.Right = delete(root.Right, minNode.Val)
    }

    return root
}

func main() {
    root := &TreeNode{Val: 5}
    insert(root, 3)
    insert(root, 7)
    insert(root, 1)
    insert(root, 9)

    fmt.Println(search(root, 7))
    fmt.Println(search(root, 4))

    root = delete(root, 5)

    fmt.Println(search(root, 5))
}
