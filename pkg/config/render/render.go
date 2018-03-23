// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

// +build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/DataDog/datadog-agent/pkg/config"
)

func main() {
	if len(os.Args[1:]) != 2 {
		panic("please use 'go run render_config.go <template_file> <destination_file>'")
	}

	tplFile, _ := filepath.Abs(os.Args[1])
	tplFilename := filepath.Base(tplFile)
	destFile, _ := filepath.Abs(os.Args[2])

	f, err := os.Create(destFile)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New(tplFilename).ParseFiles(tplFile))
	err = t.Execute(f, config.GetGroups())
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully wrote", destFile)
}
