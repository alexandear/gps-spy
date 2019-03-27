package storage

import (
	"fmt"
	"strings"
	"time"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/tidwall/buntdb"

	"github.com/devchallenge/spy-api/internal/model"
)

type Storage struct {
	db *buntdb.DB
}

func New(db *buntdb.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func key(number string) string {
	return fmt.Sprintf("%s:%s", number, uuid.New())
}

func (s *Storage) Save(phone model.Phone, coordinate model.Coordinate, timestamp time.Time) error {
	item := &item{
		IMEI:       phone.IMEI,
		IP:         phone.IP,
		Coordinate: coordinate,
		Timestamp:  timestamp,
	}
	if err := s.db.Update(func(tx *buntdb.Tx) error {
		saved, err := item.Save()
		if err != nil {
			return errors.Wrap(err, "failed to save item")
		}
		if _, _, err := tx.Set(key(phone.Number), saved, nil); err != nil {
			return errors.Wrap(err, "failed to set")
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to db update")
	}
	return nil
}

func (s *Storage) Read(number string) ([]*item, error) {
	var items []*item
	if err := s.db.View(func(tx *buntdb.Tx) error {
		return tx.Ascend("", func(key, value string) bool {
			if strings.HasPrefix(key, number) {
				item := &item{}
				if err := item.Load(value); err != nil {
					return false
				}
				items = append(items, item)
			}

			return true
		})
	}); err != nil {
		return nil, errors.Wrap(err, "failed to view db")
	}
	return items, nil
}
