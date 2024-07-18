# WordCount (WC) clone in Go

Wanted to try out go, so I made this clone.

It works the same as the original one with the exception that a file must be passed using the `-f` flag.

---

## How to use:

```sh
usage: wcgo [-h|--help] [-c|--bytes] [-l|--lines] [-w|--words]
            [-m|--characters] [-f|--file <file>]

            Count stuff in a given file

Arguments:

  -h  --help        Print help information
  -c  --bytes       Prints the number of bytes in the file. Default: false
  -l  --lines       Prints the number of lines in the file. Default: false
  -w  --words       Prints the number of words in the file. Default: false
  -m  --characters  Prints the number of characters in the file. Default: false
  -f  --file        File to process
```

