package healthcheckmodule

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/silverspell/rabbitmodule"
)

type RabbitConnection struct{}

func (r *RabbitConnection) Connect() error {
	uniqueChan, _ := uuid.NewUUID()

	sendChan, receiveChan := make(chan string), make(chan string)

	go rabbitmodule.ConnectSubscriber(receiveChan, uniqueChan.String())
	go rabbitmodule.ConnectPublisher(sendChan, uniqueChan.String())
	var err error
	sendChan <- "Are you alive"
	loop := true
	for loop {
		select {
		case <-time.After(time.Duration(1 * time.Second)):
			err = fmt.Errorf("rabbit timeout")
			loop = false
		case <-receiveChan:
			err = nil
			loop = false
		}
	}
	return err
}
