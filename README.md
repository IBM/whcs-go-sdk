# IBM Watson Health Cognitive Services Go SDK Version 0.3.1

## Overview

<b>Deprecated:</b>  Consider migrating to [Java](https://github.com/IBM/whcs-java-sdk) or [Python](https://github.com/IBM/whcs-python-sdk).

The IBM Watson Health Cognitive Services (WHCS) Go SDK allows developers to programmatically interact with the following IBM Cloud service:

Service Name | Package name
--- | ---
[Annotator for Clinical Data](https://cloud.ibm.com/apidocs/wh-acd) | annotatorforclinicaldataacdv1 |

NOTE: You must be signed into IBM Cloud to see the docs.

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your service instance. For instructions about copying the credentials of an existing service instance, or creating a new service instance, refer to Getting Started [here](https://cloud.ibm.com/docs/wh-acd?topic=wh-acd-getting-started).
* Go version 1.14 or above.

## Installation
The current version of this SDK: 0.3.1

#### Go modules
If your application uses Go modules for dependency management (recommended), just add an import to your
Go application.
Here is an example:

```go
import (
        "github.com/IBM/whcs-go-sdk/annotatorforclinicaldataacdv1"
)
```

Next, run `go build` or `go mod tidy` to download and install the new dependency and update your application's
`go.mod` file.

#### `go get` command  
Alternatively, you can use the `go get` command to download and install the package to allow your application to
use it:

```
go get -u github.com/IBM/whcs-go-sdk@v0.3.1
```
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
