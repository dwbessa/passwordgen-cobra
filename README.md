Password generator in Go

Learning CLI framework in Go and using Cobra as a initial one

This program is not a CRUD yet, just a CR for now

To build use:
```go build -o passwordgen```

We can create with 
```./passwordgen generate```
We register the password filling first Network Name (generic name for where this password will be used) and then putting the e-mail

We can read with
```./passwordgen search [network name]```
