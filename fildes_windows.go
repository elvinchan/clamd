package clamd

import (
	"errors"
	"net"
	"net/textproto"
)

const fildesUnsupportErr = "Fildes is not supported"

func (c *Client) fildesScan(_ *textproto.Conn, _ net.Conn, _ string) (err error) {
	return errors.New(fildesUnsupportErr)
}
