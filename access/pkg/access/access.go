package access

import (
	access "app/access/pkg/access/gen"
	"app/internal/broker"

	"google.golang.org/protobuf/proto"
)

type Client struct {
	b *broker.RabbitMQ
}

func New(b *broker.RabbitMQ) *Client {
	return &Client{
		b: b,
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
