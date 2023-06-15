package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	LoadServer(cfg)
	LoadData(cfg)
}

func LoadServer(cfg *ini.File) {
	AppMode = cfg.Section("server").Key("AppMode").String()
	HttpPort = cfg.Section("server").Key("HttpPort").String()
}

func LoadData(cfg *ini.File) {
	Db = cfg.Section("database").Key("Db").String()
	DbHost = cfg.Section("database").Key("DbHost").String()
	DbPort = cfg.Section("database").Key("DbPort").String()
	DbUser = cfg.Section("database").Key("DbUser").String()
	DbPassWord = cfg.Section("database").Key("DbPassWord").String()
	DbName = cfg.Section("database").Key("DbName").String()
}
