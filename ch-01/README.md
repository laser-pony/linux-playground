# Linux Playground Chapter 1

## Don't panic

- I don't know the tools,
- neither the online resources,
- mentioned in this ticket by heart.


- I just want to encourage the following strategy:
  - know that something is possible,
  - remember the buzzwords,
  - know how to google / search / ask the AI.

---

## In this playground we are going to

  - Learn how to open a socket in Java (classic echo server),
    - also see vim in action,
    - as well as the Java compiler and the Java 'interpreter'.
  - Learn how to write to the socket using netcat from the command line,
  - using tcpdump to observe what's happening over the network,
  - using netstat to observe what is happening to the sockets.

## Creating a Java source file

We use the editor Vim 

official website: https://www.vim.org/ \
learning vim: https://vimschool.netlify.app/introduction/vimtutor/

to create the following source file, named `EchoServer.java`.

There is a highlighted line within that file, that we will either leave as it is, or delete / comment, in order to see the behaviour of the OS regarding the sockets

https://en.wikipedia.org/wiki/Network_socket \
https://en.wikipedia.org/wiki/Transmission_Control_Protocol \
https://stackoverflow.com/questions/5992211/list-of-possible-internal-socket-statuses-from-proc

```java
import java.io.*;
import java.net.*;

public class EchoServer {

    public static void main(String[] args) throws IOException {

        int port = 12345;
        ServerSocket serverSocket = new ServerSocket(port);
        System.out.println("Echo server listening on port " + port);

        while (true) {
            Socket clientSocket = serverSocket.accept();
            System.out.println("Accepted connection from " + clientSocket.getRemoteSocketAddress());

            // Handle one client per thread
            new Thread(() -> {
                try (
                    BufferedReader in = new BufferedReader(new InputStreamReader(clientSocket.getInputStream()));
                    PrintWriter out = new PrintWriter(clientSocket.getOutputStream(), true)
                ) {
                    String inputLine;
                    while ((inputLine = in.readLine()) != null) {
                        out.println("echo: " + inputLine);
                    }
                    // Client closed connection, now we close the socket
                    System.out.println("Client disconnected: " + clientSocket.getRemoteSocketAddress());
                    /////////////////////////////////////////////////
                    clientSocket.close(); // <-----------------------
                    /////////////////////////////////////////////////
                    Thread.sleep(100000);
                } catch (Exception e) {
                    e.printStackTrace();
                }
            }).start();
        } // end while

    } // end main

} // end EchoServer
```

## Starting the program

```bash
java EchoServer.java
# ... quit with ctrl-c
```

### you can also compile and then execute

```bash
javac EchoServer.java
ls -1
EchoServer.class
EchoServer.java
# then, also posssble
java EchoServer # executes the class file
```

The output might look something like this:

```text
$ java EchoServer
Echo server listening on port 12345

... after some time when people connect:

Accepted connection from /127.0.0.1:35600

... after some time when people disconnect:

Client disconnected: /127.0.0.1:35600
```

## Connect to the server via netcat

```bash
# the first "Hello" is typed by us, the second one is received
peterpan@1-23-0058:~$ netcat localhost 12345
Hello
echo: Hello
```

## Observing what we send (and receive) via tcpdump


```bash

tcpdump -A -i any 'tcp port 12345'

... we will type "||asdfasdf||" via netcat

16:16:06.640112 lo    In  IP localhost.36412 > localhost.12345: Flags [P.], seq 1:14, ack 1, win 512, options [nop,nop,TS val 1899332370 ecr 1899327948], length 13
E..A..@.@.s..........<09RJ.#.7.T.....5.....
q5..q5q.||asdfasdf||

...

16:16:06.640174 lo    In  IP localhost.12345 > localhost.36412: Flags [.], ack 14, win 512, options [nop,nop,TS val 1899332370 ecr 1899332370], length 0
E..4$.@.@...........09.<.7.TRJ.0.....(.....
q5..q5..
16:16:06.641300 lo    In  IP localhost.12345 > localhost.36412: Flags [P.], seq 1:20, ack 14, win 512, options [nop,nop,TS val 1899332371 ecr 1899332370], length 19
E..G$.@.@...........09.<.7.TRJ.0.....;.....
q5..q5..echo: ||asdfasdf||
```

## Observing what happens to the sockets:

While we play around with the EchoServer (alternating between using or not using the close() method), we might observe the following via

```bash
netstat -tupen
```

Output when closing the socket properly:

```text
Active Internet connections (w/o servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       User       Inode      PID/Program name
tcp        0      0 127.0.0.1:40170         127.0.0.1:12345         TIME_WAIT   0          0          -
tcp        0      0 127.0.0.1:56720         127.0.0.1:12345         TIME_WAIT   0          0          -
tcp        0      0 127.0.0.1:56728         127.0.0.1:12345         TIME_WAIT   0          0          -
```

Output when forgetting to close the socket:

```text
Active Internet connections (w/o servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       User       Inode      PID/Program name
tcp        0      0 127.0.0.1:57988         127.0.0.1:12345         FIN_WAIT2   0          0          -
tcp        0      0 127.0.0.1:58008         127.0.0.1:12345         FIN_WAIT2   0          0          -
tcp        0      0 127.0.0.1:57994         127.0.0.1:12345         FIN_WAIT2   0          0          -
tcp6       0      0 127.0.0.1:12345         127.0.0.1:57994         CLOSE_WAIT  1002       150149     26599/java
tcp6       0      0 127.0.0.1:12345         127.0.0.1:57988         CLOSE_WAIT  1002       150148     26599/java
tcp6       0      0 127.0.0.1:12345         127.0.0.1:58008         CLOSE_WAIT  1002       150150     26599/java
```