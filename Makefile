.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto


gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend --module github.com/cloudwego/biz-demo/gomall/app/frontend -I ../../idl/frontend/auto_page.proto


gen_rpc:
	@cd rep_gen && cwgo client --type RPC --service user --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/user.proto

gen_app_user:
	@cd app/user && cwgo server --type RPC --service user --module github.com/cloudwego/biz-demo/gomall/app/user --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto
