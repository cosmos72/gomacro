// this file was generated by gomacro command: import _b "sync"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"sync"
)

// reflection: allow interpreted code to import "sync"
func init() {
	Packages["sync"] = Package{
		Binds: map[string]Value{
			"NewCond": ValueOf(sync.NewCond),
		}, Types: map[string]Type{
			"Cond":      TypeOf((*sync.Cond)(nil)).Elem(),
			"Locker":    TypeOf((*sync.Locker)(nil)).Elem(),
			"Map":       TypeOf((*sync.Map)(nil)).Elem(),
			"Mutex":     TypeOf((*sync.Mutex)(nil)).Elem(),
			"Once":      TypeOf((*sync.Once)(nil)).Elem(),
			"Pool":      TypeOf((*sync.Pool)(nil)).Elem(),
			"RWMutex":   TypeOf((*sync.RWMutex)(nil)).Elem(),
			"WaitGroup": TypeOf((*sync.WaitGroup)(nil)).Elem(),
		}, Proxies: map[string]Type{
			"Locker": TypeOf((*P_sync_Locker)(nil)).Elem(),
		},
	}
}

// --------------- proxy for sync.Locker ---------------
type P_sync_Locker struct {
	Object  interface{}
	Lock_   func(interface{})
	Unlock_ func(interface{})
}

func (P *P_sync_Locker) Lock() {
	P.Lock_(P.Object)
}
func (P *P_sync_Locker) Unlock() {
	P.Unlock_(P.Object)
}
