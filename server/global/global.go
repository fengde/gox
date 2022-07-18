package global

import (
	"flag"

	"github.com/fengde/gocommon/confx"
	_ "github.com/joho/godotenv/autoload"
)

var (
	Conf Config
)

func Init() {
	flag.Parse()
	confx.MustLoad(*flag.String("f", "config.yaml", "the config file"), &Conf, confx.UseEnv())
}
