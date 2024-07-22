# Number
The number layer returns the first character of its passed input, if it contains a number.  Otherwise it returns an error.

## Examples
Here's a valid an invalid example.
### Valid
Here's a valid example.
#### Input
```
324 test
```

#### Output
```
3
```

### Invalid
Here's a valid example.

#### Input
```
invalid
```

#### Output
```
Error code: 1
Error message: The first index of the provided data does not contain a number
```