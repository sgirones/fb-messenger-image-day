package handlers

import (
	"time"
	"fmt"
	"github.com/paked/messenger"
)

// Setup a handler to be triggered when a message is read
func MessageDelivered(d messenger.Delivery, res *messenger.Response) {
	fmt.Println(d.Watermark().Format(time.UnixDate))
}