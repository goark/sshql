package pgdrv

import (
	"database/sql"
	"database/sql/driver"
	"net"
	"time"

	"github.com/lib/pq"
)

const DriverName = "postgres+ssh"

// Driver is driver.Driver and pq.Dialer for PostgreSQL via SSH.
type Driver struct {
	Dialer
}

var _ pq.Dialer = (*Driver)(nil)
var _ driver.Driver = (*Driver)(nil)

// New returns new Driver instance.
func New(d Dialer) *Driver {
	return &Driver{d}
}

// Open opens connection to the server (compatible driver.Driver interface).
func (d *Driver) Open(s string) (driver.Conn, error) {
	if err := d.Connect(); err != nil {
		return nil, err
	}
	return pq.DialOpen(d, s)
}

// DialTimeout makes socket connection via SSH (compatible pq.Dialer interface).
func (d *Driver) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return d.Dial(network, address)
}

// Register makes a database driver available by name "postgres+ssh".
func (d *Driver) Register(name string) {
	if len(name) == 0 {
		name = DriverName
	}
	sql.Register(name, d)
}

/* MIT License
 *
 * Copyright 2022 Spiegel (forked from github.com/mattn/pqssh package)
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
