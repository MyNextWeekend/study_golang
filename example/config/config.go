package config

import (
	"gopkg.in/ini.v1"
)

var (
	HttpPort        string
	WebSocketScheme string
	WebSocketHost   string
	WebSocketPath   string
)

func InitConfig() {
	file, err := ini.Load("./example/config/config.ini")
	if err != nil {
		panic("init load failed" + err.Error())
	}
	HttpPort = file.Section("server").Key("HttpPort").String()
	WebSocketScheme = file.Section("websocket").Key("WebSocketScheme").String()
	WebSocketHost = file.Section("websocket").Key("WebSocketHost").String()
	WebSocketPath = file.Section("websocket").Key("WebSocketPath").String()
}
