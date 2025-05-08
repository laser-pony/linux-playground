# Linux Playground Chapter 3

## Goals

In this chapter we are going to cover the absolute basics on how to interact with the linux operating system.

### Birds view:

- An operating system (OS) is sort of the 'main program' running on a computer on any time.
    - Windows,
    - Linux,
    - iOS,
    - Android... are known to most people.
- The definition of the (sub-) program that users utilize to interact with an operating system is called a **shell**

A **shell** could be a graphical user interface, such as the Windows-GUI ( Start menu, taskbar, file explorer, and window management) you see on a Windows laptop,

or the graphical interface (Aqua) on your MacBook,

or the command line interface, when you log into a Linux server.
- neither the online resources,
- mentioned in this ticket by heart.

or SwiftUI/UIKit on an I-Phone, or the Android UI.

---

The term **shell**, today, is more often assiciated with a **command line
interface (CLI)***, sometimes also called a **promt**.

---

## How a **shell** works

### First example - the 'ls' command

```bash
$ ls -l /opt
```

You get an answer

```text
total 0
drwxr-xr-x 1 root root   16 Dec 17  2023 brother
drwx--x--x 1 root root   12 Jun 16  2024 containerd
drwxr-xr-x 1 root root   12 May  2 05:51 google
drwxr-xr-x 1 root root   22 Jul 28  2024 vagrant
drwxr-xr-x 1 root root 1064 Apr 21 11:02 vivaldi
```

**Explanation**

In the above command,

- 'ls' is the actual command - **list** a specific directory.
- '-l' is an option (also called flag) of that command, that could be used, but is not mandatory. In this case it stands for
  **long** format. That means, ls should display not only the content of
  the current working directory, but also show the owner and the group
  membership of each entyr - and show last modification times, as
  well as file permissions for the entry.
- '/opt' is a directory that hosts optional, most commonly manual / non-distribution software installations on a linux system.

So, the general syntax is:

```
<command-name> <some-option/flag> <some-other-option/flag> <some-parameter>
```

### Second example - the 'echo' command

On some system the prompt might look different - usually a dollar-sign
marks the begin of a promt.

The echo command just outputs parameters as they are received.

The echo command understands an (not mandatory) parameter '-n' - if that
parameter is used, echo does **not** add a newline character \
(-n means **n**o newline)

```bash
$ echo "Hello"
Hello
$ echo -n "Hello"
Hello$ <more commands follow here> 
```

### Third example - the 'man' command

The 'man' command is different from echo and ls in the sense that it is
a modal interface. That means, after sending the command to the OS,
you don't get an immediate response and re-enter the shell, instead
you are 'stuck' in a pager, that lets you navigate up and down a help
page.

You can navigate that help page with the \<up\> and \<down\> keys, and
you can leave the 'modal interface' by pressing the 'q' key on your
keyboard:

Enter:
```bash
man ls
```

You will see:

```
LS(1)                        User Commands                       LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List  information about the FILEs (the current directory by de‐
       fault).  Sort entries alphabetically if none of  -cftuvSUX  nor
       --sort is specified.
...
```

AI says:

```
What sections are typically part of a man page?

ChatGPT said:
Typical sections of a Unix man page include:

NAME – The command or function name with a short description.
SYNOPSIS – Usage syntax, including command options and arguments.
DESCRIPTION – Detailed explanation of what the command or function does.
OPTIONS – List and explanation of available flags/options.
EXIT STATUS – Describes possible exit codes and their meanings.
RETURN VALUE – For library calls, the return value and its meaning.
ENVIRONMENT – Relevant environment variables.
FILES – Files used or affected by the command.
EXAMPLES – Usage examples.
SEE ALSO – Related commands or functions.
BUGS – Known bugs or limitations.
AUTHORS – Credits for the program or documentation.
COPYRIGHT / LICENSE – Licensing information.
```

-> So it makes sense to search for that keywords, in case a certain category of information is desired.

The man command knows different sections, section 1 (user commands) and section 8 (sys admin commands) are the most important ones for us:

man sections are:

1. User Commands – Executable programs (e.g., ls, cp)
2. System Calls – Kernel-level functions (e.g., open, read)
3. Library Calls – Functions from libraries (e.g., printf, malloc)
4. Special Files – Device files and drivers (e.g., /dev/null)
5. File Formats – Configuration files and formats (e.g., /etc/passwd)
6. Games – Games and screensavers
7. Miscellaneous – Conventions, standards, and macro packages (e.g., man(7))
8. System Administration – Admin commands (e.g., mount, iptables)
9. Kernel Routines – (Rarely used on most systems)

Examples:

```bash
man 1 ls # user command 'ls'
man 2 open # syscall to open a file
man 3 printf # C lib for formatted printing
man 4 null # a black-hole-like device
man 5 crontab # a file to schedule tasks
man 6 nethack # a famous game
man 7 signal # POSIX signal names.
man 8 ip # a very common network command
man 9 printk # https://manpages.debian.org/jessie/linux-manual-3.16/printk.9.en.html
```

Searching for man pages:

```
$ man -k '.*pdump.*'
tcpdump (8)          - dump traffic on a network
utmpdump (1)         - dump UTMP and WTMP files in raw format

$ apropos tcpdump
tcpdump (8)          - dump traffic on a network
```

