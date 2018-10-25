package main

import (
	"fmt"

	"github.com/skabbass1/sepoy/nodes"
	"howett.net/plist"
)

func main() {
	data := &nodes.Plist{
		nodes.Label{Value: "com.vault"},
		nodes.RunAtLoad{Value: false},
		nodes.Disabled{Value: false},
		nodes.KeepAlive{Value: false},
		nodes.LaunchOnlyOnce{Value: true},
		nodes.ProgramArguments{Value: []string{"hello.py", "world", "hi"}},
		nodes.StartCalendarInterval{},
	}
	out, err := plist.MarshalIndent(data, plist.XMLFormat, "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

}
