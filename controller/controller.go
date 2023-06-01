package controller

import (
	"context"
	common_model "github.com/cellargalaxy/go_common/model"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/go_web_scaffold/config"
	"github.com/cellargalaxy/go_web_scaffold/db"
	"github.com/cellargalaxy/go_web_scaffold/model"
	"github.com/cellargalaxy/go_web_scaffold/service"
	"github.com/cellargalaxy/go_web_scaffold/static"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Init() {
	engine := gin.Default()
	engine.Use(util.GinLog)
	engine.Use(claims)

	debug := engine.Group(common_model.DebugPath, validate)
	pprof.RouteRegister(debug, common_model.PprofPath)

	engine.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, common_model.StaticPath) {
			c.Header("Cache-Control", "max-age=86400")
		}
	})
	engine.StaticFS(common_model.StaticPath, http.FS(static.StaticFile))

	engine.GET(common_model.PingPath, util.Ping)
	engine.POST(common_model.PingPath, validate, util.Ping)

	engine.GET(model.LoginGetPath, util.NewGinGet("LoginGet", func(ctx context.Context, req struct{}) (any, error) {
		return common_model.HttpData{Object: config.Config}, nil
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

	type AddUserReq struct {
		Object []*model.User `json:"object"`
	}
	engine.POST(model.AddUserPath, validate, util.NewGinPost("AddUser", func(ctx context.Context, req AddUserReq) (any, error) {
		object, err := db.User.Add(ctx, req.Object...)
		if err != nil {
			return nil, err
		}
		return common_model.HttpData{Object: object}, nil
	}))

	engine.POST(model.RemoveUserPath, validate, util.NewGinPost("RemoveUser", func(ctx context.Context, req model.UserInquiry) (any, error) {
		err := db.User.Remove(ctx, req)
		if err != nil {
			return nil, err
		}
		return common_model.HttpData{}, nil
	}))

	type EditUserReq struct {
		Object *model.User `json:"object"`
	}
	engine.POST(model.EditUserPath, validate, util.NewGinPost("EditUser", func(ctx context.Context, req EditUserReq) (any, error) {
		object, count, err := db.User.Edit(ctx, req.Object)
		if err != nil {
			return nil, err
		}
		return common_model.HttpData{Object: object, Count: count}, nil
	}))

	engine.GET(model.ListUserPath, validate, util.NewGinGet("ListUser", func(ctx context.Context, req model.UserInquiry) (any, error) {
		object, count, err := db.User.List(ctx, req)
		if err != nil {
			return nil, err
		}
		return common_model.HttpData{Object: object, Count: count}, nil
	}))

	err := engine.Run(model.ListenAddress)
	if err != nil {
		panic(err)
	}
}

func claims(c *gin.Context) {
	util.ClaimsGin(c, config.Config.Password)
}
func validate(c *gin.Context) {
	util.ValidateGin(c, config.Config.Password)
}
