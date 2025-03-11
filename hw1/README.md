# HW1

B10705013 資管四 陳彥廷

## Quick Start

I use go 1.24.0 to run the code.

```bash
go run main.go

# or

./run.sh
```

## Architecture

Below is my file structure.

```txt
├─ go.mod
├─ infra
│  └─ db.go
├─ interface
│  └─ cli.go
├─ main.go
├─ model
│  ├─ listing.go
│  └─ user.go
├─ README.md
├─ run.sh
└─ service
   ├─ listing.go
   └─ user.go
```

- `main.go` is the entry point of the program.
- `infra/` contains the mock database and data layer.
- `model/` defines the data structure and methods.
- `service/` contains the business logic.
- `interface/` contains the command line interface.

## Results

![alt text](/docs/image.png)
