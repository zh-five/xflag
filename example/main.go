package main

import (
	"fmt"
	"os"

	"github.com/zh-five/xflag"
)

func main() {
	flag := xflag.Flag()

	flag.Desc("用法: "+flag.Name()+" [cmd...] [option...] [param...]", "\n作者: xxx")

	// 绑定子命令
	flag.BindCmd("clone", cmdClone, "克隆资源")
	flag.BindCmd("add", cmdAdd, "添加文件内容到索引")
	flag.BindCmd("help", nil, "查看帮助信息")

	help := flag.Bool("h", false, "查看帮助信息")
	flag.Parse()

	if *help || flag.CmdName == "" || flag.CmdName == "help" {
		flag.Usage()
	}
}

func cmdClone(flag *xflag.XFlagSet) {
	flag.Desc("用法: "+flag.Name()+" [-b 分支] repo [dir]", "")
	branch := flag.String("b", "", "分支名称")
	flag.Parse()

	repo := flag.Arg(0) //仓库
	dir := flag.Arg(1)  //目录

	if repo == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("branch: %s, repo: %s, dir: %s\n", *branch, repo, dir)
	fmt.Println("os.args", os.Args[1:])
}

func cmdAdd(flag *xflag.XFlagSet) {
	flag.Desc("添加文件到索引\n用法: "+flag.Name()+" [-f] <文件列表>", "")
	force := flag.Bool("f", false, "force, 允许添加忽略的文件")
	flag.Parse()

	files := flag.Args() //文件列表
	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("force:", *force, "files:", files)
}
