.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto


gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/order_page.proto --service frontend --module github.com/cloudwego/biz-demo/gomall/app/frontend -I ../../idl


gen_rpc:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/user.proto

gen_app_user:
	@cd app/user && cwgo server --type RPC --service user --module github.com/cloudwego/biz-demo/gomall/app/user --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto


gen_app_product:
	@cd rpc_gen && cwgo client --type RPC --service product --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module github.com/cloudwego/biz-demo/gomall/app/product --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto

gen_category:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/product_page.proto --service frontend --module github.com/cloudwego/biz-demo/gomall/app/frontend -I ../../idl


gen_app_cart:
	@cd rpc_gen && cwgo client --type RPC --service cart --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module github.com/cloudwego/biz-demo/gomall/app/cart --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.proto


gen_app_payment:
	@cd rpc_gen && cwgo client --type RPC --service payment --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module github.com/cloudwego/biz-demo/gomall/app/payment --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/payment.proto

gen_checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module github.com/cloudwego/biz-demo/gomall/app/checkout --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/checkout.proto

gen_order:
	@cd rpc_gen && cwgo client --type RPC --service order --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module github.com/cloudwego/biz-demo/gomall/app/order --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order.proto


gen_email:
	@cd rpc_gen && cwgo client --type RPC --service email --module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC --service email --module github.com/cloudwego/biz-demo/gomall/app/email --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/email.proto




run: ## run {svc} server. example: make run svc=product
	@scripts/run.sh ${svc}