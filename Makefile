shop:
	nohup cd app/frontend && go run . &
	nohup cd app/cart && go run . &
	nohup cd app/order && go run . &
	nohup cd app/product && go run . &
	nohup cd app/cart && go run . &
	nohup cd app/payment && go run . &
	nohup cd app/user && go run . &
	nohup cd app/email && go run . &
