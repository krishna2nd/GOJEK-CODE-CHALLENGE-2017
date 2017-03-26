// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package commands base command features, shell, file processor
package commands

import (
	"bufio"
	"fmt"
	"os"
)

// FileCommandProcessor to store file name for processing
type FileCommandProcessor struct {
	fileName string
}

// NewFileCmdProcessor to create cmd file processor
func NewFileCmdProcessor(fileName string) *FileCommandProcessor {
	return &FileCommandProcessor{
		fileName: fileName,
	}
}

// Process method to process command file
func (fcp *FileCommandProcessor) Process() error {
	cmdFile, err := os.Open(fcp.fileName)
	if err != nil {
		return err
	}
	defer cmdFile.Close()

	cmdScanner := bufio.NewScanner(cmdFile)
	cmdMgr := NewManager()
	var cmdString string
	for cmdScanner.Scan() {
		cmdString = cmdScanner.Text()
		out, err := cmdMgr.Run(cmdString)
		processOutput(out, err)
	}

	if err := cmdScanner.Err(); err != nil {
		return err
	}
	return nil
}

// processOutput will output user according to valid output, error
func processOutput(out string, err error) {
	if nil == err {
		fmt.Println(out)
	} else {
		fmt.Println(err)
	}
}
