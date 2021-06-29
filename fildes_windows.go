package clamd

import (
	"errors"
	"net"
	"net/textproto"
)

const fildesUnsupportErr = "Fildes is not supported"

func (c *Client) fildesScan(tc *textproto.Conn, conn net.Conn, p string) (err error) {
	return errors.New(fildesUnsupportErr)
}
