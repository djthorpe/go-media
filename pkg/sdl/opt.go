package sdl

import (

	// Namespace imports
	. "github.com/djthorpe/go-errors"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type opts struct {
	surface bool
	title   string
	width   int
	height  int
	logfn   LogFn
	verbose bool
}

type Opt func(*opts) error
type LogFn func(args ...any)

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func OptTitle(v string) Opt {
	return func(o *opts) error {
		if v == "" {
			return ErrBadParameter.Withf("empty title")
		}
		o.title = v
		return nil
	}
}

func OptWindowSize(w, h int) Opt {
	return func(o *opts) error {
		if w < 0 || h < 0 {
			return ErrBadParameter.Withf("invalid window size %dx%d", w, h)
		}
		o.width = w
		o.height = h
		return nil
	}
}

func OptOffscreenRenderer() Opt {
	return func(o *opts) error {
		o.surface = true
		return nil
	}
}

func OptLog(verbose bool, v LogFn) Opt {
	return func(o *opts) error {
		o.verbose = verbose
		o.logfn = v
		return nil
	}
}
