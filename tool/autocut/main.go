package main

import (
	"flag"
	"log"
	"os"
	"regexp"
	"strings"
)

// main 修改 md 文件，去除无声片段
func main() {
	var f string
	flag.StringVar(&f, "f", "", "")
	flag.Parse()
	if f == "" {
		log.Fatalln("filename can't empty")
	}
	content, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln("read file error", err)
	}

	reg, err := regexp.Compile(`- \[ \] \[.*\].*`)
	if err != nil {
		log.Fatalln("regexp compile error", err)
	}
	subs := reg.FindAllString(string(content), -1)
	// 前两个元素是为了美观
	// 第三个元素是固定开头表明编辑完成
	selected := []string{"# edit", "", "- [x] <-- Mark if you are done editing."}
	for _, s := range subs {
		// 只选择有声片段
		if !strings.Contains(s, "No Speech") {
			selected = append(selected, strings.ReplaceAll(s, "[ ]", "[x]"))
		} else {
			selected = append(selected, s)
		}
	}

	fd, err := os.Create(f)
	if err != nil {
		log.Fatalln("write file error", err)
	}
	for _, s := range selected {
		fd.WriteString(s + "\n")
	}
}
