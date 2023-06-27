# Go Assoc Query
Searches for and retrieves a file or protocol association-related string from the registry.
## Install
```
go get github.com/goevil/go-assoc-query
```
## Example
```
package main

import (
	"fmt"
	"syscall"

	"github.com/goevil/go-assoc-query"
)

const (
	MAX_PATH            = 260
	ASSOCSTR_EXECUTABLE = 2
)

func main() {
	var path [MAX_PATH]uint16
	var size uint32 = MAX_PATH
	extPtr, _ := syscall.UTF16PtrFromString(".html")
	_ = assoc.AssocQueryString(0, ASSOCSTR_EXECUTABLE, extPtr, nil, &path[0], &size)
	fmt.Println(syscall.UTF16ToString(path[:size]))
}
// If we assume that Brave is the default browser, the possible output would be as follows.
// C:\Program Files\BraveSoftware\Brave-Browser\Application\brave.exe
```

## Reference
[AssocQueryStringW function (shlwapi.h)](https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-assocquerystringw)

## License
[MIT](https://choosealicense.com/licenses/mit/)