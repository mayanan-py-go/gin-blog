package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PageSize int
	JwtSecret string
	TokenTimeout time.Duration
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}
var AppSetting = &App{}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
}
var DatabaseSetting = &Database{}

type Redis struct {
	Host string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}
var RedisSetting = &Redis{}

func Setup() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("fail to parse 'conf/app.ini': %v", err)
	}

	if err = cfg.Section("app").MapTo(AppSetting); err != nil {
		log.Fatalf("cfg.Mapto AppSetting err:%v", err)
	}
	AppSetting.ImageMaxSize *= 1024 * 1024

	if err = cfg.Section("server").MapTo(ServerSetting); err != nil {
		log.Fatalf("cfg.Mapto ServerSetting err:%v", err)
	}
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second

	if err = cfg.Section("database").MapTo(DatabaseSetting); err != nil {
		log.Fatalf("cfg.Mapto DatabaseSetting err:%v", err)
	}

	if err = cfg.Section("redis").MapTo(RedisSetting); err != nil {
		log.Fatalf("cfg.Mapto RedisSetting err:%v", err)
	}
	RedisSetting.IdleTimeout *= time.Second
}
