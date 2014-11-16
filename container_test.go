package container

import (
	"bytes"
	"testing"

	"github.com/clipperhouse/typewriter"
)

func TestValidate(t *testing.T) {
	var err error

	g := NewContainerWriter()

	typ2 := typewriter.Type{
		Name: "SomeType2",
		Tags: typewriter.TagSlice{
			typewriter.Tag{
				Name:   "containers",
				Values: []typewriter.TagValue{},
			},
		},
	}

	var b2 bytes.Buffer
	err = g.Write(&b2, typ2)

	if b2.Len() > 0 {
		t.Errorf("empty 'containers' tag should not write")
	}

	if err != nil {
		t.Error(err)
	}

	typ3 := typewriter.Type{
		Name: "SomeType3",
		Tags: typewriter.TagSlice{
			typewriter.Tag{
				Name: "containers",
				Values: []typewriter.TagValue{
					{"List", nil},
				},
			},
		},
	}

	var b3 bytes.Buffer
	err = g.Write(&b3, typ3)

	if b3.Len() == 0 {
		t.Errorf("'containers' tag with List should write (and ignore others)")
	}

	if err != nil {
		t.Error(err)
	}
}
