// Copyright 2015, Cyrill @ Schumacher.fm and the CoreStore contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mail

import (
	"errors"
	"github.com/corestoreio/csfw/utils/log"
	"github.com/go-gomail/gomail"
	"time"
)

var ErrMailChannelClosed = errors.New("The mail channel has been closed.")

// Daemon represents a daemon which must be created via NewDaemon() function
type Daemon struct {
	msgChan chan *gomail.Message
	dialer  *gomail.Dialer
	closed  bool
	// SMTPTimeout closes the connection to the SMTP server if no email was
	// sent in the last default 30 seconds.
	SMTPTimeout int
}

// Start listens to a channel and sends all incoming messages. Errors will be logged.
// Use code snippet:
//		d := NewDaemon(...)
// 		go func(){
//			if err := d.Worker(); err != nil {
// 				panic(err) // for example
// 			}
// 		}()
//		d.Send(*gomail.Message)
//		d.Stop()
func (dm *Daemon) Worker() error {
	if dm.closed {
		return ErrMailChannelClosed
	}

	var s gomail.SendCloser
	var err error
	open := false
	for {
		select {
		case m, ok := <-dm.msgChan:
			if !ok {
				dm.closed = true
				return nil
			}
			if !open {
				if s, err = dm.dialer.Dial(); err != nil {
					return log.Error("mail.daemon.Start.Dial", "err", err, "message", m)
				}
				open = true
			}
			if err := gomail.Send(s, m); err != nil {
				log.Error("mail.daemon.Start.Send", "err", err, "message", m)
			}
		// Close the connection to the SMTP server if no email was sent in
		// the last 30 seconds.
		case <-time.After(dm.SMTPTimeout * time.Second):
			if open {
				if err := s.Close(); err != nil {
					return log.Error("mail.daemon.Start.Close", "err", err, "message", m)
				}
				open = false
			}
		}
	}

}

// Stop closes the channel stops the daemon
func (dm *Daemon) Stop() error {
	if dm.closed {
		return ErrMailChannelClosed
	}
	close(dm.msgChan)
	dm.closed = true
	return nil
}

// Send sends a mail
func (dm *Daemon) Send(m *gomail.Message) error {
	if dm.closed {
		return ErrMailChannelClosed
	}
	dm.msgChan <- m
	return nil
}

// SendPlain sends a simple plain text email
func (dm *Daemon) SendPlain(from, to, subject, body string) error {
	if dm.closed {
		return ErrMailChannelClosed
	}
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	dm.Send(m)
	return nil
}

// Options applies optional arguments to the daemon
// struct. It returns the last set option. More info about the returned function:
// http://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
func (dm *Daemon) Option(opts ...DaemonOption) (previous DaemonOption) {
	for _, o := range opts {
		if o != nil {
			previous = o(dm)
		}
	}
	return previous
}

// DaemonOption can be used as an argument in NewDaemon to configure a daemon.
type DaemonOption func(*Daemon) DaemonOption

// DefaultDialer connects to localhost on port 25.
var DefaultDialer = gomail.NewPlainDialer("localhost", "25", "", "")

// SetMessageChannel sets your custom channel to listen to.
func SetMessageChannel(mailChan chan *gomail.Message) DaemonOption {
	return func(da *Daemon) DaemonOption {
		previous := da.msgChan
		da.msgChan = mailChan
		da.closed = false
		return SetMessageChannel(previous)
	}
}

// SetDialer sets a channel to listen to.
func SetDialer(di *gomail.Dialer) DaemonOption {
	if di == nil {
		di = DefaultDialer
	}
	return func(da *Daemon) DaemonOption {
		previous := da.dialer
		da.dialer = di
		return SetDialer(previous)
	}
}

// NewDaemon creates a new daemon to send default to localhost:25 and creates
// a default unbuffered channel which can be used via the Send*() function.
func NewDaemon(opts ...DaemonOption) *Daemon {
	d := &Daemon{
		dialer:      DefaultDialer,
		SMTPTimeout: 30,
	}
	d.Option(opts...)
	if d.msgChan == nil {
		d.msgChan = make(chan *gomail.Message)
	}
	return d
}
