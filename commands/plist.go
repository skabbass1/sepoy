package commands

import (
	"fmt"
	"io/ioutil"

	goplist "github.com/DHowett/go-plist"
	"github.com/skabbass1/sepoy/plist"
	"github.com/urfave/cli"
)

var Plist = cli.Command{
	Name:  "plist",
	Usage: "display plist attributes for the scheduled task",
	Action: func(c *cli.Context) error {

		if len(c.Args()) == 0 {
			return cli.NewExitError("task name must be provided as first argument", 1)
		}
		myplist, err := TaskPlist(c.Args().Get(0))
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		fmt.Printf("%+v", myplist)
		return nil

	},
}

func TaskPlist(taskName string) (plist.Plist, error) {

	myplist := plist.Plist{}
	bytes, _ := ioutil.ReadFile(PlistLocation(taskName))
	_, err := goplist.Unmarshal(bytes, &myplist)
	return myplist, err

}
