# opencensus trace propogation chaining

This Go package allows one to try multiple trace propogation formats on incoming HTTP requests
and to set multiple output propogations, if needed.

## Why?

I wrote a test app that I ran in various providers that provided differing trace propogation headers.  So, I decided to try multiple ones.  In hindsight, a command line flag may have been more appropriate...

## Usage

`go get -u github.com/bakins/opencensus-chain-propogation`

And to use with OpenCensus:

```go
import (
    "go.opencensus.io/plugin/ochttp"
    chain "github.com/bakins/opencensus-chain-propogation"
    stackdriver "contrib.go.opencensus.io/exporter/stackdriver/propagation"
    "go.opencensus.io/trace/propagation"
    "go.opencensus.io/plugin/ochttp/propagation/b3"
)

func main() {

    c := Chain{
        // try stackdriver, then b3. Chain stops after first one returns ok.
        Incoming: []propogation.HTTPFormat{
            &stackdriver.HTTPFormat{},
            &b3.HTTPFormat{},
        },
        // only propogate b3 on outgoing
        Outgoing: []propogation.HTTPFormat{
            &b3.HTTPFormat{},
        },
    }

    h := ochttp.Handler{
        Propagation: &c,
    }
}

```

## LICENSE

See [LICENSE](./LICENSE)
