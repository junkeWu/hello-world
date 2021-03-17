package main

import (
    "fmt"
	"os"
	"flag"
	"strings"
)

func main() {
	// 
	var p string
	flag.StringVar(&p,"type","","type's value can't null")
    flag.Parse()

    for _, a := range os.Args[1:] {
		var b []string
		b = strings.Split(a,"=")
		switch b[1]{
		case "1":
			fmt.Println("最是一年春好处，绝胜烟柳满皇都。")
		case "2":
			fmt.Println("自春来、惨绿愁红，芳心是事可可。")
		default:
			fmt.Println("使用type参数时请输入1或者2")
		}
	}

}
		