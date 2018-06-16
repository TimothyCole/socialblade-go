# Social Blade API in Go [<img align="right" src="https://socialblade.com/images/media/red/2.png">](https://github.com/TimothyCole/socialblade-go)

[![Build Status](https://travis-ci.com/TimothyCole/socialblade-go.svg?branch=master)](https://travis-ci.com/TimothyCole/socialblade-go)
[![GoDoc](https://godoc.org/github.com/TimothyCole/socialblade-go?status.svg)](https://godoc.org/github.com/TimothyCole/socialblade-go)
[![Personal Discord](https://img.shields.io/discord/313591755180081153.svg?label=Personal%20Discord&colorB=308bcd&maxAge=3600)](https://discordapp.com/invite/YFtfGwq)
[![Social Blade Discord](https://img.shields.io/discord/125022847562285056.svg?label=Social%20Blade%20Discord%20(Not%20for%20Support)&colorB=c84329&maxAge=3600)](https://socialblade.com/discord)

---

### Information and Disclaimer
Just because I work for Social Blade does not mean this is an official library. This is an unofficial library created for my personal convenience on my personal time.

This library supports most of the `api.socialblade.com` endpoints that are being exposed by either the official Social Blade Browser Extensions, Apps, or Website. Any endpoints that are not exposed publically will not be added until done so.

Social Blade's API is a private API used for Social Blade official tools and for companies with express written permission. Use of the Social Blade API without express permission is prohibited by the active [Terms of Service](https://socialblade.com/info/terms) and may result in being blocked from the Social Blade website and all of it's data.

---

## Install
```
go get github.com/TimothyCole/socialblade-go
```

## Usage
Import the library into your project.
```go
import "github.com/TimothyCole/socialblade-go"
```

Construct a new client which will be used to access the API by either the third-party `Auth` function or first-party `AuthAsUser` function.
```go
// Third-Party Auth
sb, err := socialblade.Auth("Third-Party API Key")

// First-Party Auth
sb, err := socialblade.AuthAsUser("User Email", "User Access Token")
```

## Example
Getting YouTube Stats
```go
sb, _ := socialblade.Auth("Third-Party API Key")
stats, _ := sb.StatsYouTube("EatTim")
fmt.Println(stats.Data.DisplayName, "has", stats.Data.Subs, "subscribers")
```

## Troubleshooting
If you have official access to the Social Blade API via express permission from the company then feel free to contact me via [Twitter](https://twitter.com/messages/compose?recipient_id=1690693537) or [email me](mailto:tim@timcole.me?cc=tim@socialblade.com&subject=Social%20Blade%20Golang%20Library%20Inquiry) for any help.