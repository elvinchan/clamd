//+build !windows

package clamd

import (
	"fmt"
	"net"
	"net/textproto"
	"os"
	"syscall"

	"github.com/elvinchan/clamd/protocol"
)

func (c *Client) fildesScan(tc *textproto.Conn, conn net.Conn, p string) (err error) {
	var f *os.File
	var vf *os.File

	fmt.Fprintf(tc.W, "n%s\n", protocol.Fildes)
	tc.W.Flush()

	if f, err = os.Open(p); err != nil {
		return
	}
	defer f.Close()

	s := conn.(*net.UnixConn)
	if vf, err = s.File(); err != nil {
		return
	}
	sock := int(vf.Fd())
	defer vf.Close()

	fds := []int{int(f.Fd())}
	rights := syscall.UnixRights(fds...)
	if err = syscall.Sendmsg(sock, nil, rights, nil, 0); err != nil {
		return
	}

	return
}
