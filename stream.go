package tesla

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// StreamingMessage represents the current state of the car.
type StreamingMessage struct {
	Timestamp    time.Time
	Speed        int
	Odometer     float64
	SOC          int
	Elevation    int
	EstHeading   int
	EstLatitude  float64
	EstLongitude float64
	Power        int
	ShiftState   int
	Range        int
	EstRange     int
	Heading      int
}

func (msg *StreamingMessage) fromCSV(data string) {
	values := strings.Split(data, ",")

	if ts, err := strconv.Atoi(values[0]); err == nil {
		sec := ts / 1000
		ms := time.Duration(ts%1000) * time.Millisecond

		msg.Timestamp = time.Unix(int64(sec), int64(ms))

	}

	if val, err := strconv.Atoi(values[1]); err == nil {
		msg.Speed = val
	}

	if val, err := strconv.ParseFloat(values[2], 64); err == nil {
		msg.Odometer = val
	}

	if val, err := strconv.Atoi(values[3]); err == nil {
		msg.SOC = val
	}

	if val, err := strconv.Atoi(values[4]); err == nil {
		msg.Elevation = val
	}

	if val, err := strconv.Atoi(values[5]); err == nil {
		msg.EstHeading = val
	}

	if val, err := strconv.ParseFloat(values[6], 64); err == nil {
		msg.EstLatitude = val
	}

	if val, err := strconv.ParseFloat(values[7], 64); err == nil {
		msg.EstLongitude = val
	}

	if val, err := strconv.Atoi(values[8]); err == nil {
		msg.Power = val
	}

	if val, err := strconv.Atoi(values[9]); err == nil {
		msg.ShiftState = val
	}

	if val, err := strconv.Atoi(values[10]); err == nil {
		msg.Range = val
	}

	if val, err := strconv.Atoi(values[11]); err == nil {
		msg.EstRange = val
	}

	if val, err := strconv.Atoi(values[12]); err == nil {
		msg.Heading = val
	}
}

type Stream struct {
	data  chan StreamingMessage
	close bool

	sem sync.RWMutex

	err error
}

func (s *Stream) Close() {
	s.sem.Lock()
	defer s.sem.Unlock()

	s.close = true
}

func (s *Stream) Err() error {
	return s.err
}

func (s *Stream) Data() <-chan StreamingMessage {
	return s.data
}

// Stream will initiate a stream of data from the car, with updates going to the returned channel.
// New messages are received approximately every 250ms, but that is not reliable. If the API closes
// the stream, the returned channel will be closed as well.
func (c *Conn) Stream(id int, token string) (*Stream, error) {
	type message struct {
		MessageType string `json:"msg_type"`
		Token       string `json:"token"`
		Value       string `json:"value"`
		Tag         string `json:"tag"`
	}

	connect := func() (*websocket.Conn, error) {
		connectMsg := message{
			MessageType: "data:subscribe_oauth",
			Token:       c.accessToken,
			Value:       "speed,odometer,soc,elevation,est_heading,est_lat,est_lng,power,shift_state,range,est_range,heading",
			Tag:         fmt.Sprintf("%d", id),
		}

		url := "wss://streaming.vn.teslamotors.com/streaming/"
		ws, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return nil, err
		}
		if err = ws.WriteJSON(connectMsg); err != nil {
			return nil, err
		}

		return ws, nil
	}

	ws, err := connect()
	if err != nil {
		return nil, err
	}

	stream := &Stream{
		data: make(chan StreamingMessage, 10),
	}

	go func(s *Stream) {
		defer close(s.data)

		for {
			if s.close {
				return
			}

			var msg message
			err := ws.ReadJSON(&msg)

			if msg.MessageType == "data:update" {
				var sm StreamingMessage
				sm.fromCSV(msg.Value)
				s.data <- sm
			} else if msg.MessageType == "data:error" {
				ws.Close()

				ws, err = connect()
				if err != nil {
					s.err = err
					return
				}
			}
		}
	}(stream)

	return stream, nil
}
