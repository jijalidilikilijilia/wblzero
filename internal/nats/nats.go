package nats

import (
	"github.com/nats-io/stan.go"
)

type NatsHandler struct {
	conn stan.Conn
}

func NewNatsHandler(clusterID, clientID, natsURL string) (*NatsHandler, error) {
	conn, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		return nil, err
	}

	return &NatsHandler{conn: conn}, nil
}

func (nh *NatsHandler) Subscribe(subject, queueGroup string, handler stan.MsgHandler) (stan.Subscription, error) {
	return nh.conn.QueueSubscribe(subject, queueGroup, handler)
}

func (nh *NatsHandler) Publish(subject string, data []byte) error {
	if err := nh.conn.Publish(subject, data); err != nil {
		return err
	}
	return nil
}

func (nh *NatsHandler) Close() {
	nh.conn.Close()
}
