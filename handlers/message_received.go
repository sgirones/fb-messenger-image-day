package handlers

import (
	"time"
	"fmt"
	"github.com/paked/messenger"
)

func MessageReceived(msg messenger.Message, res *messenger.Response) {
	fmt.Printf("%v (Sent, %v)\n", msg.Text, msg.Time.Format(time.UnixDate))

	res.Text(fmt.Sprintf("Hello, %v!", msg.Sender.ID))
}