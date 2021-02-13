package setting

import (
	"log"

	"gopkg.in/ini.v1"
)

type App struct {
	PrefixUrl  string
	TimeFormat string
}

var AppSetting = &App{}

type Server struct {
	RunMode  string
	HttpPort int
}

var ServerSetting = &Server{}

type Db struct {
	ConnUrl string
}

var DbSetting = &Db{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("db", DbSetting)

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
