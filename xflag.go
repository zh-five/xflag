package xflag

import (
	"flag"
	"fmt"
	"os"
)

type XFlagSet struct {
	*flag.FlagSet
	level uint                //当前子命令级别, 顶层为1
	cmds  map[string]*cmdInfo //下一级子命令信息列表
}

type cmdInfo struct {
	desc string
	cb   CmdFunc
}

type CmdFunc func(flag *XFlagSet)

//根节点
var root *XFlagSet = newXFlagSet(os.Args[0], 1)

func newXFlagSet(name string, level uint) *XFlagSet {
	return &XFlagSet{
		level:   level,
		FlagSet: flag.NewFlagSet(name, flag.ExitOnError),
		cmds:    make(map[string]*cmdInfo),
	}
}

// 子命令绑定
func SubCmd(name, desc string, cb CmdFunc) {
	root.SubCmd(name, desc, cb)
}

// 执行解析
func Parse() {
	root.Parse()
}

func (x *XFlagSet) SubCmd(name, desc string, cb CmdFunc) {
	x.cmds[name] = &cmdInfo{
		desc: desc,
		cb:   cb,
	}
}

func (x *XFlagSet) Parse() {
	h1 := x.Bool("h", false, "查看帮助")
	h2 := x.Bool("help", false, "查看帮助")

	// 有配置子命令
	if len(x.cmds) > 0 && len(os.Args) > int(x.level) { // 参数里有子命令
		s_cmd := os.Args[x.level]
		if info, ok := x.cmds[s_cmd]; ok { // 子命令合法
			tmp := newXFlagSet(x.Name()+" "+s_cmd, x.level+1)
			info.cb(tmp)
			return
		}
	}

	x.FlagSet.Parse(os.Args[x.level:]) //解析参数
	if *h1 || *h2 {
		x.PrintDefaults()
		os.Exit(0)
	}

}

func (x *XFlagSet) PrintDefaults() {
	if len(x.cmds) > 0 {
		fmt.Fprintln(x.Output(), "Subcommand of", x.Name())
		for k, v := range x.cmds {
			fmt.Fprintln(x.Output(), "    ", k, "\t", v.desc)
		}
	}

	x.Usage()
}
