package access

import (
	"app/access/internal/service"
	access "app/access/pkg/access/gen"
	"app/internal/broker"

	"google.golang.org/protobuf/proto"
)

type Client struct {
	b *broker.RabbitMQ
	s service.Service
}

func New(b *broker.RabbitMQ, s service.Service) *Client {
	return &Client{
		b: b,
		s: s,
	}
}

func (c *Client) UserAccess(req *access.UserAccessRequest) (*access.UserAccessResponse, error) {
	body, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	respBody, err := c.b.RPC("UserAccess", body, "UserAccess")
	if err != nil {
		return nil, err
	}
	resp := &access.UserAccessResponse{}
	err = proto.Unmarshal(respBody, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
