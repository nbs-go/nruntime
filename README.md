# nruntime

[![Go Report Card](https://goreportcard.com/badge/github.com/nbs-go/nruntime)](https://goreportcard.com/report/github.com/nbs-go/nruntime)
[![GitHub license](https://img.shields.io/github/license/nbs-go/nruntime)](https://github.com/nbs-go/nruntime/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/nbs-go/nruntime/branch/master/graph/badge.svg?token=IUK6YSGRNT)](https://codecov.io/gh/nbs-go/nruntime)

Define Application runtime with Manifest and intercept common program arguments such as `help` and `version`. 

Application Manifest value is stored globally. You can retrieve values from anywhere with getter functions. An Application Manifest is defined by:
1. Application Name
2. Application Version
3. Build Signature (e.g. Git Commit Hash)

## Installing

```shell
go get github.com/nbs-go/nruntime
```

## Example

### Application Manifest

```go
package main

import (
	"fmt"
	"github.com/nbs-go/nruntime"
)

const (
	// Define Application Manifest
	appName        = "MyApp"
	appVersion     = "v1.0.0"
	buildSignature = "9f816a75ca6fcd264c791a70a86d8a4a498cf22d"
)

func init() {
	nruntime.NewRuntime(appName, appVersion, buildSignature)
}

func main() {
	// Get application manifest values
	fmt.Printf("AppName=%s AppVersion=%s BuildSignature=%s\n",
		nruntime.AppName(),nruntime.AppVersion(), nruntime.BuildSignature())
	
	// Get application uptime with started at
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Printf("Uptime=%s\n", time.Since(nruntime.StartedAt()))
}
```

### Intercept Common Program Arguments

> TODO

#### Usage

> TODO

## Contributors

- Saggaf Arsyad <saggaf.arsyad@gmail.com>