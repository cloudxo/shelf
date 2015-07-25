package actions

import (
	"github.com/codegangsta/cli"
	"strings"
	"github.com/matthistuff/shelf/helpers"
	"fmt"
	"github.com/matthistuff/shelf/data"
)

func AddTag(c *cli.Context) {
	if len(c.Args()) < 2 {
		helpers.ErrExit(true, "Not enough arguments provided")
	}

	objectId := c.Args().First()
	value := strings.Join(c.Args().Tail(), " ")

	object, err := data.GetObject(objectId)
	helpers.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	if (object.HasTag(value)) {
		return
	}

	object.AddTag(value)
}

func RemoveTag(c *cli.Context) {
	if len(c.Args()) < 2 {
		helpers.ErrExit(true, "Not enough arguments provided")
	}

	objectId := c.Args().First()
	value := strings.Join(c.Args().Tail(), " ")

	object, err := data.GetObject(objectId)
	helpers.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	object.RemoveTag(value)
}