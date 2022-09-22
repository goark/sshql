package mysqldrv

import (
	"context"
	"net"

	"github.com/go-sql-driver/mysql"
	"github.com/goark/errs"
)

const DialName = "ssh+tcp"

// Driver is driver.Driver and pq.Dialer for MySQL via SSH.
type Driver struct {
	Dialer
}

// New returns new Driver instance.
func New(d Dialer) *Driver {
	return &Driver{d}
}

// Dial makes socket connection via SSH (compayible mysql/RegisterDial type).
func (d *Driver) Dial(address string) (net.Conn, error) {
	if err := d.Dialer.Connect(); err != nil {
		return nil, errs.Wrap(err)
	}
	return d.Dialer.Dial("tcp", address)
}

// DialContext makes socket connection via SSH (compayible mysql/RegisterDialContext type).
func (d *Driver) DialContext(_ context.Context, address string) (net.Conn, error) {
	return d.Dial(address)
}

// Register makes a dial available by name "mysql+ssh".
func (d *Driver) RegisterDial(name string) {
	if len(name) == 0 {
		name = DialName
	}
	mysql.RegisterDialContext(name, d.DialContext)
}

/* MIT License
 *
 * Copyright 2022 Spiegel
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
