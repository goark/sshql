package sshql

import (
	"errors"
	"testing"
)

func TestNil(t *testing.T) {
	var d *Dialer // initialize by nil
	err := d.Connect()
	if err == nil || !errors.Is(err, ErrNoConnection) {
		t.Errorf("<nil>.Connect()  = '%v', want '%v'.", err, ErrNoConnection)
	}
	_, err = d.Dial("tcp", "")
	if err == nil || !errors.Is(err, ErrNoConnection) {
		t.Errorf("<nil>.Connect()  = '%v', want '%v'.", err, ErrNoConnection)
	}
	err = d.Close()
	if err != nil {
		t.Errorf("<nil>.Close()  = '%v', want <nil>.", err)
	}
}

func TestEmpty(t *testing.T) {
	d := &Dialer{}
	err := d.Connect()
	if err == nil {
		t.Errorf("<nil>.Connect()  = '%v', do not want <nil>.", err)
	}
	_, err = d.Dial("tcp", "")
	if err == nil || !errors.Is(err, ErrNoConnection) {
		t.Errorf("<nil>.Connect()  = '%v', want '%v'.", err, ErrNoConnection)
	}
	err = d.Close()
	if err != nil {
		t.Errorf("<nil>.Close()  = '%v', want <nil>.", err)
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
