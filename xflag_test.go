package xflag

import (
	"os"
	"testing"
)

func TestFlag(t *testing.T) {
	f := Flag()
	f.Desc("top", "bottom")

	if f.Parsed() {
		t.Fail()
	}
}

func TestDesc(t *testing.T) {
	f := Flag()
	f.Desc("top", "bottom")
}

func TestPaser(t *testing.T) {
	os.Args = []string{"test", "-f", "file1", "file2"}

	fg := Flag()
	f := fg.Bool("f", false, "ff")
	fg.Parse()

	if !*f {
		t.Fail()
	}
	if fg.Arg(0) != "file1" {
		t.Fail()
	}
	if fg.Arg(1) != "file2" {
		t.Fail()
	}

	//os.Args = os_args
}

func TestPaser1(t *testing.T) {
	os.Args = []string{"test", "file1", "-f", "file2"}

	fg := Flag()
	f := fg.Bool("f", false, "ff")
	fg.Parse()

	if !*f {
		t.Fail()
	}
	if fg.Arg(0) != "file1" {
		t.Fail()
	}
	if fg.Arg(1) != "file2" {
		t.Fail()
	}
}

func TestPaser2(t *testing.T) {
	os.Args = []string{"test", "file1", "file2", "-f"}

	fg := Flag()
	f := fg.Bool("f", false, "ff")
	fg.Parse()

	if !*f {
		t.Fail()
	}
	if fg.Arg(0) != "file1" {
		t.Fail()
	}
	if fg.Arg(1) != "file2" {
		t.Fail()
	}
}

func TestCmd(t *testing.T) {
	os.Args = []string{"test", "cmd", "file2", "-f"}
	cmd_ok := false
	fg := Flag()
	fg.BindCmd("cmd", func(flag *XFlagSet) {
		cmd_ok = true
		f := flag.Bool("f", false, "ff")
		flag.Parse()

		if !*f {
			t.Fail()
		}
		if flag.Arg(0) != "file2" {
			t.Fail()
		}
	}, "cmd...")
	fg.Parse()

	if !cmd_ok {
		t.Fail()
	}
	if fg.CmdName != "cmd" {
		t.Fail()
	}
}

func TestCmd1(t *testing.T) {
	os.Args = []string{"test", "cmd", "cmd", "-f", ""}
	cmd_ok := false
	fg := Flag()
	fg.BindCmd("cmd", func(flag *XFlagSet) {
		cmd_ok = true
		f := flag.Bool("f", false, "ff")
		flag.BindCmd("cmd1", func(flag *XFlagSet) {

		}, "cmd1...")
		flag.Parse()

		if *f {
			t.Fail()
		}
		if flag.Arg(0) != "file2" {
			t.Fail()
		}
	}, "cmd...")
	fg.Parse()

	if !cmd_ok {
		t.Fail()
	}
	if fg.CmdName != "cmd" {
		t.Fail()
	}
}
