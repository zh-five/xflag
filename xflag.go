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

func Flag() *XFlagSet {
	return newXFlagSet(os.Args[0], 1)
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
			return
		} else {
			fmt.Fprintln(x.Output(), "Subcommand error:", s_cmd)
			os.Exit(2)
		}
	}

	x.sortParse(os.Args[x.level:]) //解析参数
}

// 调整顺序后解析
func (x *XFlagSet) sortParse(args []string) {
	all_num := len(args)
	param := make([]string, 0, all_num)
	new_args := args[0:0]

	start := 0
	for start < all_num {
		x.FlagSet.Parse(args[start:])
		tmp := x.FlagSet.Args()
		idx := len(tmp)
		offset := len(args[start:]) - idx
		new_args = append(new_args, args[start:start+offset]...)
		start += offset

		for i, v := range tmp {
			if v[0] == '-' {
				idx = i
				break
			}
		}
		//fmt.Println("arg:", args[start:], "tmp:", tmp, "idx:", idx, "new_args:", new_args, "param:", param)
		param = append(param, tmp[0:idx]...)
		start += idx
		//fmt.Println("arg:", args[start:], "tmp:", tmp, "idx:", idx, "new_args:", new_args, "param:", param, "\n ")
	}
	new_args = append(new_args, param...)
	//fmt.Println("end args:", new_args, args)
	x.FlagSet.Parse(new_args)
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
