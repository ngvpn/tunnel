package gost

import (
	"io"
	"net"

	ymux "github.com/hashicorp/yamux"
)

type muxStreamConn struct {
	net.Conn
	stream io.ReadWriteCloser
}

func (c *muxStreamConn) Read(b []byte) (n int, err error) {
	return c.stream.Read(b)
}

func (c *muxStreamConn) Write(b []byte) (n int, err error) {
	return c.stream.Write(b)
}

func (c *muxStreamConn) Close() error {
	return c.stream.Close()
}

type muxSession struct {
	conn    net.Conn
	session muxSessionInterface
}

type muxSessionInterface interface {
	Open() (io.ReadWriteCloser, error)
	Accept() (io.ReadWriteCloser, error)
	Close() error
	IsClosed() bool
	NumStreams() int
}

type ymuxSession struct {
	session *ymux.Session
}

func (y *ymuxSession) Open() (io.ReadWriteCloser, error) {
	return y.session.Open()
}

func (y *ymuxSession) Accept() (io.ReadWriteCloser, error) {
	return y.session.Accept()
}

func (y *ymuxSession) Close() error {
	return y.session.Close()
}

func (y *ymuxSession) IsClosed() bool {
	return y.session.IsClosed()
}

func (y *ymuxSession) NumStreams() int {
	return y.session.NumStreams()
}

func (session *muxSession) GetConn() (net.Conn, error) {
	stream, err := session.session.Open()
	if err != nil {
		return nil, err
	}
	return &muxStreamConn{Conn: session.conn, stream: stream}, nil
}

func (session *muxSession) Accept() (net.Conn, error) {
	stream, err := session.session.Accept()
	if err != nil {
		return nil, err
	}
	return &muxStreamConn{Conn: session.conn, stream: stream}, nil
}

func (session *muxSession) Close() error {
	if session.session == nil {
		return nil
	}
	return session.session.Close()
}

func (session *muxSession) IsClosed() bool {
	if session.session == nil {
		return true
	}
	return session.session.IsClosed()
}

func (session *muxSession) NumStreams() int {
	return session.session.NumStreams()
}
