package williamnharveyfrontend

import "context"

var InteractableBackend Backend

type Backend interface {
	// DoSomethingWithInterfaceStruct does something with the instance
	DoSomethingWithInterfaceStruct(context.Context, *InterfaceStruct) (*InterfaceStruct, error)
}

type InterfaceStruct struct {
	Name string
}

func DoSomethingWithInterfaceStruct(ctx context.Context, interfaceStruct *InterfaceStruct) (*InterfaceStruct, error) {
	return InteractableBackend.DoSomethingWithInterfaceStruct(ctx, interfaceStruct)
}
