package conf

import (
	"github.com/go-ini/ini"
	"log"
	"skrshop-api/utils"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort string

	JwtSecret     string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	DBType        string
	DBSource      string
	RedisHost     string
	RedisPassword string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	loadBase()
	loadServer()
	loadApp()
	loadDB()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	port := sec.Key("HTTP_PORT").MustInt(8000)
	HTTPPort = ":" + utils.IntToString(port)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}

func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}
func loadDB() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	DBType = sec.Key("TYPE").MustString("mysql")
	DBSource = sec.Key("SOURCE").MustString("")
}
func loadRedis() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatalf("Fail to get section 'redis': %v", err)
	}
	RedisHost = sec.Key("HOST").MustString("127.0.0.1:6379")
	RedisPassword = sec.Key("PASSWORD").MustString("")
}
