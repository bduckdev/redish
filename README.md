# Redish

When mom says "we have Redis at home"

## How to use

1. In one terminal run: `bash
git clone https://github.com/bduckdev/redish \
cd redish \
go run .`

2. In another run: `
nc localhost 6379`
3. Behold the functionality: `SET foo bar
OK
GET foo
bar
GET notakey
(nil)`

## Features

- you can connect, and you can set a key, then you can get it back

## todo

- basically everything else
