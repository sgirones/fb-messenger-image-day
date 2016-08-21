package handlers

import (
	"time"
	"fmt"
	"github.com/sgirones/messenger"
)

var (
	imageUrl = "http://cdn.bulbagarden.net/upload/thumb/7/73/004Charmander.png/250px-004Charmander.png"
)

func MessageReceived(msg messenger.Message, res *messenger.Response) {
	fmt.Printf("%v (Sent, %v)\n", msg.Text, msg.Time.Format(time.UnixDate))

	res.ImageWithURL(imageUrl)
}