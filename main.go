package main

import (
	"fmt"
)

func main() {
	// var showmenu = menu()
	// showmenu()

	linktable := initLinkTable()
	// fmt.Println(linktable.headNode)

	// insertHeadNode(1, linktable)
	// insertHeadNode(2, linktable)
	// insertHeadNode(3, linktable)
	// showLinkTableData(linktable)
	// fmt.Println()

	insertTailNode(4, linktable)
	insertTailNode(5, linktable)
	insertTailNode(6, linktable)
	// showLinkTableData(linktable)de
	// fmt.Println()
	// append(3, 7, linktable)
	// fmt.Println(findNodeByIndex(1, linktable))
	// fmt.Println(findNodeByValue(6, linktable))

	deleteNode(2, linktable)
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
