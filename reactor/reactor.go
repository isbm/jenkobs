package jenkobs_reactor

import (
	"fmt"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	"github.com/streadway/amqp"
)

type Reactor struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue

	user     string
	password string
	fqdn     string
	port     int

	wzlib_logger.WzLogger
}

func NewReactor() *Reactor {
	rtr := new(Reactor)
	return rtr
}

// SetAMQPDial string
func (rtr *Reactor) SetAMQPDial(user string, password string, fqdn string, port int) *Reactor {
	rtr.user = user
	rtr.password = password
	rtr.fqdn = fqdn
	rtr.port = port
	return rtr
}

func (rtr *Reactor) connectAMQP() error {
	if rtr.user == "" || rtr.fqdn == "" {
		err := fmt.Errorf("Error connecting to the AMQP server: user or FQDN are missing")
		rtr.GetLogger().Error(err.Error())
		return err
	}
	var err error
	var connstr string
	if rtr.port > 0 {
		connstr = fmt.Sprintf("amqps://%s:%s@%s:%d/", rtr.user, rtr.password, rtr.fqdn, rtr.port)
	} else {
		connstr = fmt.Sprintf("amqps://%s:%s@%s/", rtr.user, rtr.password, rtr.fqdn)
	}
	rtr.conn, err = amqp.Dial(connstr)
	if err != nil {
		rtr.GetLogger().Errorf("Error connecting to the AMQP server: %s", err.Error())
		return err
	} else {
		rtr.GetLogger().Infof("Connected to AMQP at %s", rtr.fqdn)
	}

	// Setup channel
	rtr.channel, err = rtr.conn.Channel()
	err = rtr.channel.ExchangeDeclarePassive("pubsub", "topic", true, false, false, false, nil)
	if err != nil {
		rtr.GetLogger().Errorf("Error creating AMQP channel: %s", err.Error())
		return err
	} else {
		rtr.GetLogger().Infof("Created AMQP channel")
	}

	// Setup queue
	rtr.queue, err = rtr.channel.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		rtr.GetLogger().Errorf("Error setting up queue: %s", err.Error())
		return err
	} else {
		rtr.GetLogger().Infof("Default queue declared")
	}

	if err = rtr.channel.QueueBind(rtr.queue.Name, "#", "pubsub", false, nil); err != nil {
		rtr.GetLogger().Errorf("Error binding queue '%s' to the channel: %s", rtr.queue.Name, err.Error())
		return err
	} else {
		rtr.GetLogger().Infof("Bound queue '%s' to the channel", rtr.queue.Name)
	}

	return nil
}

// LoadConfig of the reactor
func (rtr *Reactor) LoadActions() *Reactor {
	return rtr
}

func (rtr *Reactor) consume() error {
	msgChannel, err := rtr.channel.Consume(rtr.queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	looper := make(chan bool)

	go func() {
		rtr.GetLogger().Debug("Listening to the events...")
		for data := range msgChannel {
			fmt.Println(string(data.Body))
			fmt.Println("---")
		}
	}()

	<-looper
	return nil
}

// Run the reactor
func (rtr *Reactor) Run() error {
	if err := rtr.connectAMQP(); err == nil {
		defer rtr.conn.Close()
		if err := rtr.consume(); err != nil {
			rtr.GetLogger().Errorf("Error consuming messages: %s", err.Error())
		}
	}

	return nil
}
