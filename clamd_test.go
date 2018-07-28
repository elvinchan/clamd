// Copyright (C) 2018 Andrew Colin Kissa <andrew@datopdog.io>
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

/*
Package clamd Golang Clamd client
Clamd - Golang clamd client
*/
package clamd

import (
	"fmt"
	"go/build"
	"os"
	"path"
	"strings"
	"testing"
	"time"
)

type checkErrorTestKey struct {
	in  string
	out error
}

var s = "Could not open file /.xxxx ERROR"
var errNf = fmt.Errorf("%s", strings.TrimRight(s, " ERROR"))
var TestcheckErrors = []checkErrorTestKey{
	{"This is a test", nil},
	{s, errNf},
}

func TestCheckErrors(t *testing.T) {
	for _, tt := range TestcheckErrors {
		if e := checkError(tt.in); e != tt.out {
			if e != nil && e.Error() != tt.out.Error() {
				t.Errorf("%q = checkError(%q), want %q", tt.out, tt.in, tt.out)
			}
		}
	}
}

func TestBasics(t *testing.T) {
	// Test Non existant socket
	_, e := NewClient("unix", "/tmp/.dumx.sock")
	if e == nil {
		t.Errorf("An error should be returned as sock does not exist")
	} else {
		expected := "The unix socket: /tmp/.dumx.sock does not exist"
		if e.Error() != expected {
			t.Errorf("Expected %q want %q", expected, e)
		}
	}
	// Test defaults
	_, e = NewClient("", "")
	if e == nil {
		t.Errorf("An error should be returned as sock does not exist")
	} else {
		expected := "The unix socket: /var/run/clamav/clamd.sock does not exist"
		if e.Error() != expected {
			t.Errorf("Got %q want %q", expected, e)
		}
	}
	// Test udp
	_, e = NewClient("udp", "127.1.1.1:3310")
	if e == nil {
		t.Errorf("Expected an error got nil")
	} else {
		expected := "Protocol: udp is not supported"
		if e.Error() != expected {
			t.Errorf("Got %q want %q", expected, e)
		}
	}
	// Test tcp
	network := "tcp"
	address := "127.1.1.1:3310"
	c, e := NewClient(network, address)
	if e != nil {
		t.Errorf("An error should not be returned")
	} else {
		if c.network != network {
			t.Errorf("Got %q want %q", c.network, network)
		}
		if c.address != address {
			t.Errorf("Got %q want %q", c.address, address)
		}
	}
}

func TestSettings(t *testing.T) {
	var e error
	var c *Client
	network := "tcp"
	address := "127.1.1.1:3310"
	if c, e = NewClient(network, address); e != nil {
		t.Errorf("An error should not be returned")
	}
	if c.connTimeout != defaultTimeout {
		t.Errorf("The default conn timeout should be set")
	}
	if c.connSleep != defaultSleep {
		t.Errorf("The default conn sleep should be set")
	}
	if c.connRetries != 0 {
		t.Errorf("The default conn retries should be set")
	}
	expected := 2 * time.Second
	c.SetConnTimeout(expected)
	if c.connTimeout != expected {
		t.Errorf("Calling c.SetConnTimeout(%q) failed", expected)
	}
	c.SetCmdTimeout(expected)
	if c.cmdTimeout != expected {
		t.Errorf("Calling c.SetCmdTimeout(%q) failed", expected)
	}
	c.SetConnSleep(expected)
	if c.connSleep != expected {
		t.Errorf("Calling c.SetConnSleep(%q) failed", expected)
	}
	c.SetConnRetries(2)
	if c.connRetries != 2 {
		t.Errorf("Calling c.SetConnRetries(%q) failed", 2)
	}
	c.SetConnRetries(-2)
	if c.connRetries != 0 {
		t.Errorf("Preventing negative values in c.SetConnRetries(%q) failed", -2)
	}
}

