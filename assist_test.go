// -----------------------------------------------------------------------------
// github.com/balacode/udpt                                    /[assist_test.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package udpt

import (
	"net"
	"strings"
)

// makeTestConn creates a UDP connection for testing.
func makeTestConn() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9876")
	if err != nil {
		panic("0xEE52A7")
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic("0xE1E9E7")
	}
	return conn
} //                                                                makeTestConn

// matchError retruns true if err contains the specified error message.
func matchError(err error, msg string) bool {
	if err == nil && (msg == "" || msg == "nil" || msg == "<nil>") {
		return true
	}
	return err != nil && strings.Contains(err.Error(), msg)
} //                                                                  matchError

// -----------------------------------------------------------------------------

// mockWriteCloser is a mock io.WriteCloser with methods you can make fail.
type mockWriteCloser struct {
	failWrite bool
	failClose bool
} //                                                             mockWriteCloser

// Write is a method of mockWriteCloser implementing io.WriteCloser.
//
// You can make it return an error by setting mockWriteCloser.failWrite.
//
func (ob *mockWriteCloser) Write(p []byte) (n int, err error) {
	if ob.failWrite {
		return 0, makeError(0xE12345, "from mockWriteCloser.Write")
	}
	return len(p), nil
} //                                                                       Write

// Close is a method of mockWriteCloser implementing io.WriteCloser.
//
// You can make it return an error by setting mockWriteCloser.failClose.
//
func (ob *mockWriteCloser) Close() error {
	if ob.failClose {
		return makeError(0xE12345, "from mockWriteCloser.Close")
	}
	return nil
} //                                                                       Close

// end
