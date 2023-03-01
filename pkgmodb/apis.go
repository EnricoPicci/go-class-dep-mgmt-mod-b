package pkgmodb

import (
	"fmt"

	"github.com/EnricoPicci/go-class-dep-mgmt-mod-c/pkgmodc"
)

func ApiThatCallsModuleCApiThatCallsModuleDApi() {
	fmt.Println("I am an API of Module B that calls an API of Module C that calls an API of Module D")
	pkgmodc.ApiThatCallsModuleDApi()
}

func ApiThatCallsModuleCApiThatCallRemainsInteral() {
	fmt.Println("I am an API of Module B that calls an API of Module C that remains internal to Module C")
	pkgmodc.ApiThatCallsInternalApi()
}

func ApiThatDoesNotCallExternalApis() {
	fmt.Println("I am an API of Module B that does not call any other API")
}
