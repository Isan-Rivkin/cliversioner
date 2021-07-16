# cliversioner

Library to check against remote API current cli version. 

If the current client version is outdated the client will be notified via the cli with a message on how to update the version dependeing on the OS and app context.
## Install 

To install the package 

```bash
go get github.com/isan-rivkin/cliversioner
```

## Example 

```go
package main 

import (
    v "github.com/isan-rivkin/cliversioner"
)

func main(){
    app := "myapp"
    url := "github.com" 
    currentVersion := "0.1.0" 

    input := v.NewInput(app, url, currentVersion,nil) 
    output, err  := v.CheckVersion(input)

    if  err != nil {
        panic(err)
    }

    if output.Outdated {
    	fmt.Printf("%s is not latest, %s, upgrade to %s", output.CurrentVersion, output.Message, output.LatestVersion)
    }
}
```

## Opt out 

The library support opting out via environment variable via `optoutEnvVar` indicator.

Assuming that `optoutEnvVar` = `MY_CLI_CHECK_VERSION` and the caller has the env var `MY_CLI_CHECK_VERSION` = `false` then the lbirary will skip version check against remote server.

```go
func NewInput(app, url, currentVersion string, optoutEnvVar *string) *VersionInput
```




