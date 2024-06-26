// this file was generated by gomacro command: import _b "sort"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	sort "sort"
)

// reflection: allow interpreted code to import "sort"
func init() {
	Packages["sort"] = Package{
		Name: "sort",
		Binds: map[string]Value{
			"Float64s":	ValueOf(sort.Float64s),
			"Float64sAreSorted":	ValueOf(sort.Float64sAreSorted),
			"Ints":	ValueOf(sort.Ints),
			"IntsAreSorted":	ValueOf(sort.IntsAreSorted),
			"IsSorted":	ValueOf(sort.IsSorted),
			"Reverse":	ValueOf(sort.Reverse),
			"Search":	ValueOf(sort.Search),
			"SearchFloat64s":	ValueOf(sort.SearchFloat64s),
			"SearchInts":	ValueOf(sort.SearchInts),
			"SearchStrings":	ValueOf(sort.SearchStrings),
			"Slice":	ValueOf(sort.Slice),
			"SliceIsSorted":	ValueOf(sort.SliceIsSorted),
			"SliceStable":	ValueOf(sort.SliceStable),
			"Sort":	ValueOf(sort.Sort),
			"Stable":	ValueOf(sort.Stable),
			"Strings":	ValueOf(sort.Strings),
			"StringsAreSorted":	ValueOf(sort.StringsAreSorted),
		}, Types: map[string]Type{
			"Float64Slice":	TypeOf((*sort.Float64Slice)(nil)).Elem(),
			"IntSlice":	TypeOf((*sort.IntSlice)(nil)).Elem(),
			"Interface":	TypeOf((*sort.Interface)(nil)).Elem(),
			"StringSlice":	TypeOf((*sort.StringSlice)(nil)).Elem(),
		}, Proxies: map[string]Type{
			"Interface":	TypeOf((*P_sort_Interface)(nil)).Elem(),
		}, 
	}
}

// --------------- proxy for sort.Interface ---------------
type P_sort_Interface struct {
	Object	interface{}
	Len_	func(interface{}) int
	Less_	func(_proxy_obj_ interface{}, i int, j int) bool
	Swap_	func(_proxy_obj_ interface{}, i int, j int) 
}
func (P *P_sort_Interface) Len() int {
	return P.Len_(P.Object)
}
func (P *P_sort_Interface) Less(i int, j int) bool {
	return P.Less_(P.Object, i, j)
}
func (P *P_sort_Interface) Swap(i int, j int)  {
	P.Swap_(P.Object, i, j)
}
