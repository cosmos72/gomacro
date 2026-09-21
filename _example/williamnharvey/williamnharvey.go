package williamnharvey

import (
	"context"
	"fmt"

	"github.com/cosmos72/gomacro/_example/williamnharvey/williamnharveybackend"
	"github.com/cosmos72/gomacro/_example/williamnharvey/williamnharveyfrontend"
)

func main() {
	williamnharveyfrontend.InteractableBackend = williamnharveybackend.InteractableBackend{}

	newStruct := williamnharveyfrontend.InterfaceStruct{}
	williamnharveyfrontend.DoSomethingWithInterfaceStruct(context.Background(), &newStruct)

	fmt.Println(newStruct.Name)
}
