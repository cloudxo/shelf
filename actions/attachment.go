package actions

import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/data"
	"os"
	"io"
	"path"
	"fmt"
	"github.com/matthistuff/shelf/helpers"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)


func AddAttachment(c *cli.Context) {
	objectId := helpers.ValidId(c.Args().First())
	filepath := c.Args().Get(1)

	object, err := data.GetObject(objectId)
	helpers.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	dbFile, err := data.Files().Create("")
	helpers.ErrPanic(err)

	file, err := os.Open(filepath)
	helpers.ErrPanic(err)
	defer file.Close()

	_, err = io.Copy(dbFile, file)
	helpers.ErrPanic(err)

	err = dbFile.Close()
	helpers.ErrPanic(err)

	attachment := data.CreateAttachment(dbFile, path.Base(filepath))
	object.Attachments = append(object.Attachments, *attachment)
	object.Update()

	fmt.Println(attachment.Id.Hex())
}

func GetAttachment(c *cli.Context) {
	objectId := helpers.ValidId(c.Args().First())

	file, err := data.Files().OpenId(bson.ObjectIdHex(objectId))
	helpers.ErrPanic(err)

	_, err = io.Copy(os.Stdout, file)
	helpers.ErrPanic(err)

	err = file.Close()
	helpers.ErrPanic(err)
}

func ListAttachments(c *cli.Context) {
	helpers.Color(c)

	objectId := helpers.ValidId(c.Args().First())

	object, err := data.GetObject(objectId)
	helpers.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	if len(object.Attachments) > 0 {
		data.ClearCache()
		defer data.FlushCache()

		for index, attachment := range object.Attachments {
			fmt.Printf("(%s) %s \"%s\"\n", helpers.ShortId(index+1), helpers.ObjectId(attachment.Id.Hex()), attachment.Filename)
			data.SetCache(strconv.Itoa(index+1), attachment.Id.Hex())
		}
	}
}