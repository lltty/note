package lib

type Set interface {
	Add(interface{})
	Remove(interface{})
	Clear()
	Contains(interface{})
	Len() int
	Same(other Set) bool
	Elements() []interface{}
	String() string

	//集合方法
	Union(other Set) []interface{}
	Intersect(other Set) []interface{}
	Diff(other Set) []interface{}
	SymmetricDiff(other Set) []interface{}
}
