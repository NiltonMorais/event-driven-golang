package queue

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"time"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	uri       string
	conn      *amqp.Connection
	listeners map[string][]Listener
}

type QueueMessage struct {
	Body []byte
}

func NewRabbitMQAdapter(uri string) *RabbitMQAdapter {
	return &RabbitMQAdapter{
		uri:       uri,
		listeners: make(map[string][]Listener),
	}
}

func (r *RabbitMQAdapter) Connect(ctx context.Context) error {
	conn, err := amqp.Dial(r.uri)
	if err != nil {
		return err
	}
	r.conn = conn
	return nil
}

func (r *RabbitMQAdapter) Disconnect(ctx context.Context) error {
	return r.conn.Close()
}

func (r *RabbitMQAdapter) Publish(ctx context.Context, eventPayload interface{}) error {
	eventName := reflect.TypeOf(eventPayload).Name()

	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		eventName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	eventJson, err := json.Marshal(eventPayload)
	if err != nil {
		return errors.New("error converting struct to json")
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(eventJson),
		})
	if err != nil {
		return err
	}
	log.Printf(" [x] Sent %s\n", eventJson)
	return nil
}

func (r *RabbitMQAdapter) StartConsuming(ctx context.Context, queueName string) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.ConsumeWithContext(
		ctx,
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			hasError := false
			for _, listener := range r.listeners[queueName] {
				eventDomain := event.DomainEvent{
					Date:    time.Now(),
					Type:    listener.eventType,
					Payload: d.Body,
				}
				err := listener.callback(ctx, eventDomain)
				if err != nil {
					log.Printf("Error processing message: %s", err)
					hasError = true
					break
				}
			}
			if !hasError {
				d.Ack(false)
			}
		}
	}()

	var forever chan struct{}
	log.Printf(" [*] Waiting for messages on queue %s. To exit press CTRL+C", queueName)
	<-forever
	return nil
}

func (r *RabbitMQAdapter) ListenerRegister(eventType reflect.Type, callback func(ctx context.Context, e event.DomainEvent) error) {
	r.listeners[eventType.Name()] = append(r.listeners[eventType.Name()], Listener{eventType, callback})
}
