package handlers

import (
	"time"
	"fmt"
	"github.com/paked/messenger"
)

func MessageDelivered(d messenger.Delivery, res *messenger.Response) {
	fmt.Println("Message read by the user!")
	fmt.Println(d.Watermark().Format(time.UnixDate))
}