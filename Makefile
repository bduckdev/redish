CC = gcc
CLAGS = -std=c89 -pedantic -Wall

build: clean
	mkdir bin
	${CC} ${CFLAGS} -o ./bin/server ./server/main.c
	${CC} ${CFLAGS} -o ./bin/client ./client/main.c
server: build
	./bin/server
client: build
	./bin/client
clean:
	rm -rf ./bin

.PHONY: build run clean
