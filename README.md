# IBM Watson Health Cognitive Services Go SDK Version 0.0.2

## Overview

The IBM Watson Health Cognitive Services (WHCS) Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[Annotator for Clinical Data](https://cloud.ibm.com/apidocs/wh-acd) | annotatorforclinicaldataacdv1 |
[Insights for Medical Literature](https://cloud.ibm.com/apidocs/wh-iml) | insightsformedicalliteratureservicev1 |

NOTE: You must be signed into IBM Cloud to see the docs.

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.12 or above.

## Installation
The current version of this SDK: 0.0.2

There are a few different ways to download and install the MySDK Go SDK project for use by your
Go application:

#### `go get` command  
Use this command to download and install the SDK to allow your Go application to
use it:

```
go get -u github.com/IBM/whcs-go-sdk
```

#### Go modules  
If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
	"github.com/IBM/whcs-go-sdk/annotatorforclinicaldataacdv1"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

#### `dep` dependency manager  
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```
[[constraint]]
  name = "github.com/IBM/whcs-go-sdk"
  version = "0.0.2"

```

then run `dep ensure`.

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md).

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/whcs-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/).

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

The IBM Watson Health Cognitive Services Go SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
