// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import namestruct "github.com/jmattheis/goverter/example/name-struct"

type RenamedConverter struct{}

func (c *RenamedConverter) Convert(source namestruct.Input) namestruct.Output {
	var exampleOutput namestruct.Output
	exampleOutput.Name = source.Name
	return exampleOutput
}