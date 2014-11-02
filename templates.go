package container

import (
	"github.com/clipperhouse/typewriter"
)

var templates = typewriter.TemplateSet{
	"List": list,
	"Ring": ring,
	"Set":  set,
}
