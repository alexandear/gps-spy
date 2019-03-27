package storage

import (
	"encoding/json"
	"net"
	"time"

	"github.com/devchallenge/spy-api/internal/model"
)

type item struct {
	IMEI       string
	IP         net.IP
	Coordinate model.Coordinate
	Timestamp  time.Time
}

func (p *item) Save() (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func (p *item) Load(value string) error {
	return json.Unmarshal([]byte(value), &p)
}
