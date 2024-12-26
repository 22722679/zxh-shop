package email

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/email/infra/mq"
	"github.com/cloudwego/biz-demo/gomall/app/email/infra/notify"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	tracer := otel.Tracer("shop-nats-consumer")
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}
		ctx := context.Background()
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(msg.Header))
		_, span := tracer.Start(ctx, "shop-email-consumer")
		defer span.End()

		noopemail := notify.NewNoopEmail()
		_ = noopemail.Send(&req)
	})
	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		mq.Nc.Close()
	})
}
