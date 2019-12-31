package mq

import "fmt"

type Msg struct {
	Content string
}

func (m *Msg) MessageContent() string {
	return m.Content
}

func (m *Msg) Consumer(data []byte) error {
	fmt.Println("AMQP 消费消息")
	fmt.Println("Amqp receive message: ", string(data))
	return nil
}

func InitEngineMQ() {
	//msg := fmt.Sprintf("This is test message")

	te := &Msg{}

	queueExchange := &QueueExchange{
		QueueName:    "kunkka.queue.match",
		RouteKey:     "match",
		ExchangeName: "kunkka.exchange.match",
		ExchangeType: "direct",
	}
	te.Content = "Hello World"
	mq := NewAmqp(queueExchange)
	mq.RegisterProducer(te)
	mq.RegisterReceiver(te)
	mq.Start()
}
