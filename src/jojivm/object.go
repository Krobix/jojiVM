package jojivm

//All types and functions related to objects go in this file

import (
	"errors"
	"strconv"
)

type JojiType struct {
	SupportsInt   bool
	SupportsStr   bool
	SupportsFloat bool
	IsDictType    bool
	Fields        map[string]JojiType
}

type JojiObject interface {
	GetInt() (int64, error)
	GetStr() (string, error)
	GetFloat() (float64, error)
	GetField(name string) (JojiObject, error)
	GetType() (JojiType, error)
	SetInt(value int64) error
	SetStr(value string) error
	SetFloat(value float64) error
	SetField(name string, value JojiObject) error
}

//Builtin JojiString type begins here

type JojiString struct {
	value string
}

func (obj JojiString) GetType() (JojiType, error) {
	return JojiType{
		SupportsInt:   true,
		SupportsStr:   true,
		SupportsFloat: true,
		IsDictType:    false,
		Fields:        nil,
	}, nil
}

func (obj *JojiString) GetStr() (string, error) {
	return obj.value, nil
}

func (obj *JojiString) GetInt() (int64, error) {
	i, err := strconv.ParseInt(obj.value, 10, 64)
	errorCheck(err)
	return i, nil
}

func (obj *JojiString) GetFloat() (float64, error) {
	f, err := strconv.ParseFloat(obj.value, 64)
	errorCheck(err)
	return f, nil
}

func (obj *JojiString) SetStr(value string) error {
	obj.value = value
	return nil
}

func (obj *JojiString) SetInt(value int64) error {
	obj.value = strconv.FormatInt(value, 10)
	return nil
}

func (obj *JojiString) SetFloat(value float64) error {
	obj.value = strconv.FormatFloat(value, 'f', -1, 64)
	return nil
}

func (obj *JojiString) GetField(name string) (JojiObject, error) {
	return nil, errors.New("jojiVM: Error: strings do not support fields")
}

func (obj *JojiString) SetField(name string, value JojiObject) error {
	return errors.New("jojiVM: Error: strings do not support fields")
}

//JojiString type ends here
