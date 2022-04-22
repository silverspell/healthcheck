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

	go func(rc chan string) {
		t := time.NewTimer(2 * time.Second)
		<-t.C
		rc <- "timeout"
	}(receiveChan)

	sendChan <- "Are you alive"

	for {
		msg := <-receiveChan
		if msg == "timeout" {
			return fmt.Errorf("timeout")
		}
		return nil
	}
}
