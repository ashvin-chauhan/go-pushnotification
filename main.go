package main

import (
	"final_task/push_notification"
)

func main() {
	pushnotification.Android("Hi, Android!", "devise token", "server key")
  pushnotification.IOS("Hi, IOS!", "devise token", "pem file path")
}
