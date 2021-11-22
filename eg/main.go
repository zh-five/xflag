package main

import (
	"fmt"

	"github.com/zh-five/xflag"
)

func main() {
	xflag.SubCmd("get", "获取资源", cmdGet)
	xflag.SubCmd("run", "运行", cmdRun)
	xflag.SubCmd("build", "编译", cmdBuild)
	xflag.Parse()

}

func cmdGet(flag *xflag.XFlagSet) {
	u := flag.String("u", "", "u 参数说明")
	s := flag.String("s", "", "s 参数说明")
	flag.Parse()
	fmt.Println("u:", *u)
	fmt.Println("s:", *s)
	fmt.Println("args:", flag.Args())
}

func cmdRun(flag *xflag.XFlagSet) {
	flag.Bool("a", false, "a 参数说明")
	flag.Parse()
}

func cmdBuild(flag *xflag.XFlagSet) {
	flag.Bool("b", false, "b 参数说明")
	flag.Parse()
}