func TestMethodsErrors(t *testing.T) {
	var e error
	var c *Client
	network := "tcp"
	address := "127.1.1.1:3310"
	if c, e = NewClient(network, address); e != nil {
		t.Errorf("An error should not be returned")
	}
	c.SetConnTimeout(500 * time.Microsecond)
	// c.SetConnRetries(1)
	expected := "dial tcp 127.1.1.1:3310: i/o timeout"
	if _, e = c.Ping(); e == nil {
		t.Errorf("An error should be returned")
	} else {
		if e.Error() != expected {
			t.Errorf("Got %q want %q", e, expected)
		}
	}
	if _, e = c.Version(); e == nil {
		t.Errorf("An error should be returned")
	} else {
		if e.Error() != expected {
			t.Errorf("Got %q want %q", e, expected)
		}
	}
	if _, e = c.Reload(); e == nil {
		t.Errorf("An error should be returned")
	} else {
		if e.Error() != expected {
			t.Errorf("Got %q want %q", e, expected)
		}
	}
	if e = c.Shutdown(); e == nil {
		t.Errorf("An error should be returned")
	} else {
		if e.Error() != expected {
			t.Errorf("Got %q want %q", e, expected)
		}
	}
	if _, e = c.Stats(); e == nil {
		t.Errorf("An error should be returned")
	} else {
		if e.Error() != expected {
			t.Errorf("Got %q want %q", e, expected)
		}
	}
	if _, e = c.VersionCmds(); e == nil {
		t.Errorf("An error should be returned")
	} else {
		if e.Error() != expected {
			t.Errorf("Got %q want %q", e, expected)
		}
	}
	if _, e = c.Scan("/tmp/bxx.syx"); e == nil {
		t.Errorf("An error should be returned")
	} else {
		if e.Error() != expected {
			t.Errorf("Got %q want %q", e, expected)
		}
	}
	if _, e = c.ContScan("/tmp/bxx.syx"); e == nil {
		t.Errorf("An error should be returned")
	} else {
		if e.Error() != expected {
			t.Errorf("Got %q want %q", e, expected)
		}
	}
	if _, e = c.MultiScan("/tmp/bxx.syx"); e == nil {
		t.Errorf("An error should be returned")
	} else {
		if e.Error() != expected {
			t.Errorf("Got %q want %q", e, expected)
		}
	}
	expected = "stat /tmp/bxx.syx: no such file or directory"
	if _, e = c.InStream("/tmp/bxx.syx"); e == nil {
		t.Errorf("An error should be returned")
	} else {
		if e.Error() != expected {
			t.Errorf("Got %q want %q", e, expected)
		}
	}
	if _, e = c.Fildes("/tmp/bxx.syx"); e == nil {
		t.Errorf("An error should be returned")
	} else {
		if e.Error() != expected {
			t.Errorf("Got %q want %q", e, expected)
		}
	}
}

