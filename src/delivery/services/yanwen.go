package services

import (
	"time"

	"github.com/gocolly/colly"
)

type Yanwen struct {
	historic []Event
	token    string
}

func (fb *Yanwen) SetToken(t string) {
	fb.token = t
}

func (fb Yanwen) GetLastEvent() string {
	if len(fb.historic) == 0 {
		return "Error"
	}
	return fb.historic[0].event
}

func (box *Yanwen) AddItem(e Event) []Event {
	box.historic = append(box.historic, e)
	return box.historic
}

func (fb *Yanwen) Get() []Event {
	c := colly.NewCollector()
	var events []Event

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		var ts time.Time
		var ev string
		e.ForEach("td", func(i int, elem *colly.HTMLElement) {
			if i%2 == 0 {
				ts, _ = time.Parse("2006-01-02 03:04", elem.Text)

			} else if i%2 == 1 {
				ev = elem.Text
				events = append(events, Event{timestamp: ts, event: ev})
			}
		})
	})
	m := make(map[string]string)
	m["InputTrackNumbers"] = fb.token
	c.Post("http://track.yw56.com.cn/en-US/", m)
	return events
}

func (fb *Yanwen) Update() bool {
	data := fb.Get()
	if len(fb.historic) == 0 && len(data) != 0 {
		fb.historic = data
	} else if len(fb.historic) != len(data) && len(data) != 0 {
		fb.historic = data
		return true
	}
	return false
}
