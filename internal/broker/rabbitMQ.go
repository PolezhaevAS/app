package broker

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Config struct {
	URL string `yaml:"url" mapstructure:"url"`
}

func NewConfig() *Config {
	return &Config{
		URL: "amqp://guest:guest@localhost:5672/",
	}
}

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	s    grpc.ServiceDesc
}

func New(service grpc.ServiceDesc, cfg *Config) (r *RabbitMQ, err error) {
	r = new(RabbitMQ)
	r.conn, err = amqp.Dial(cfg.URL)
	if err != nil {
		return nil, err
	}

	r.ch, err = r.conn.Channel()
	if err != nil {
		return nil, err
	}

	r.s = service

	for _, m := range service.Methods {
		_, err = r.QueueDeclare(m.MethodName)
		if err != nil {
			return nil, err
		}
	}

	return r, nil
}

func (r *RabbitMQ) queueName(name string) string {
	return fmt.Sprintf("%s.%s", r.s.ServiceName, name)
}

func (r *RabbitMQ) Close() {
	r.ch.Close()
	r.conn.Close()
}

func (r *RabbitMQ) QueueDeclare(name string) (amqp.Queue, error) {
	q, err := r.ch.QueueDeclare(
		r.queueName(name), // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	return q, err
}

func (r *RabbitMQ) Consume(name string, autoAck bool) (<-chan amqp.Delivery, error) {

	msgs, err := r.ch.Consume(
		r.queueName(name), // queue
		"",                // consumer
		autoAck,           // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (r *RabbitMQ) replyConsume(q amqp.Queue) (<-chan amqp.Delivery, error) {

	msgs, err := r.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (r *RabbitMQ) Publish(routeKey string, body []byte, corrId string, replyTo string) error {

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	}

	if corrId != "" {
		msg.CorrelationId = corrId
	}

	if replyTo != "" {
		msg.ReplyTo = replyTo
	}

	err := r.ch.Publish(
		"",                    // exchange
		r.queueName(routeKey), // routing key
		false,                 // mandatory
		false,                 // immediate
		msg)
	if err != nil {
		return err
	}
	return nil
}

func (r *RabbitMQ) ReplyTo(reply amqp.Delivery, msg protoreflect.ProtoMessage) error {

	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	err = r.ch.Publish(
		"",
		reply.ReplyTo,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: reply.CorrelationId,
			Body:          body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *RabbitMQ) RPC(routeKey string, body []byte, replyTo string) ([]byte, error) {

	q, err := r.QueueDeclare("")
	if err != nil {
		return nil, err
	}

	msgs, err := r.replyConsume(q)
	if err != nil {
		return nil, err
	}
	corrId := randomString(32)
	err = r.Publish(routeKey, body, corrId, q.Name)
	if err != nil {
		return nil, err
	}

	for d := range msgs {
		log.Println("Msg", d.Body)
		if corrId == d.CorrelationId {
			return d.Body, nil
		}
	}

	return nil, nil
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