func TestMethods(t *testing.T) {
	var e error
	var b bool
	var c *Client
	var result []*Response
	var vcmds []string
	var network, address, rs string

	address = os.Getenv("CLAMD_ADDRESS")
	if address == "" {
		address = "/opt/local/var/run/clamav/clamd.socket"
		if _, e = os.Stat(address); os.IsNotExist(e) {
			address = "/var/run/clamav/clamd.ctl"
		}
	}
	if strings.HasPrefix(address, "/") {
		network = "unix"
	} else {
		network = "tcp4"
	}
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	fn := path.Join(gopath, "src/github.com/baruwa-enterprise/clamd/examples/eicar.txt")
	if c, e = NewClient(network, address); e != nil {
		t.Errorf("An error should not be returned")
	}
	if b, e = c.Ping(); e != nil {
		t.Errorf("An error should not be returned")
	} else {
		if !b {
			t.Errorf("Expected %t got %t", true, b)
		}
	}
	if rs, e = c.Version(); e != nil {
		t.Errorf("An error should not be returned")
	} else {
		if !strings.HasPrefix(rs, "Clam") {
			t.Errorf("Expected version starting with Clam, got %q", rs)
		}
	}
	if rs, e = c.Stats(); e != nil {
		t.Errorf("An error should not be returned")
	} else {
		if !strings.HasPrefix(rs, "POOLS:") {
			t.Errorf("Expected version starting with POOLS:, got %q", rs)
		}
	}
	if vcmds, e = c.VersionCmds(); e != nil {
		t.Errorf("An error should not be returned")
	} else {
		if len(vcmds) == 0 {
			t.Errorf("Expected a slice of strings:, got %q", vcmds)
		} else {
			if vcmds[0] != "SCAN" {
				t.Errorf("Expected SCAN:, got %q", vcmds[0])
			}
		}
	}
	if result, e = c.Scan(fn); e != nil {
		t.Errorf("An error should not be returned")
	} else {
		l := len(result)
		if l == 0 {
			t.Errorf("Expected a slice of Response objects:, got %q", result)
		} else if l > 1 {
			t.Errorf("Expected a slice of Response 1 object:, got %d", l)
		} else {
			mb := result[0]
			if mb.Filename != fn {
				t.Errorf("Expected %q, got %q", fn, mb.Filename)
			}
			if mb.Signature != "Eicar-Test-Signature" {
				t.Errorf("Expected %q, got %q", "Eicar-Test-Signature", mb.Signature)
			}
		}
	}
	if result, e = c.ContScan(path.Dir(fn)); e != nil {
		t.Errorf("An error should not be returned")
	} else {
		l := len(result)
		if l == 0 {
			t.Errorf("Expected a slice of Response objects:, got %q", result)
		} else if l > 2 {
			t.Errorf("Expected a slice of Response 2 objects:, got %d", l)
		} else {
			mb := result[0]
			if mb.Filename != fn {
				t.Errorf("Expected %q, got %q", fn, mb.Filename)
			}
			if mb.Signature != "Eicar-Test-Signature" {
				t.Errorf("Expected %q, got %q", "Eicar-Test-Signature", mb.Signature)
			}
			mb = result[1]
			if mb.Signature != "Eicar-Test-Signature" {
				t.Errorf("Expected %q, got %q", "Eicar-Test-Signature", mb.Signature)
			}
		}
	}
	if result, e = c.MultiScan(fn); e != nil {
		t.Errorf("An error should not be returned")
	} else {
		l := len(result)
		if l == 0 {
			t.Errorf("Expected a slice of Response objects:, got %q", result)
		} else if l > 1 {
			t.Errorf("Expected a slice of Response 1 object:, got %q", l)
		} else {
			mb := result[0]
			if mb.Filename != fn {
				t.Errorf("Expected %q, got %q", fn, mb.Filename)
			}
			if mb.Signature != "Eicar-Test-Signature" {
				t.Errorf("Expected %q, got %q", "Eicar-Test-Signature", mb.Signature)
			}
		}
	}
	if result, e = c.InStream(fn); e != nil {
		t.Errorf("An error should not be returned")
	} else {
		l := len(result)
		if l == 0 {
			t.Errorf("Expected a slice of Response objects:, got %q", result)
		} else if l > 1 {
			t.Errorf("Expected a slice of Response 1 object:, got %q", l)
		} else {
			mb := result[0]
			if mb.Filename != "stream" {
				t.Errorf("Expected %q, got %q", "stream", mb.Filename)
			}
			if mb.Signature != "Eicar-Test-Signature" {
				t.Errorf("Expected %q, got %q", "Eicar-Test-Signature", mb.Signature)
			}
		}
	}
	if network == "unix" {
		if result, e = c.Fildes(fn); e != nil {
			t.Errorf("An error should not be returned")
		} else {
			l := len(result)
			if l == 0 {
				t.Errorf("Expected a slice of Response objects:, got %q", result)
			} else if l > 1 {
				t.Errorf("Expected a slice of Response 1 object:, got %q", l)
			} else {
				mb := result[0]
				if !strings.HasPrefix(mb.Filename, "fd[") {
					t.Errorf("Expected name starting with fd[, got %q", mb.Filename)
				}
				if mb.Signature != "Eicar-Test-Signature" {
					t.Errorf("Expected %q, got %q", "Eicar-Test-Signature", mb.Signature)
				}
			}
		}
	}
	if b, e = c.Reload(); e != nil {
		t.Errorf("An error should not be returned")
	} else {
		if !b {
			t.Errorf("Expected true, got %t", b)
		}
	}
}
