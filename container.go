package container

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewContainerWriter())
	if err != nil {
		panic(err)
	}
}

type ContainerWriter struct{}

func NewContainerWriter() *ContainerWriter {
	return &ContainerWriter{}
}

func (c *ContainerWriter) Name() string {
	return "container"
}

func (c *ContainerWriter) Imports(t typewriter.Type) (result []typewriter.ImportSpec) {
	// none
	return result
}

func (c *ContainerWriter) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(c)

	if !found {
		return nil
	}

	writeLicenses(w, tag)

	for _, v := range tag.Values {
		tmpl, err := templates.ByTagValue(t, v)
		if err != nil {
			return err
		}

		if err := tmpl.Execute(w, t); err != nil {
			return err
		}
	}

	return nil
}

func writeLicenses(w io.Writer, tag typewriter.Tag) error {
	var list, ring, set bool

	for _, v := range tag.Values {
		if v.Name == "List" {
			list = true
		}
		if v.Name == "Ring" {
			ring = true
		}
		if v.Name == "Set" {
			set = true
		}
	}

	if list {
		license := `// List is a modification of http://golang.org/pkg/container/list/
`
		if _, err := w.Write([]byte(license)); err != nil {
			return err
		}
	}

	if ring {
		license := `// Ring is a modification of http://golang.org/pkg/container/ring/
`
		if _, err := w.Write([]byte(license)); err != nil {
			return err
		}
	}

	if list || ring {
		license := `// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE

`
		if _, err := w.Write([]byte(license)); err != nil {
			return err
		}
	}

	if set {
		license := `// Set is a modification of https://github.com/deckarep/golang-set
// The MIT License (MIT)
// Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)
`

		if _, err := w.Write([]byte(license)); err != nil {
			return err
		}
	}

	return nil
}
