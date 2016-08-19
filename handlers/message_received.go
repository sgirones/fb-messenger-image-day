package handlers

import (
	"time"
	"fmt"
	"github.com/paked/messenger"
	_ "image/png"
	"os"
	"log"
	"image"

)

var (
	imagepath = "./charmander.png"
)

func MessageReceived(msg messenger.Message, res *messenger.Response) {
	fmt.Printf("%v (Sent, %v)\n", msg.Text, msg.Time.Format(time.UnixDate))

	reader, err := os.Open(imagepath)
	if err != nil {
	    log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Replying with Charmander's picture!")
	res.Image(m)
}