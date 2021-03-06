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

		a, err := b.s.Access(context.Background(), req.GetId())
		if err != nil {
			log.Println(err)
			continue
		}

		m := make(map[string]*access.Methods)
		for k, v := range a {
			m[k] = &access.Methods{Name: v}
		}

		resp.Access = m

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
