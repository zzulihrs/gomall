// Code generated by hertz generator.

package order

import (
	"github.com/cloudwego/biz-demo/gomall/app/frontend/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Auth()}
}

func _orderlistMw() []app.HandlerFunc {
	// your code...
	return nil
}
