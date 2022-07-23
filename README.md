# FSL

This is an example of fictitious scripting language (FSL).  FSL is written in JSON.

This system receives FSL as input. The FSL defines variables and functions. Functionality is limited to create, delete, update, add, subtract, multiply, divide, print, as well as function calls. Variables are always numeric.

## Limitations

There are 8 predefined operations:
- create
```json
    {
        "cmd": "create",
        "id": "var1",
        "value": 5
    }
```
- delete
```json
    {
        "cmd": "delete",
        "id": "var1"
    }
```
- update
```json
    {
        "cmd": "update",
        "id": "var1",
        "value": 5
    }
```
- add
```json
    {
        "cmd": "add",
        "id": "var1",
        "operand1": 1,
        "operand2": 2
    }
```
- subtract
```json
    {
        "cmd": "subtract",
        "id": "var1",
        "operand1": 1,
        "operand2": 2
    }
```
- multiply
```json
    {
        "cmd": "multiply",
        "id": "var1",
        "operand1": 1,
        "operand2": 2
    }
```
- divide
```json
    {
        "cmd": "divide",
        "id": "var1",
        "operand1": 1,
        "operand2": 2
    }
```
- print
```json
    {
        "cmd": "print",
        "value": 5
    }
```

Custom function cannot have the same name, overriding is not allowed.

## Usage


Method **`json.NewParser(io.Writer)`** receives writer (needed for print method) and return **`Parser`** interface.

**`Parser`** receives *raw JSON* and call `init` function. To process several `FSL JSONs` need to call *Parse* method for each of them one by one.

### Example

Some of examples of FSL are stored in `sample` directory

### Run script

```shell
go run main.go -f sample/sample1.json -f sample/sample2.json
```
