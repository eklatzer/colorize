# colorize

`colorize` is a small CLI tool that helps to analyze logs by coloring lines based on regex matching.

## Usage

`colorize` can be used by directly providing the path to a logfile (absolute or relative) or by piping the output of some other command like `cat` or `tail` into it.

1. With file:
```bash
colorize <path to file>
```
2. By piping the output of commands like `cat` or `tail`:
```bash
cat example.log | colorize
```

```bash
tail example.log | colorize    
```
