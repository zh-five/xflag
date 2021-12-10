package xflag

import (
	"flag"
	"fmt"
	"os"
)

type XFlagSet struct {
	*flag.FlagSet
	level   uint                //当前子命令级别, 顶层为1 , Parse(os.Arg[1:])
	cmds    map[string]*cmdInfo //下一级子命令信息列表
	CmdName string              //被执行的子命令
}

type cmdInfo struct {
	desc string
	cb   CmdFunc
}

type CmdFunc func(flag *XFlagSet)

//根节点
var root *XFlagSet = newXFlagSet(os.Args[0], 1)

func Flag() *XFlagSet {
	if root.Parsed() {
		return nil
	}
	return root
}

func newXFlagSet(name string, level uint) *XFlagSet {
	x := &XFlagSet{
		level:   level,
		FlagSet: flag.NewFlagSet(name, flag.ExitOnError),
		cmds:    make(map[string]*cmdInfo),
	}

	usage := x.FlagSet.Usage
	x.FlagSet.Usage = func() {
		x.cmdList()
		usage()
	}

	return x
}

// 绑定子命令处理函数
func (x *XFlagSet) BindCmd(name string, cb CmdFunc, desc string) {
	x.cmds[name] = &cmdInfo{
		desc: desc,
		cb:   cb,
	}
}

// 解析
func (x *XFlagSet) Parse() {
	// 有配置子命令
	if len(x.cmds) > 0 && len(os.Args) > int(x.level) { // 参数里有子命令
		s_cmd := os.Args[x.level]
		if info, ok := x.cmds[s_cmd]; ok { // 子命令合法
			x.CmdName = s_cmd
			if info.cb != nil {
				tmp := newXFlagSet(x.Name()+" "+s_cmd, x.level+1)
				info.cb(tmp)
			}
		}
	}

	x.FlagSet.Parse(os.Args[x.level:]) //解析参数
}

func (x *XFlagSet) Desc(top, bottom string) {
	usage := x.FlagSet.Usage
	x.FlagSet.Usage = func() {
		if top != "" {
			fmt.Println(top)
		}
		usage()
		if bottom != "" {
			fmt.Println(bottom)
		}
	}
}

func (x *XFlagSet) cmdList() {
	if len(x.cmds) > 0 {
		fmt.Fprintln(x.Output(), "Subcommand of", x.Name()+":")
		for k, v := range x.cmds {
			fmt.Fprintln(x.Output(), "    ", k, "\t", v.desc)
		}
	}
}
