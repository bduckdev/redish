CC = gcc
CLAGS = -std=c89 -pedantic -Wall

build: main.c
	${CC} ${CFLAGS} -o ./main ./main.c
run: build
	./main
	rm -f ./main
clean:
	rm -f ./main

.PHONY: build run clean
