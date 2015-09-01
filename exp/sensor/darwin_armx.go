// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin
// +build arm arm64

package sensor

import (
	"errors"
	"time"
)

type manager struct {
}

func (m *manager) initialize() {
}

func (m *manager) enable(s Sender, t Type, delay time.Duration) error {
	return errors.New("sensor: no sensors available")
}

func (m *manager) disable(t Type) error {
	return errors.New("sensor: no sensors available")
}

func (m *manager) read(e []Event) (n int, err error) {
	return 0, errors.New("sensor: no sensor data available")
}

func (m *manager) close() error {
	return nil
}