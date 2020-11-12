# IBM Watson Health Cognitive Services Go SDK Version 0.2.1

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
* An IAM API key to allow the SDK to access your service instance. For instructions about copying the credentials of an existing service instance, or creating a new service instance, refer to Getting Started [here](https://cloud.ibm.com/docs/wh-acd?topic=wh-acd-getting-started).
* Go version 1.12 or above.

## Installation
The current version of this SDK: 0.2.1

There are a few different ways to download and install the WHCS Go SDK project for use by your
Go application:

#### `go get` command  
Use this command to download and install the SDK to allow your Go application to
use it:

```
go get -u github.com/IBM/whcs-go-sdk@v0.2.1
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
  version = "0.2.1"

```

then run `dep ensure`.

## Using the SDK
The ability to store client configuration information in an external credentials file is supported.  The credentials file contains environment variables including authentication data.  Create a credentials file called `ibm-credentials.env`.  Refer to the Prerequisites section for instructions of how to obtain the credentials for your service instance, e.g. URL and API Key.  Define the environment variables in the credentials file as follows:
```
ACD_SERVICE_URL=<service_instance_url, e.g. https://us-south.wh-acd.cloud.ibm.com/wh-acd/api>
ACD_SERVICE_APIKEY=<service_instance_apikey>
ACD_SERVICE_VERSION=<YYYY-MM-DD>
```

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
