package pgdrv

import (
	"database/sql"
	"database/sql/driver"
	"net"
	"time"

	"github.com/goark/sshql"
	"github.com/lib/pq"
)

const DriverName = "postgres+ssh"

// Driver is driver.Driver and pq.Dialer for PostgreSQL via SSH.
type Driver struct {
	sshql.Dialer
}

// var _ pq.Dialer = (*Driver)(nil)
// var _ driver.Driver = (*Driver)(nil)

// Open opens connection to the server.
func (d *Driver) Open(s string) (driver.Conn, error) {
	if err := d.Connect(); err != nil {
		return nil, err
	}
	return pq.DialOpen(d, s)
}

// Dial makes socket connection via SSH.
func (d *Driver) Dial(network, address string) (net.Conn, error) {
	return d.Dialer.Dial(network, address)
}

// DialTimeout makes socket connection via SSH.
func (d *Driver) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return d.Dialer.Dial(network, address)
}

// Register makes a database driver available by name "postgres+ssh".
func (d *Driver) Register() {
	sql.Register(DriverName, d)
}

/**
 * Frok from github.com/mattn/pqssh package
 */
