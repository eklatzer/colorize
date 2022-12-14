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

## Configuration

The regex-matching and the used colors can be configured in two ways:
* Config file at the home directory: `$HOME/.colorized.yaml`
* By defining a config file with the flag `--config`

If no config file is provided default values are taken.
": .*color.*
### Example Configuration

```yaml
colors:
  - expression: (?i).*fatal.*
    color: "#FF0000"
  - expression: (?i).*error.*
    color: "#FFA500"
  - expression: (?i).*warn.*
    color: "#FFFF00"
default:
  color: "#FFFFF"
```
