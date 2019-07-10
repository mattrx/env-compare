# env-compare

Tool to compare multiple env files. Ignores empty lines and comments from env files. Only checks if the variables exist, not the content. Can also handle json files.

## Install
```
$ go install github.com/mattrx/env-compare
```

## Help
```
$ env-compare --help

Usage of env-compare:
  -cs
    	Check should be case sensitive
  -v	Print more verbose output
```

## Usage
```
$ env-compare file1 file2

PARAM_1 from file1:39 missing in file2
PARAM_2 from file1:40 missing in file2
PARAM_9 from file2:1 missing in file1
```
