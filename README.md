# golibs
[![GoDoc](https://godoc.org/github.com/rogelioConsejo/golibs?status.svg)](https://godoc.org/github.com/rogelioConsejo/golibs)

`golibs` provides a collection of Go libraries that address common development needs, including but not limited to
 CORS handling, user authentication, and more.

 The project aims to offer easy-to-integrate solutions that can help developers save time and effort in setting up
 foundational features of their applications. 
 
 ## Features
    - Basic CORS handling
    - Basic User Authentication
    - Simple File Persistence with concurrency support

## Getting Started
### Installation
To include golibs in your project, you can fetch the library using go get:
```sh
go get github.com/rogelioConsejo/golibs
```

### Usage
To use the libraries in your project, you can import them in your Go source files:
```go
import "github.com/rogelioConsejo/golibs/cors"
import "github.com/rogelioConsejo/golibs/auth"
import "github.com/rogelioConsejo/golibs/file"
```

## Contributing
If you have any suggestions, feature requests, or bug reports, please feel free to open an issue or submit a pull request.

- Fork the repository.
- Create a new branch for your changes.
- Implement your changes.
- Submit a pull request with a detailed description of your modifications.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
