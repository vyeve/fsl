# FSL

This is example of fictitious scripting language (FSL).  FSL is written in JSON.

This system receives FSL as input. The FSL defines variables and functions. Functionality is limited to create, delete, update, add, subtract, multiply, divide, print, as well as function calls. Variables are always numeric.

## Limitations

There are 8 predefined operations:
- create
- delete
- update
- add
- subtract
- multiply
- divide
- print

Custom function cannot have the same name, overriding is not allowed.

## Usage

Method **`json.NewParser(io.Writer)`** receives writer (needed for print method) and return **`Parser`** interface.

**`Parser`** receives *raw JSON* and call `init` function. To process several `FSL JSONs` need to call *Parse* method for each one by one.

### Example

```shell
go run main.go -f sample/sample1.json -f sample/sample2.json
```
