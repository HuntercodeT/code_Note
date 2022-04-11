package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 通过标准输入建立二叉树
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	str := strings.Split(scanner.Text(), ",")
	var root *TreeNode
	list := make([]*TreeNode, len(str))
	value := 0
	// node := root
	for i := 0; i < len(str); i++ {
		var node *TreeNode
		if str[i] != "nil" {
			value, _ = strconv.Atoi(str[i])
			node = &TreeNode{Val: value}
		}
		list[i] = node
		if i == 0 {
			root = node
		}
	}
	for i := 0; i < len(str)/2; i++ {
		if list == nil {
			continue
		}
		if 2*i+1 < len(str) {
			list[i].Left = list[2*i+1]
		}
		if 2*i+2 < len(str) {
			list[i].Right = list[2*i+2]
		}
	}
	//  前序遍历二叉树
	Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Println(root *TreeNode) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		fmt.Println(node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
}
