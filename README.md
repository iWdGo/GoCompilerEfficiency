# Benchmarking good practices

To avoid bad habits, it is useful to benchmark them against the compiler.

***Logic***
- The most efficient loop is using a [locally defined variable](./src/forloops/README.md)
- Don't bother to use anything but the len() to get the [size of an array](./src/lenarray/README.md)
- Don't duplicate code to handle the [first item of a list](./src/firstitem/README.md)
- Don't try to [keep a single exit](./src/singleexit/README.md)
- [Array initialization](./src/arrayinit/README.md) using append is costly.

***Templates***
- Keep a global [template to parse forms](src/formparse/README.md)
- Keep a global var to hold a [template read from file](src/tmplfile/README.md)

***Buffers***
- No difference when [writing to file](src/tofile/README.md) 
 
 