package push_notification

import (
	"fmt"
	"log"
	"path/filepath"

	fcm "github.com/NaySoftware/go-fcm"
	"github.com/anachronistic/apns"
	"github.com/manucorporat/try"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	"github.com/sideshow/apns2/token"
)

func Android(message string, device_token string, server_key string) {
	try.This(func() {
		tokens := []string{
			device_token,
		}
		data := map[string]string{"message": message}

		client := fcm.NewFcmClient(server_key)
		client.NewFcmRegIdsMsg(tokens, data)

		status, err := client.Send()
		if err == nil {
			fmt.Println("No error")
			status.PrintResults()
		} else {
			fmt.Println("[PUSH NOTIFICATION ERROR]", err)
		}
	}).Finally(func() {
		fmt.Println("Finally")
	}).Catch(func(e try.E) {
		fmt.Println("Catch", e)
	})
}

func IOS(message string, device_token string, apns_file_path string) {
	try.This(func() {
		apns_path, path_err := filepath.Abs(apns_file_path)
		if path_err != nil {
			fmt.Println(path_err)
		}
		payload := apns.NewPayload()
		payload.Alert = message
		payload.Sound = "Default"
		payload.Badge = 1

		pn := apns.NewPushNotification()
		pn.DeviceToken = device_token
		pn.AddPayload(payload)

		var client *apns.Client
		client = apns.NewClient("gateway.push.apple.com:2195", apns_path, apns_path)
		resp := client.Send(pn)

		alert, _ := pn.PayloadString()
		fmt.Println("Alert:", alert)
		fmt.Println("Success:", resp.Success)
		fmt.Println("Error:", resp.Error)
	}).Finally(func() {
		fmt.Println("Finally block")
	}).Catch(func(e try.E) {
		fmt.Println("Catch", e)
	})
}

func IOSUsingP8(message string, devise_token string, apns_file_path string, key_id string, team_id string, topic string) {
	try.This(func() {
		authKey, err := token.AuthKeyFromFile(apns_file_path)
		if err != nil {
			fmt.Println("token error:", err)
		}

		token := &token.Token{
			AuthKey: authKey,
			KeyID:   key_id,
			TeamID:  team_id,
		}

		payload := payload.NewPayload().Alert(message).Badge(1).Custom("key", "val")

		notification := &apns2.Notification{
			DeviceToken: devise_token,
			Topic:       topic,
			Payload:     payload,
		}

		client := apns2.NewTokenClient(token)
		res, err := client.Push(notification)

		if err != nil {
			log.Println("[PUSH NOTIFICATION ERROR]", err)
			return
		}

		if res.Sent() {
			log.Println("Sent:", res.ApnsID)
		} else {
			fmt.Printf("Not Sent: %v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
		}

	}).Finally(func() {
		fmt.Println("Finally block")
	}).Catch(func(e try.E) {
		fmt.Println("Catch", e)
	})
}
