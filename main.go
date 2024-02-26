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
	"fmt"
	"os"
)

func help() {
	fmt.Print(`
srx: The fast Symbol Ranking based compressor, version 0.0.0-go
Copyright (C) 2023-2024  Mai Thanh Minh (a.k.a. thanhminhmr)

To   compress: srx c <input-file> <output-file>
To decompress: srx d <input-file> <output-file>
`)
	os.Exit(0)
}

func main() {
	var args = os.Args
	if len(args) != 4 {
		help()
	}

	// todo
}
