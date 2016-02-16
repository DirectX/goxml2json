# goxml2json [![Build Status](https://drone.io/github.com/basgys/goxml2json/status.png)](https://drone.io/github.com/basgys/goxml2json/latest)

Go package that converts XML to JSON

### Install

    go get -u github.com/basgys/goxml2json

### Importing

    import github.com/basgys/goxml2json

### Usage

**Code example**

```go
  package main

  import (
  	"fmt"
  	"strings"

  	xj "github.com/basgys/goxml2json"
  )

  func main() {
  	// xml is an io.Reader
  	xml := strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?><hello>world</hello>`)
  	json, err := xj.Convert(xml)
  	if err != nil {
  		panic("That's embarrassing...")
  	}

  	fmt.Println(json.String())
  	// {"hello": "world"}
  }

```

**Input**

```xml
  <?xml version="1.0" encoding="UTF-8"?>
  <osm version="0.6" generator="CGImap 0.0.2">
   <bounds minlat="54.0889580" minlon="12.2487570" maxlat="54.0913900" maxlon="12.2524800"/>
   <foo>bar</foo>
  </osm>
```

**Output**

```json
  {
    "osm": {
      "-version": "0.6",
      "-generator": "CGImap 0.0.2",
      "bounds": {
        "-minlat": "54.0889580",
        "-minlon": "12.2487570",
        "-maxlat": "54.0913900",
        "-maxlon": "12.2524800"
      },
      "foo": "bar"
    }
  }
```

### XML Attriutes mapping

By default all attributes are being decorated on decoding using `"-"` prefix. Thus this example XML
```xml
<?xml version="1.0" encoding="UTF-8"?>
<hello foo="bar">
  <whole>world</whole>
</hello>
```
will be transformed to
```json
{ "hello": { "-foo": "bar", "whole": "world" } }
```
To override this behaviour function
```go
json, err := xml2json.ConvertWithAttrPrefix(xml, "")
```
can be used instead.
```json
{ "hello": { "foo": "bar", "whole": "world" } }
```

### Disclaimer
This project has been hacked in a few hours and is definitely not production ready.

### Contributing
Feel free to contribute to this project if you want to fix/extend/improve it.

### Contributors

  - [DirectX](https://github.com/directx)

### TODO

   * Extract data types in JSON (numbers, boolean, ...)
   * Categorise errors
   * Option to prettify the JSON output
   * Benchmark
