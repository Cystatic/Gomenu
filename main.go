package main

import "fmt"

func main() {
	var showmenu = menu()
	showmenu()
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
