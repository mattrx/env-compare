# env-compare

Tool to compare multiple env files. Ignores empty lines and comments. Only checks if the variables exist, not the content.

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
$ env-compare path1 path2

PARAM_1 from .env.dist:39 missing in .env
PARAM_2 from .env.dist:40 missing in .env
PARAM_9 from .env:1 missing in .env.dist
```
