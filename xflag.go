package xflag

import (
	"flag"
	"fmt"
	"os"
)

type XFlagSet struct {
	*flag.FlagSet
	args []string
	cmds map[string]*cmdInfo
}

type cmdInfo struct {
	name string
	desc string
	cb   CmdFunc
}

type CmdFunc func(x *XFlagSet, args []string) error

var root *XFlagSet

func init() {
	root = &XFlagSet{
		args:    os.Args,
		FlagSet: flag.NewFlagSet("root", flag.ExitOnError),
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
		name: name,
		desc: desc,
		cb:   cb,
	}
}

func (x *XFlagSet) Parse() {
	//子命令参数解析
	if len(x.cmds) == 0 {
		x.FlagSet.Parse(x.args)
		return
	}

	// 解析子命令

}

func (x *XFlagSet) PrintDefaults() {
	fmt.Println("PrintDefaults()...")
}
