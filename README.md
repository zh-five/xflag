# xflag 
一个go语言命令行参数解析库, 基于flag模块扩展了以下功能:
- 支持无限级别的子命令. 如: `app run say -s hello`
- option(`-`开头的选项和数据)其余参数无顺序要求, 如: `rm -r ./dir/ -f`

命令格式为: `app [cmd...] [option or param]`
`cmd` 为子命令,必须逐级按顺序填写
`option` 是以`-`开头的选项和数据(如果有),如, `-f` `-o a.log`
`param` 其余参数. 如`rm -r ./dir/ -f` 中的`./dir/`


# 安装
```base
go get github.com/zh-five/xflag
```

# 示例
```go
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

```

```
$ ./test -h
用法: ./test [cmd...] [option...] [param...]
Subcommand of ./test:
     clone       克隆资源
     add         添加文件内容到索引
     help        查看帮助信息
Usage of ./test:
  -h    查看帮助信息

作者: xxx

$ ./test clone -h
用法: ./test clone [-b 分支] repo [dir]
Usage of ./test clone:
  -b string
        分支名称

```
