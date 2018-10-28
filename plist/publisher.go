package plist

import (
	"io/ioutil"

	goplist "github.com/DHowett/go-plist"
)

func PublishPlist(data Plist, location string) error {

	bytes, err := goplist.MarshalIndent(data, goplist.XMLFormat, "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(location, bytes, 0644)
	return err
}
