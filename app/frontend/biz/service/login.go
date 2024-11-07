package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	auth "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	err = session.Save()

	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}

	return
}
