// this file was generated by gomacro command: import _b "flag"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	"flag"
	. "reflect"
)

// reflection: allow interpreted code to import "flag"
func init() {
	Packages["flag"] = Package{
		Binds: map[string]Value{
			"Arg":             ValueOf(flag.Arg),
			"Args":            ValueOf(flag.Args),
			"Bool":            ValueOf(flag.Bool),
			"BoolVar":         ValueOf(flag.BoolVar),
			"CommandLine":     ValueOf(&flag.CommandLine).Elem(),
			"ContinueOnError": ValueOf(flag.ContinueOnError),
			"Duration":        ValueOf(flag.Duration),
			"DurationVar":     ValueOf(flag.DurationVar),
			"ErrHelp":         ValueOf(&flag.ErrHelp).Elem(),
			"ExitOnError":     ValueOf(flag.ExitOnError),
			"Float64":         ValueOf(flag.Float64),
			"Float64Var":      ValueOf(flag.Float64Var),
			"Int":             ValueOf(flag.Int),
			"Int64":           ValueOf(flag.Int64),
			"Int64Var":        ValueOf(flag.Int64Var),
			"IntVar":          ValueOf(flag.IntVar),
			"Lookup":          ValueOf(flag.Lookup),
			"NArg":            ValueOf(flag.NArg),
			"NFlag":           ValueOf(flag.NFlag),
			"NewFlagSet":      ValueOf(flag.NewFlagSet),
			"PanicOnError":    ValueOf(flag.PanicOnError),
			"Parse":           ValueOf(flag.Parse),
			"Parsed":          ValueOf(flag.Parsed),
			"PrintDefaults":   ValueOf(flag.PrintDefaults),
			"Set":             ValueOf(flag.Set),
			"String":          ValueOf(flag.String),
			"StringVar":       ValueOf(flag.StringVar),
			"Uint":            ValueOf(flag.Uint),
			"Uint64":          ValueOf(flag.Uint64),
			"Uint64Var":       ValueOf(flag.Uint64Var),
			"UintVar":         ValueOf(flag.UintVar),
			"UnquoteUsage":    ValueOf(flag.UnquoteUsage),
			"Usage":           ValueOf(&flag.Usage).Elem(),
			"Var":             ValueOf(flag.Var),
			"Visit":           ValueOf(flag.Visit),
			"VisitAll":        ValueOf(flag.VisitAll),
		}, Types: map[string]Type{
			"ErrorHandling": TypeOf((*flag.ErrorHandling)(nil)).Elem(),
			"Flag":          TypeOf((*flag.Flag)(nil)).Elem(),
			"FlagSet":       TypeOf((*flag.FlagSet)(nil)).Elem(),
			"Getter":        TypeOf((*flag.Getter)(nil)).Elem(),
			"Value":         TypeOf((*flag.Value)(nil)).Elem(),
		}, Proxies: map[string]Type{
			"Getter": TypeOf((*P_flag_Getter)(nil)).Elem(),
			"Value":  TypeOf((*P_flag_Value)(nil)).Elem(),
		},
	}
}

// --------------- proxy for flag.Getter ---------------
type P_flag_Getter struct {
	Object  interface{}
	Get_    func(interface{}) interface{}
	Set_    func(interface{}, string) error
	String_ func(interface{}) string
}

func (P *P_flag_Getter) Get() interface{} {
	return P.Get_(P.Object)
}
func (P *P_flag_Getter) Set(unnamed0 string) error {
	return P.Set_(P.Object, unnamed0)
}
func (P *P_flag_Getter) String() string {
	return P.String_(P.Object)
}

// --------------- proxy for flag.Value ---------------
type P_flag_Value struct {
	Object  interface{}
	Set_    func(interface{}, string) error
	String_ func(interface{}) string
}

func (P *P_flag_Value) Set(unnamed0 string) error {
	return P.Set_(P.Object, unnamed0)
}
func (P *P_flag_Value) String() string {
	return P.String_(P.Object)
}
