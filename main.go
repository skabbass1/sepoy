package main

import(
	"fmt"
	"howett.net/plist"
)


type Label struct {
	Value string `plist:"Label"`
}

type RunAtLoad struct {
	Value bool `plist:"RunAtLoad"`
}

type Disabled struct {
	Value bool `plist:"Disabled"`
}

type KeepAlive  struct {
	Value bool `plist:"KeepAlive"`
}

type LaunchOnlyOnce  struct {
	Value bool `plist:"LaunchOnlyOnce"`
}

type ProgramArguments  struct {
	Value []string `plist:"ProgramArguments"`
}


type Plist struct {
	Label
	RunAtLoad
	Disabled
	KeepAlive
	LaunchOnlyOnce
	ProgramArguments
}

func main() {
	data := &Plist{
		Label{Value: "com.vault"},
		RunAtLoad{Value: false},
		Disabled{Value: false},
		KeepAlive{Value: false},
		LaunchOnlyOnce{Value: true},
		ProgramArguments{Value: []string{"hello.py", "world", "hi"}},
	}
	out, err := plist.MarshalIndent(data, plist.XMLFormat, "  ")
	if err != nil {
		    fmt.Println(err)
			}
			fmt.Println(string(out))

}
