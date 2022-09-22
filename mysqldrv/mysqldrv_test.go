package mysqldrv

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
	"testing"
)

var ErrConnect = errors.New("error for Connect test")

type dummyDialer struct{}

func (d *dummyDialer) Connect() error {
	return ErrConnect
}
func (d *dummyDialer) Dial(network, address string) (net.Conn, error) {
	return nil, fmt.Errorf("error for Dial test: network = %q, address = %q", network, address)
}
func (d *dummyDialer) Close() error {
	return errors.New("error for Close test")
}

func TestNil(t *testing.T) {
	var d *dummyDialer // initialize by nil
	drv := New(d)
	drv.RegisterDial("")
	db, err := sql.Open("mysql", fmt.Sprintf("dbuser:dbpassword@%s(localhost:3306)/dbname", DialName))
	if err != nil {
		t.Errorf("sql.Open()  = '%v', want <nil>.", err)
	}
	_, err = db.Query("SELECT id, name FROM tablename ORDER BY id")
	if !errors.Is(err, ErrConnect) {
		t.Errorf("<nil>.Open()  = '%v', want '%v'.", err, ErrConnect)
	}
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
