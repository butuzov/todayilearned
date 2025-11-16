package main

import (
	"errors"
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type VariableType int

const (
	IsString VariableType = iota
	IsInt
	IsIntSlice
)

type Variable struct {
	Type  VariableType
	Value interface{}
}

func (v *Variable) UnmarshalYAML(unmarshal func(interface{}) error) error {
	errCantValidate := errors.New("can't unmarshal value")

	type plainUint uint
	number := new(uint)
	if err := unmarshal((*plainUint)(number)); err == nil {

		v.Type = IsInt
		v.Value = *number
		return nil
	}

	type plainString string
	str := new(string)
	if err := unmarshal((*plainString)(str)); err == nil {

		v.Type = IsString
		v.Value = *str
		return nil
	}

	type plainIntSlice []int
	var intSlice []int
	if err := unmarshal((*plainIntSlice)(&intSlice)); err == nil {

		v.Type = IsIntSlice
		v.Value = intSlice
		return nil
	}

	return errCantValidate
}

func main() {
	yamls := []string{
		`variable: 1`,
		`variable: boo`,
		`variable: [1,2]`,
	}

	for _, body := range yamls {

		var v struct {
			Data Variable `yaml:"variable"`
		}

		if err := yaml.Unmarshal([]byte(body), &v); err != nil {
			log.Printf("error: %w\n", err)
		} else {
			fmt.Printf("yaml -  %s\n", body)
			fmt.Printf("var  -  %#v\n", v)
		}

		fmt.Println()
	}
}
