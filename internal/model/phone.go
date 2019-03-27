package model

import (
	"net"

	"github.com/pkg/errors"
)

type Phone struct {
	Number string
	IMEI   string
	IP     net.IP
}

func NewPhone(number string, ip net.IP, imei string) (Phone, error) {
	if number == "" {
		return Phone{}, errors.Wrap(ErrInvalidArgument, "number must be not empty")
	}
	if imei == "" {
		return Phone{}, errors.Wrap(ErrInvalidArgument, "IMEI must be not empty")
	}
	return Phone{
		Number: number,
		IMEI:   imei,
		IP:     ip,
	}, nil
}
