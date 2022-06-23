package broker

import (
	"app/access/internal/service"
	access "app/access/pkg/access/gen"
	"app/internal/broker"
	"context"
	"log"

	"google.golang.org/protobuf/proto"
)

type Broker struct {
	s service.Service
	b *broker.RabbitMQ
}

func New(b *broker.RabbitMQ, s service.Service) *Broker {
	return &Broker{
		s: s,
		b: b,
	}
}

func (b *Broker) Run() {
	go b.UserAccess()
}

func (b *Broker) UserAccess() {
	msgs, err := b.b.Consume("UserAccess", false)
	if err != nil {
		log.Fatal(err)
	}

	for d := range msgs {
		req := &access.UserAccessRequest{}
		resp := &access.UserAccessResponse{}

		err := proto.Unmarshal(d.Body, req)
		if err != nil {
			log.Println(err)
			continue
		}

		a, err := b.s.UserAccesses(context.Background(), req.GetId())
		if err != nil {
			log.Println(err)
			continue
		}

		for s, m := range a {
			resp.Access = append(resp.Access, &access.Service{
				Name:    s,
				Methods: m,
			})
		}

		err = b.b.ReplyTo(d, resp)
		if err != nil {
			log.Println(err)
		}

		d.Ack(false)

	}
}

func (b *Broker) Stop() {
	b.b.Close()
}