## Important commands

```bash
whoami
pwd # print working directory - where am I?
cd # navigate to my users home folder
cd / # navigate to the root directory
cd <some dir>
ls -lah # list contents of current directory with details and human readable size  
who # is logged on at this system
date # query server time
touch filename # create file name
cat filename # print contents of the file
cat /etc/os-release # what kind of system am I logged into 
uname -a # similar
ps -aux # what processes are running at the moment
```

## The difference between a program and a script:

```bash
$ cat prog.c
#include <stdio.h>
#include <unistd.h>
#include <limits.h>

int main() {
    char cwd[PATH_MAX];
    if (getcwd(cwd, sizeof(cwd)) != NULL) {
        printf("Current working dir: %s\n", cwd);
    } else {
        perror("getcwd() error");
        return 1;
    }
    return 0;
}

$ gcc prog.c -o prog
$ hexdump -C prog | head
00000000  7f 45 4c 46 02 01 01 00  00 00 00 00 00 00 00 00  |.ELF............|
00000010  03 00 3e 00 01 00 00 00  c0 10 00 00 00 00 00 00  |..>.............|
00000020  40 00 00 00 00 00 00 00  20 37 00 00 00 00 00 00  |@....... 7......|
00000030  00 00 00 00 40 00 38 00  0d 00 40 00 1f 00 1e 00  |....@.8...@.....|
00000040  06 00 00 00 04 00 00 00  40 00 00 00 00 00 00 00  |........@.......|
00000050  40 00 00 00 00 00 00 00  40 00 00 00 00 00 00 00  |@.......@.......|
00000060  d8 02 00 00 00 00 00 00  d8 02 00 00 00 00 00 00  |................|
00000070  08 00 00 00 00 00 00 00  03 00 00 00 04 00 00 00  |................|
00000080  18 03 00 00 00 00 00 00  18 03 00 00 00 00 00 00  |................|
00000090  18 03 00 00 00 00 00 00  1c 00 00 00 00 00 00 00  |................|
$ ./prog
Current working dir: /home/peterpan
```

-> '.ELF' is the Linux 'magic number'

```bash
$ pwd
/mnt/c/Users/XXX.XXX/Downloads
s$ hexdump -C Postman-win64-Setup.exe | head
00000000  4d 5a 90 00 03 00 00 00  04 00 00 00 ff ff 00 00  |MZ..............|
00000010  b8 00 00 00 00 00 00 00  40 00 00 00 00 00 00 00  |........@.......|
00000020  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000030  00 00 00 00 00 00 00 00  00 00 00 00 18 01 00 00  |................|
00000040  0e 1f ba 0e 00 b4 09 cd  21 b8 01 4c cd 21 54 68  |........!..L.!Th|
00000050  69 73 20 70 72 6f 67 72  61 6d 20 63 61 6e 6e 6f  |is program canno|
00000060  74 20 62 65 20 72 75 6e  20 69 6e 20 44 4f 53 20  |t be run in DOS |
00000070  6d 6f 64 65 2e 0d 0d 0a  24 00 00 00 00 00 00 00  |mode....$.......|
00000080  56 fc 57 fd 12 9d 39 ae  12 9d 39 ae 12 9d 39 ae  |V.W...9...9...9.|
00000090  59 e5 3a af 17 9d 39 ae  59 e5 3c af 9e 9d 39 ae  |Y.:...9.Y.<...9.|
```
-> 'MZ' is the Windows/DOS 'magic number'

```bash
$ cat myls.sh
#/bin/env bash

echo "Current working dir: $(pwd)"

$ ./myls.sh
Current working dir: /home/peterpan
```

-> '#!' identifies a script


```
+-------------------------------------+
| ALL COMMANDS YOU RUN ON AN OS ARE   |
| JUST EXECUTABLES OR SCRIPTS         |
+-------------------------------------+
```

```
+------------------------------------------+
| ULTIMATELY, ALL OF THEM USE SO CALLED    |
|        SYSTEM CALLS                      |
| TO IMPLEMENT THEIR TASKS                 |
+------------------------------------------+
```


---

In Linux, almost everything is accessible as a file, following the "everything is a file" philosophy — but there are exceptions:

Not accessible as regular files:

- Network sockets (e.g. TCP/UDP connections):

  - You can interact via descriptors (e.g., via send()/recv()), but they don’t appear as regular files. 
  - /proc/net/ gives info, but not direct socket access.

- Process memory:
  - Parts are exposed (e.g., /proc/[pid]/mem), but access may be restricted and not everything is mapped cleanly.

- Hardware interrupts / CPU instructions:
 - You can't trigger an IN/OUT instruction or RDTSC through a file interface — requires CPU-level access.

- Kernel-internal resources:
  - Scheduler queues, locks, internal kernel data structures — not all are exposed via /proc or /sys.

-> Most resources are accessible through the file abstraction (e.g., /dev, /proc, /sys), but not everything — especially ephemeral or privileged low-level resources.

---

---

PL APIs for syscalls:

Java:
https://java-native-access.github.io/jna/5.13.0/javadoc/ \
https://docs.oracle.com/en/java/javase/17/docs/specs/jni/index.html

Python: \
https://docs.python.org/3/library/os.html

Golang: \
https://pkg.go.dev/syscall