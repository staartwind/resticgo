# Resticgo
Minimal Go wrapper around the [restic](https://restic.readthedocs.io/) backup cli.

## Installation
```shell
go get github.com/staartwind/resticgo
```
Alternatively the same can be achieved if you use import in a package:
```shell
import "github.com/staartwind/resticgo"
```

## Usage
```shell
import "github.com/staartwind/resticgo"
```

Configuration is done using the builder pattern. Every usual restic config for each command is implemented. Have a look at the code to configure it
```go
resticClient := resticgo.NewRestic(resticgo.WithoutCache)

// Get all the snapshots
snapshots, err := resticClient.Snapshots()
if err != nil {
	panic(err)
}

for _, elem := range snapshots {
	fmt.Println(elem)
}
```

## TODO:
- [ ] Make the json responses typed using structs, maybe not because we cannot be sure which restic version is being used
- [ ] Add unit tests
- [ ] Write docs with all the possible methods etc...