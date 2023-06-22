package eventconsumer

import (
	"log"
	"time"

	"github.com/markraiter/bot/pkg/events"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())
			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(time.Second * 1)
			continue
		}

		if err := c.HandleEvents(gotEvents); err != nil {
			log.Print(err)
			continue
		}
	}
}

func (c *Consumer) HandleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new even^ %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())
			continue
		}
	}

	return nil
}
