// redish server
#include "../config.h"
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

static void die(const char *msg);
static void do_a_thing(int connfd);

int main() {
  // get socket handle
  int fd = socket(AF_INET, SOCK_STREAM, 0);

  // set socket options
  int val = 1;
  setsockopt(fd, SOL_SOCKET, SO_REUSEADDR, &val, sizeof(val));

  // bind
  struct sockaddr_in addr = {};
  addr.sin_family = AF_INET;
  addr.sin_port = htons(PORT);
  addr.sin_addr.s_addr = htonl(0);
  int rv = bind(fd, (const struct sockaddr *)&addr, sizeof(addr));
  if (rv) {
    die("bind()");
  }

  // listen
  rv = listen(fd, SOMAXCONN);
  if (rv) {
    die("listen()");
  }

  for (;;) {
    // accept
    struct sockaddr_in client_addr = {};
    socklen_t addrlen = sizeof(client_addr);
    int connfd = accept(fd, (struct sockaddr *)&client_addr, &addrlen);
    if (connfd < 0) {
      continue;
    }

    do_a_thing(connfd);
    close(connfd);
  }

  printf("server exited\n");
  return 0;
}

static void die(const char *msg) {
  perror(msg);
  exit(1);
}

// dummy function for sending/recieving data
static void do_a_thing(int connfd) {
  char rbuf[64] = {};
  ssize_t n = read(connfd, rbuf, sizeof(rbuf) - 1);
  if (n < 0) {
    return;
  }
  printf("message from client: %s\n", rbuf);

  char wbuf[] = "hah, goteem";
  write(connfd, wbuf, strlen(wbuf));
}
