# Aseprite API

Do not have Aseprite on your machine? Is not able to compile it? Still needs to export a .aseprite file? Just do it via HTTP.

> This is not a production app, more like a fun little thing I hacked together. There is no actual endpoint available.

## Usage

Send a file through web form. See cURL example below:

```sh
curl -F upload=@path/to/file.aseprite localhost:80/export -o file.png
```

## Running locally

1. Aseprite API relies on a local installation of Aseprite. Create a `.env` file and set a variable `ASEPRITE` pointing to the [Aseprite CLI](https://www.aseprite.org/docs/cli/#platform-specific-details) path.
2. Assuming you already have Go's development tools, run the main.go file.

```sh
go run main.go
```

Alternativelly, you can build it and execute the binary.

```sh
go build
./aseprite-api
```
