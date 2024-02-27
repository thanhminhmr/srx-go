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
	"fmt"
	"log"
	"os"
	"time"
)

func compress(reader *CountingReader, writer *CountingWriter) error {
	// todo
	return nil
}

func decompress(reader *CountingReader, writer *CountingWriter) error {
	// todo
	return nil
}

const HelpString = `
srx: The fast Symbol Ranking based compressor, version 0.0.0-go
Copyright (C) 2023-2024  Mai Thanh Minh (a.k.a. thanhminhmr)

To   compress: srx c <input-file> <output-file>
To decompress: srx d <input-file> <output-file>
`

func main() {
	// check argument
	var args = os.Args
	if len(args) != 4 || args[1] != "c" && args[1] != "d" {
		log.Fatal(HelpString)
	}

	var isCompressing = args[1] == "c"
	var inputFilePath = args[2]
	var outputFilePath = args[3]

	// open input file
	var inputFile, inputFileError = os.Open(inputFilePath)
	if inputFileError != nil {
		log.Fatal(inputFileError)
	}
	var countingInput = NewCountingReader(bufio.NewReader(inputFile))

	// create output file
	var outputFile, outputFileError = os.Create(outputFilePath)
	if outputFileError != nil {
		log.Fatal(outputFileError)
	}
	var countingOutput = NewCountingWriter(bufio.NewWriter(outputFile))

	// do compress/decompress
	var startTime = time.Now()

	var compressDecompressError error
	if isCompressing {
		compressDecompressError = compress(countingInput, countingOutput)
	} else {
		compressDecompressError = decompress(countingInput, countingOutput)
	}
	if compressDecompressError != nil {
		log.Fatal(compressDecompressError)
	}

	var durationInSeconds = time.Since(startTime).Seconds()

	// print summary
	var inputSize = countingInput.Count()
	var outputSize = countingOutput.Count()

	var percentage = float64(outputSize) / float64(inputSize) * 100.0
	var speedInMebibytePerSeconds = float64(inputSize) / 1048576.0 / durationInSeconds

	fmt.Printf("%d -> %d (%.2f%%) in %.2f seconds (%.2f MiB/s)",
		inputSize, outputSize, percentage, durationInSeconds, speedInMebibytePerSeconds)
}
