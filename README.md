# ctr-result-parser

Command line tool for parsing and displaying information about CTR (Nintendo 3DS) Result codes.

## Running

This repository provides [releases](https://github.com/nikita-petko/ctr-result-parser/releases).

## Building

Ensure you have [Go 1.20.3+](https://go.dev/dl/)

1. Clone the repository via `git`:

    ```txt
    git clone git@github.com:nikita-petko/ctr-result-parser.git
    cd s3-forcer
    ```

2. Build via [make](https://www.gnu.org/software/make/)

    ```txt
    make build-debug
    ```

## Usage

`cd src && go run main.go --help` (use the build binary found in the bin directory if you downloaded a prebuilt or built it yourself)

```txt
Usage: ctr-error-code-parser
Build Mode: debug
Commit:  
	[-h|--help] [--results] [...errorCodes]

  -help
    	Print usage.
  -results string
    	A comma separated list of error codes to parse, if not specified we will try parse other command line arguments as error codes, if that fails we will parse stdin. (Environment variable: RESULTS)
```

You can input codes any of the following ways:

1. ctr-result-parser -results 0x1234567890,0x1234567890,0x1234567890
2. ctr-result-parser 0x1234567890 0x1234567890 0x1234567890
3. echo 0x1234567890,0x1234567890,0x1234567890 | ctr-result-parser
4. ctr-result-parser -> Will prompt for input

## License

```txt
Copyright 2023 Nikita Petko <petko@vmminfra.net>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
