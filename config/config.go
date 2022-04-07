package config

import "gopkg.in/ini.v1"

var Config *ini.File

// APP
var WORKER int
var QUEUE_SIZE int
var LANGUAGE string
var MIN_STAR int
var DEBUG bool

// Web
var GITHUB_TOKEN []string
var TRYOUT int

// Repo
var CLONE_WORKER_NUMBER int
var CLONE_LOWER_BOUND int
var CLONE_STORAGE_PATH string
var CLONE_BATCH_SIZE int

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
	DEBUG = APPSection.Key("DEBUG").MustBool(false)

	// load WEB section
	WEBSection, err := Config.GetSection("WEB")
	if err != nil {
		panic(err)
	}
	GITHUB_TOKEN = WEBSection.Key("GITHUB_TOKEN").ValueWithShadows() // parse token list from config
	TRYOUT = WEBSection.Key("TRYOUT").MustInt(5)

	// load REPO section
	REPOSection, err := Config.GetSection("REPO")
	if err != nil {
		panic(err)
	}
	CLONE_WORKER_NUMBER = REPOSection.Key("CLONE_WORKER_NUMBER").MustInt(16)
	CLONE_LOWER_BOUND = REPOSection.Key("CLONE_LOWER_BOUND").MustInt(100)
	CLONE_STORAGE_PATH = REPOSection.Key("CLONE_STORAGE_PATH").String()
	CLONE_BATCH_SIZE = REPOSection.Key("CLONE_BATCH_SIZE").MustInt(256)
}

func ParseKey(section *ini.Section, key string) string {
	return section.Key(key).String()
}
