# go-git-fame

GOLANG port of [fame](https://github.com/oleander/git-fame-rb)

## Example Output
```
+-----------+-------+---------+------+-------------------+
| AUTHOR    | FILES | COMMITS |  LOC | DISTRIBUTION      |
+-----------+-------+---------+------+-------------------+
| Todd Bush |    17 |      66 | 1274 | 94.44/97.06/98.84 |
| todd-bush |     3 |       2 |   15 | 16.67/2.94/1.16   |
+-----------+-------+---------+------+-------------------+
```


## Building

Execute 'make setup' to fetch everything you'll need to build

## Installing
Execute ```make install```

## Running

Execute ```go-git-fame``` from your git repository.

### Options

```--sort``` will allow to sort by commits (default), loc, or files
