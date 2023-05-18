package controller

import (
	"context"
	common_model "github.com/cellargalaxy/go_common/model"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/go_web_scaffold/model"
	"github.com/cellargalaxy/go_web_scaffold/service"
	"github.com/cellargalaxy/go_web_scaffold/static"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func Controller() error {
	engine := gin.Default()
	engine.Use(util.GinLog)

	engine.GET(common_model.PingPath, util.Ping)

	engine.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, common_model.StaticPath) {
			c.Header("Cache-Control", "max-age=86400")
		}
	})
	engine.StaticFS(common_model.StaticPath, http.FS(static.StaticFile))

	type LoginGetReq struct {
		Username string `json:"username" form:"username" query:"username"`
		Password string `json:"password" form:"password" query:"password"`
	}
	engine.POST(model.LoginGetPath, util.NewGinPost("LoginGet", func(ctx context.Context, req LoginGetReq) (any, error) {
		object, err := service.Login(ctx, req.Username, req.Password)
		if err != nil {
			return nil, err
		}
		return common_model.HttpData{Object: object}, nil
	}))

	type LoginPostReq struct {
		Username string `json:"username" form:"username" query:"username"`
		Password string `json:"password" form:"password" query:"password"`
	}
	engine.POST(model.LoginPostPath, util.NewGinPost("LoginPost", func(ctx context.Context, req LoginPostReq) (any, error) {
		object, err := service.Login(ctx, req.Username, req.Password)
		if err != nil {
			return nil, err
		}
		return common_model.HttpData{Object: object}, nil
	}))

	err := engine.Run(model.ListenAddress)
	if err != nil {
		errors.Errorf("web服务启动，异常: %+v", err)
	}
	return nil
}
