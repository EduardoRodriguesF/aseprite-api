# Aseprite API

Do not have Aseprite on your machine? Is not able to compile it? Still needs to export a .aseprite file? Just do it via HTTP.

> This is not a production app, more like a fun little thing I hacked together. There is no actual endpoint available.

## Usage

Send a file through web form. See cURL example below:

```sh
curl -F upload=@path/to/file.aseprite localhost:80/export -o file.png
```

## Running locally

Assuming you already have Go's development tools, run the main.go file providing the path to an Aseprite CLI.

```sh
ASEPRITE=path/to/aseprite go run main.go
```
