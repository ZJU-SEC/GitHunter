package config

import "gopkg.in/ini.v1"

var Config *ini.File

// APP
var WORKER int
var QUEUE_SIZE int
var LANGUAGE string
var MIN_STAR int
var REPOS_PATH string
var DEBUG bool
var CLONE_LIMIT int

// Web
var GITHUB_TOKEN []string
var TRYOUT int


func Init() {
	var err error

	// load config.ini file
	Config, err = ini.ShadowLoad("config.ini")
	if err != nil {
		panic(err)
	}

	// load APP section
	APPSection, err := Config.GetSection("APP")
	if err != nil {
		panic(err)
	}
	WORKER = APPSection.Key("WORKER").MustInt(16)
	QUEUE_SIZE = APPSection.Key("QUEUE_SIZE").MustInt(128)
	LANGUAGE = APPSection.Key("LANGUAGE").String()
	MIN_STAR = APPSection.Key("MIN_STAR").MustInt(20)
    REPOS_PATH = APPSection.Key("REPOS_PATH").String()
	DEBUG = APPSection.Key("DEBUG").MustBool(false)
	CLONE_LIMIT = APPSection.Key("CLONE_LIMIT").MustInt(100)

	// load WEB section
	WEBSection, err := Config.GetSection("WEB")
	if err != nil {
		panic(err)
	}
	GITHUB_TOKEN = WEBSection.Key("GITHUB_TOKEN").ValueWithShadows() // parse token list from config
	TRYOUT = WEBSection.Key("TRYOUT").MustInt(5)
}

func ParseKey(section *ini.Section, key string) string {
	return section.Key(key).String()
}
