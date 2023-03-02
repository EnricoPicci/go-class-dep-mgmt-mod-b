package pkgmodb1

import (
	"fmt"
)

func ApiThatDoesNotCallExternalApis() {
	fmt.Println("I am an API of Module B in pkgmodb1 that does not call any other API ")
}
