### `jsonpb` Library - dfuse fork

This is simply a copy of Google Protocol Buffer `jsonpb` package
where `[]byte` types by default outputs to hexadecimal encoding instead
of base64 encoding.

#### Usage

```
package main

func main() {
   block := &pbdeos.Block{}
   blockJSON := jsonpb.MarshalToString(block)
}
```
