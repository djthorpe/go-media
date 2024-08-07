package sdl

import (
	"encoding/json"

	// Namespace imports

	. "github.com/mutablelogic/go-media"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type Device struct {
	t    Type
	name string
}

////////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	DefaultAudioSource = &Device{AUDIO | INPUT, ""}
	DefaultAudioSink   = &Device{AUDIO | OUTPUT, ""}
)

////////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newDevice(t Type, name string) *Device {
	return &Device{t, name}
}

////////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (d *Device) MarshalJSON() ([]byte, error) {
	type j struct {
		Type Type   `json:"type"`
		Name string `json:"name"`
	}
	return json.Marshal(j{
		Type: d.t,
		Name: d.name,
	})
}

func (d *Device) String() string {
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (d *Device) Type() Type {
	return d.t
}

func (d *Device) Name() string {
	return d.name
}
