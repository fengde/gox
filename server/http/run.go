package http

import (
	"server/global"
	"server/http/router"
	"syscall"
	"time"

	"github.com/fengde/gocommon/logx"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func Run() error {
	r := gin.Default()

	router.Init(r)

	s := endless.NewServer(global.Conf.HttpAddress, r)
	s.ReadTimeout = 30 * time.Second
	s.WriteTimeout = 30 * time.Second
	s.MaxHeaderBytes = 1 << 24 // 16M大小

	s.BeforeBegin = func(add string) {
		logx.Infof("Actual pid is %d", syscall.Getpid())
	}
	if err := s.ListenAndServe(); err != nil {
		logx.Errorf("ListenAndServe err: %v", err)
		return err
	}

	return nil
}
