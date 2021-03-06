package jojivm

//All types and functions related to objects go in this file

import (
	"errors"
	"fmt"
	"math"
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
	GetField(name string) (*JojiObject, error)
	GetType() (*JojiType, error)
	SetInt(value int64) error
	SetStr(value string) error
	SetFloat(value float64) error
	SetField(name string, value *JojiObject) error
}

//Static primitive type definitions

var JojiStringType JojiType = JojiType{
	SupportsInt:   true,
	SupportsStr:   true,
	SupportsFloat: true,
	IsDictType:    false,
	Fields:        nil,
}

var JojiIntType JojiType = JojiType{
	SupportsInt:   true,
	SupportsStr:   true,
	SupportsFloat: true,
	IsDictType:    false,
	Fields:        nil,
}

var JojiFloatType JojiType = JojiType{
	SupportsInt:   true,
	SupportsStr:   true,
	SupportsFloat: true,
	IsDictType:    false,
	Fields:        nil,
}

//Builtin JojiString type begins here

type JojiString struct {
	value string
}

func (obj JojiString) GetType() (*JojiType, error) {
	return &JojiStringType, nil
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

func (obj *JojiString) GetField(name string) (*JojiObject, error) {
	return nil, errors.New("jojiVM: Error: strings do not support fields")
}

func (obj *JojiString) SetField(name string, value *JojiObject) error {
	return errors.New("jojiVM: Error: strings do not support fields")
}

//JojiString type ends here
//JojiInt type begins here

type JojiInt struct {
	value int64
}

func (obj *JojiInt) GetType() (*JojiType, error) {
	return &JojiIntType, nil
}

func (obj *JojiInt) GetInt() (int64, error) {
	return obj.value, nil
}

func (obj *JojiInt) GetStr() (string, error) {
	str := strconv.FormatInt(obj.value, 10)
	return str, nil
}

func (obj *JojiInt) GetFloat() (float64, error) {
	return float64(obj.value), nil
}

func (obj *JojiInt) SetInt(value int64) error {
	obj.value = value
	return nil
}

func (obj *JojiInt) SetStr(value string) error {
	i, err := strconv.ParseInt(value, 10, 64)
	errorCheck(err)
	obj.value = i
	return nil
}

func (obj *JojiInt) SetFloat(value float64) error {
	f := math.Floor(value)
	obj.value = int64(f)
	return nil
}

func (obj *JojiInt) GetField(name string) (*JojiObject, error) {
	return nil, errors.New("jojiVM: Error: ints do not support fields")
}

func (obj *JojiInt) SetField(name string, value *JojiObject) error {
	return errors.New("jojiVM: Error: ints do not support fields")
}

//End of JojiInt type - beginning of JojiFloat type

type JojiFloat struct {
	value float64
}

func (obj *JojiFloat) GetType() (*JojiType, error) {
	return &JojiFloatType, nil
}

func (obj *JojiFloat) GetFloat() (float64, error) {
	return obj.value, nil
}

func (obj *JojiFloat) GetInt() (int64, error) {
	f := math.Floor(obj.value)
	return int64(f), nil
}

func (obj *JojiFloat) GetStr() (string, error) {
	return fmt.Sprintf("%f", obj.value), nil
}

func (obj *JojiFloat) SetFloat(value float64) error {
	obj.value = value
	return nil
}

func (obj *JojiFloat) SetInt(value int64) error {
	obj.value = float64(value)
	return nil
}

func (obj *JojiFloat) SetStr(value string) error {
	f, err := strconv.ParseFloat(value, 64)
	errorCheck(err)
	obj.value = f
	return nil
}

func (obj *JojiFloat) GetField(name string) (*JojiObject, error) {
	return nil, errors.New("jojiVM: Error: float objects do not support fields")
}

func (obj *JojiFloat) SetField(name string, value *JojiObject) error {
	return errors.New("jojiVM: Error: float objects do not support fields")
}

//end JojiFloat type
