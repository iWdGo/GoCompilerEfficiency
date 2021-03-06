[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/GoCompilerEfficiency)](https://goreportcard.com/report/github.com/iwdgo/GoCompilerEfficiency)

# Benchmarking good practices

Go is [strong typed and object oriented](./src/forloops/README.md) by definition.
To distinguish bad habits, benchmarking against the compiler performance is often convincing.

***Logic***
- The most efficient loop is using a [locally defined variable](./src/forloops/README.md)
- Don't bother to use anything but the len() to get the [size of an array](./src/lenarray/README.md)
- Don't duplicate code to handle the [first item of a list](./src/firstitem/README.md)
- Don't try to [keep a single exit](./src/singleexit/README.md)
- Always use [shadowing inside a for loop](./src/isvalid/README.md) even for a boolean.
- Don't bother about [irrelevant cases](./src/switch/README.md) in a switch.
- Use [bit masking](./src/bitorif/README.md) instead of a second if.
- Get the [bits representation of a byte](./src/bytetobits/README.md) using bits package is simple and efficient.

***Math***
- math.Log2() is the most efficient to check that 2 is the [sole divisor](./src/switch/README.md). 

***Arrays***
- [Array of struct](./src/arraysstruct/README.md) is as efficient as multiple arrays.
- [Array initialization](./src/arrayinit/README.md) using append is costly.
- Using a method is marginally more efficient to [add an element to an array](./src/arrayofarray/README.md).

***Templates***
- Use [implicit parsing for forms](src/formparse/README.md)
- Keep a global var to hold a [template read from file](src/tmplfile/README.md)

***Buffers***
- No difference when [writing to file](src/tofile/README.md) 
- [Returning a buffer](src/returnbuffer/README.md) as a string is marginally more expensive.
- Package `io` is much more efficient than `fmt` which offers many more possibilities:
    - [write a simple string](src/writestring/README.md).
    - [write a buffer to a buffer](src/writebuffer/README.md).
- When handling every byte of a buffer, [use ReadByte](src/readbyte/README.md).

***Builder***
- No difference for a url builder between [returning the string or the url.URL](src/urlbuilder/README.md).
