package sdk

import (
	"fmt"
	"net/http"
	"time"

	"mhy-sdk/conf"
	"mhy-sdk/hkrpg"
	"mhy-sdk/logger"
	"mhy-sdk/sdk/mdk"

	"github.com/gin-gonic/gin"
)

type Sdk struct {
	router  *gin.Engine
	httpNet *conf.HttpNet
	hkRpg   *hkrpg.HkRpg
}

func NewSdk(cfg *conf.Config) *Sdk {
	gin.SetMode(gin.ReleaseMode)
	s := &Sdk{
		httpNet: cfg.HttpNet,
		hkRpg:   hkrpg.NewHkRpg(),
	}
	if logger.GetLogLevel() == logger.DEBUG {
		s.router = gin.Default()
	} else {
		s.router = gin.New()
	}
	s.router.Use(gin.Recovery())
	s.newRouter()
	mdk.IsAutoCreate = cfg.AutoCreate
	// 初始化hkrpg
	s.hkRpg.NewRouter(s.router)

	go s.UpUpstreamServer() // 定时器

	return s
}

func (s *Sdk) Run() error {
	addr := fmt.Sprintf("%s:%s", s.httpNet.InnerAddr, s.httpNet.InnerPort)
	logger.Info("http监听地址:%s", addr)
	server := &http.Server{Addr: addr, Handler: s.router}
	return server.ListenAndServe()
}

func (s *Sdk) UpUpstreamServer() {
	ticker := time.NewTicker(time.Minute * 5)
	for {
		select {
		case <-ticker.C:
			s.hkRpg.NewConf()
		}
	}
}
