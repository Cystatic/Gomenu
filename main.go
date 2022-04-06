package main

import (
	"fmt"
)

func main() {
	// var showmenu = menu()
	// showmenu()
	linktable := initLinkTable()

	insertTailNode(1, linktable)
	insertTailNode(2, linktable)
	insertTailNode(3, linktable)
	insertTailNode(4, linktable)
	insertTailNode(5, linktable)
	insertTailNode(6, linktable)

	fmt.Print("初始链表：")
	showLinkTableData(linktable)

	fmt.Print("在索引为3处插入值为7的节点：")
	append(3, 7, linktable)
	showLinkTableData(linktable)

	fmt.Print("查找索引为1的节点：")
	fmt.Println(findNodeByIndex(1, linktable))

	fmt.Print("查找值为6的节点：")
	fmt.Println(findNodeByValue(6, linktable))

	fmt.Print("删除第6个节点：")
	deleteNode(6, linktable)
	showLinkTableData(linktable)
}
func menu() func() {
	var message string = ""
	fmt.Print("请输入指令：h-help,v-version,q-quit\n")
	fun := func() {
		for {
			fmt.Print(">>")
			fmt.Scan(&message)
			switch message {
			case "h":
				fmt.Println("help")
			case "v":
				fmt.Println("version-1.0")
			case "q":
				return
			default:
				fmt.Println("error command")
			}
		}
	}
	return fun
}
