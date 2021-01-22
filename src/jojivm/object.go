package jojivm

//All types and functions related to objects go in this file

type JojiType struct {
	supportsInt   bool
	supportsStr   bool
	supportsFloat bool
	isDictType    bool
	fields        map[string]JojiType
}

type JojiObject interface {
	GetInt() (int64, error)
	GetStr() (string, error)
	GetFloat() (float64, error)
	SetInt() error
	SetStr() error
	SetFloat() error
}
