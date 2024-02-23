# blog-go

[![CI](https://github.com/pplmx/blog-go/workflows/CI/badge.svg)](https://github.com/pplmx/blog-go/actions)
[![Coverage Status](https://coveralls.io/repos/github/pplmx/blog-go/badge.svg?branch=main)](https://coveralls.io/github/pplmx/blog-go?branch=main)

blog-go, written in go.

Technologies used:

- [Go Fiber](https://gofiber.io/) for web framework
- [GORM](https://gorm.io/) for ORM
- [lib/pq](https://github.com/lib/pq) for PostgreSQL driver in pure Go
- [glebarez/go-sqlite](https://github.com/glebarez/go-sqlite) for SQLite driver in pure Go
- [Cobra](https://github.com/spf13/cobra) for CLI
- [Viper](https://github.com/spf13/viper) for configuration

## Usage

```bash
# initialize the project
make init

# if you add new repository, or service, or anything that needs to be wired
# if you just want to run the server, you don't need to
make wire

# run the server
make run

# build the binary
make build

# run with the binary
./bin/blog-go
```

## Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you, as defined in the Apache-2.0 license, shall be
dual licensed as above, without any additional terms or conditions.

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

Licensed under either of

* Apache License, Version 2.0
  ([LICENSE-APACHE](LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)
* MIT license
  ([LICENSE-MIT](LICENSE-MIT) or http://opensource.org/licenses/MIT)

at your option.
