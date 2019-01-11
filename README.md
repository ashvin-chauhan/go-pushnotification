## Getting started
This package is about push notification for Andoid and IOS

```
Go version: 1.10.4
```

### Install
```
go get github.com/ashvin-chauhan/go-pushnotification
```

### Android
```
push_notification.Android("Hi, Android!", "Devise Token", "Server key")
```

### IOS Using pem file
```
push_notification.IOS("Hi, IOS!", "Devise Token", "APNs pem file path")
```

### IOS Using p8 file
```
push_notification.IOSUsingP8("Hi, IOS!", "Devise Token", "APNs p8 file path", "Key ID", "Team ID", "Topic", "Environment production/development")
```
