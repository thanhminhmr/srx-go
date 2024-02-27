/*
 * srx: The fast Symbol Ranking based compressor.
 * Copyright (C) 2023-2024  Mai Thanh Minh (a.k.a. thanhminhmr)
 *
 * This program is free software: you can redistribute it and/or modify it under
 * the terms of the GNU General Public License as published by the Free Software
 * Foundation, either  version 3 of the  License,  or (at your option) any later
 * version.
 *
 * This program  is distributed in the hope  that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
 * FOR  A PARTICULAR PURPOSE. See  the  GNU  General  Public   License  for more
 * details.
 *
 * You should have received a copy of the GNU General Public License along with
 * this program. If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"bufio"
	"sync/atomic"
)

type CountingReader struct {
	*bufio.Reader
	count uint64
}

func NewCountingReader(reader *bufio.Reader) *CountingReader {
	return &CountingReader{
		Reader: reader,
	}
}

func (r *CountingReader) Read(buf []byte) (int, error) {
	n, err := r.Reader.Read(buf)
	atomic.AddUint64(&r.count, uint64(n))
	return n, err
}

func (r *CountingReader) Count() uint64 {
	return atomic.LoadUint64(&r.count)
}

type CountingWriter struct {
	*bufio.Writer
	count uint64
}

func NewCountingWriter(writer *bufio.Writer) *CountingWriter {
	return &CountingWriter{
		Writer: writer,
	}
}

func (r *CountingWriter) Write(buf []byte) (int, error) {
	n, err := r.Writer.Write(buf)
	if n >= 0 {
		atomic.AddUint64(&r.count, uint64(n))
	}
	return n, err
}

func (r *CountingWriter) Count() uint64 {
	return atomic.LoadUint64(&r.count)
}
