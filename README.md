# chouse

Simple [Golang](https://golang.org/) wrapper for the [Companies House Rest API service](https://www.gov.uk/government/organisations/companies-house).

Allows to perform API calls to look up company profile information among other details like company officers, filling history etc.

All API endpoints are supported and data returned is parsed JSON objects.

This project has no tests however it is being using in production at [coreportsly.com](http://coreportsly.com)

### Usage

* Sign up for the service & register your app at https://developer.companieshouse.gov.uk/api/docs/
* Put your API key in environment variable

```shell
export CHOUSE_APIKEY='<your_secret_api_key>'
```

* Example code using this package

```golang
package main

import (
   "github.com/coreportsly/chouse"
)

func main() {
    ch := chouse.Explore('companyNumber')
    ch.Company() // returns Company profile
}
```

* ...
* Profit!
