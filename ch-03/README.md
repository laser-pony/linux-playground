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
