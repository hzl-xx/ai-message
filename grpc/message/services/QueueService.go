package services

import (
	"messageserver/grpc/message/protos"
	"github.com/streadway/amqp"
	"messageserver/grpc/message/configure"
	"messageserver/utils/log"
	"encoding/json"
)

type QueueService struct {
	dialHost string
	queueName string
	rabbitMqConn *amqp.Connection
}

func NewMessageService() *QueueService {
	queue := &QueueService{
		dialHost:configure.RABBMITMQ_HOST,
		queueName:configure.RABBMITMQ_NAME,
	}
	if err := queue.CreateDial(); err != nil{
		log.Info("CreateDial %+v", err)
	}

	return queue
}

func (q *QueueService) CreateDial() error  {
	conn, err := amqp.Dial(q.dialHost)

	if err != nil {
		return err
	}
	q.rabbitMqConn = conn
	return nil
}

func (q *QueueService) CloseRabbitMqConn()  {
	if err := q.rabbitMqConn.Close();err != nil {
		log.Info("CloseRabbitMqConn err %+v", err)
	}
}

func (q *QueueService) QueueDeclare(ch *amqp.Channel) (amqp.Queue,error)  {
	return ch.QueueDeclare(
		q.queueName, // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
}

func (q *QueueService) ConsumeMessage() (<-chan amqp.Delivery,error) {

	ch, err := q.rabbitMqConn.Channel()

	//defer func() {
	//	if err:= ch.Close();err != nil {
	//		log.Info("CloseChConn err %+v", err)
	//	}
	//}()

	que, err := q.QueueDeclare(ch)
	messagelist := make(<-chan amqp.Delivery)
	if err != nil {
		return messagelist,err
	}

	messagelist, err = ch.Consume(
		que.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	return messagelist,err
}



func (q *QueueService) PushMessage(req *protos.Message) error {


	body,err := json.Marshal(req)
	if err != nil {
		log.Info("parse json request err :", err)
		return err
	}
	ch, err := q.rabbitMqConn.Channel()
	defer func() {
		if err:= ch.Close();err != nil {
			log.Info("CloseChConn err %+v", err)
		}
	}()

	que, err := q.QueueDeclare(ch)

	if err != nil {
		return err
	}

	err = ch.Publish(
		"",     // exchange
		que.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(string(body)),
		})

	return err
}

