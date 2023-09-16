// Copyright 2016 The go-aigccoin Authors
// This file is part of the go-aigccoin library.
//
// The go-aigccoin library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-aigccoin library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-aigccoin library. If not, see <http://www.gnu.org/licenses/>.

//go:build !go1.5
// +build !go1.5

// no-op implementation of tracing methods for Go < 1.5.

package debug

import "errors"

func (*HandlerT) StartGoTrace(string) error {
	return errors.New("tracing is not supported on Go < 1.5")
}

func (*HandlerT) StopGoTrace() error {
	return errors.New("tracing is not supported on Go < 1.5")
}
