# printlint

This package warns when `fmt.Println` or friends are used.

## Example

```bash
$ go install github.com/bahlo/printlint/cmd/printlint
$ go vet -vettool=$GOBIN/printlint ./test
# github.com/bahlo/printlint/test
test/print.go:8:2: fmt.Println found "fmt.Println(\"foo bar\")"
```

## References

https://github.com/fatih/addlint
