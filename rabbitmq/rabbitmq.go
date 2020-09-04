package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn     *amqp.Connection
	Channel  *amqp.Channel
	Name     string //queue name
	Exchange string //exchange name
}

func New() *RabbitMQ {
	//connect rabbitmq
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/%s", "test", "test888", "10.12.35.8", "5672", "test-host")) //host
	if err != nil {
		panic(err)
	}
	//create a channel
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	//create a queue
	/*queue, err := ch.QueueDeclare(
		"",    //queue name
		true,  //durable
		true,  //delete when unused
		false, //exclusive
		false, //no-wait
		nil,   //arguments
	)
	if err != nil {
		panic(err)
	}*/
	mq := RabbitMQ{Channel: ch, Conn: conn}
	return &mq
}

func (mq *RabbitMQ) BindExchange(exchange string) {
	//queue bind exchange
	err := mq.Channel.QueueBind(
		mq.Name,  //queue name
		"",       //routing key
		exchange, //exchange
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	mq.Exchange = exchange
}

func (mq *RabbitMQ) Send(exchangeName string, queueName string, body interface{}) {
	msg, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	err = mq.Channel.Publish(
		exchangeName,
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		panic(err)
	}
}

func (mq *RabbitMQ) Receive(queueName string) <-chan amqp.Delivery {
	msg, err := mq.Channel.Consume(
		queueName, //queue name
		"",        //consumer name
		true,      //autoAck
		false,     //exclusive
		false,     //nolcoal
		false,     //nowait
		nil,       //args
	)
	if err != nil {
		panic(err)
	}
	return msg
}

func (mq *RabbitMQ) Close() {
	err := mq.Channel.Close()
	if err != nil {
		panic(err)
	}
	err = mq.Conn.Close()
	if err != nil {
		panic(err)
	}
}
