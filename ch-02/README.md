# Linux Playground Chapter 2

---

## In this playground we are going to

  - Learn how to use different
    - text processing tools / stream editors/ filters.
    - the common 'trio' here is sed, grep and awk.
    - I will also talk about vim, as this editor can deal with huge files and exists on every unix-like system (posix say it shold, [1], https://en.wikipedia.org/wiki/POSIX)  

[1]: https://pubs.opengroup.org/onlinepubs/9699919799/utilities/vi.html?utm_source=chatgpt.com

## Out testfile for learning purposes

```
15:02:25 NORMAL   Job: 12345 [groceries] started job, budget=13
15:02:26 NORMAL   Job: 12345 [groceries] added sausages
15:02:27 NORMAL   Job: 12345 [groceries] added ketchup
15:02:28 NORMAL   Job: 12345 [groceries] added chips
15:02:29 NORMAL   Job: 12345 [groceries] skipped salad
15:02:30 NORMAL   Job: 12345 [groceries] vitamins unreachable 404
15:02:31 NORMAL   Job: 12345 [groceries] finished in 123456 ms
15:02:32 NORMAL   Job: 12346 [office] started job, budget=1000
15:02:33 NORMAL   Job: 12346 [office] solved ticket nb. 12
15:02:34 NORMAL   Job: 12346 [office] solved ticket nb. 13
15:02:35 NORMAL   Job: 12346 [office] tested ticket nb. 12 successfully
15:02:36 NORMAL   Job: 12346 [office] tested ticket nb. 13 with errors
15:02:37 NORMAL   Job: 12346 [office] finished in 123 ms
```

-> Please save this as `test-logfile.log`

## grep

grep is a very, very efficient programm using the Boyer-Moore string searching algorithm. There are some fancy new implementations (ripgrep being one of these), but from an algorithmic and practical point of view grep could be considered a worthy benchmark.

The syntax is:

```text
grep [OPTION...] PATTERNS [FILE...]

# @see
# man grep
```

grep has several standard options, that should be know - as with any unix program, you could find all infos (much more than you usually want or need) in its man page (https://en.wikipedia.org/wiki/Man_page).

- -H : Always prints the filename with the matching line, even if only one file is searched.
- -a : Process a binary file as if it were text (avoids "Binary file matches" messages). I use this if I desperately search a certain string, e.g. also in compiled classes.
- -i : Ignore case when matching (abc matches ABC, etc.).
- -n : Prefix each line of output with its line number.
- -R : Recursively search all files and subdirectories (follows symbolic links).

I call these the 'HainR' options (similar to the German name 'Heiner'). If you include '-F' for literal matches (i.e. the following pattern is not interpreted as regex, but as a literal string), or '-v' to invert the match, there is another, less kosher mnemonic.

Another useful combination of options is

- -o : only print the matched
- -P : use Perl compatible regex

-> I usually search for a simple text, usually combining that search with the ignore-case-flag, or, when using Patterns, I tend to use the perl-compatile variants, as I perceive them as more versatile (https://en.wikipedia.org/wiki/Perl_Compatible_Regular_Expressions)

### Basic Examples

```bash
$ pwd
    ~/linux-schulung/ch-02

$ ls -1
    README.md
    subdir

$ grep -H ad.ed subdir/test-logfile.log
    subdir/test-logfile.log:15:02:26 NORMAL   Job: 12345 [groceries] added sausages
    subdir/test-logfile.log:15:02:27 NORMAL   Job: 12345 [groceries] added ketchup
    subdir/test-logfile.log:15:02:28 NORMAL   Job: 12345 [groceries] added chips

$ grep -Hi AD.ED subdir/test-logfile.log
    subdir/test-logfile.log:15:02:26 NORMAL   Job: 12345 [groceries] added sausages
    subdir/test-logfile.log:15:02:27 NORMAL   Job: 12345 [groceries] added ketchup
    subdir/test-logfile.log:15:02:28 NORMAL   Job: 12345 [groceries] added chips

$ grep -Hin AD.ED subdir/test-logfile.log
    subdir/test-logfile.log:2:15:02:26 NORMAL   Job: 12345 [groceries] added sausages
    subdir/test-logfile.log:3:15:02:27 NORMAL   Job: 12345 [groceries] added ketchup
    subdir/test-logfile.log:4:15:02:28 NORMAL   Job: 12345 [groceries] added chips

$ grep -HinF AD.ED subdir/test-logfile.log
# no match here

$ grep -HiRnF ADDED subdir
    subdir/test-logfile.log:2:15:02:26 NORMAL   Job: 12345 [groceries] added sausages
    subdir/test-logfile.log:3:15:02:27 NORMAL   Job: 12345 [groceries] added ketchup
    subdir/test-logfile.log:4:15:02:28 NORMAL   Job: 12345 [groceries] added chips
    README.md:18:15:02:26 NORMAL   Job: 12345 [groceries] added sausages
    README.md:19:15:02:27 NORMAL   Job: 12345 [groceries] added ketchup
    README.md:20:15:02:28 NORMAL   Job: 12345 [groceries] added chips


# finding tasks, that have a 3 digit budged (or more)
$ grep -iP 'budget=\d{3,}' subdir/test-logfile.log
15:02:32 NORMAL   Job: 12346 [office] started job, budget=1000 
```

There are two commands, that often make sense together

- sort - sorts all lines in a file, dictionary order
- uniq - deletes subsequent identical lines
  - when used with the '-c' option, for 'count', it counts subsequent lines

Extending our examples:

```bash
# printing all the names in square brackets
# match up to a closing bracket, or up to a comma:
$ grep -ioP '\[[^,]*?(\,|\]) $ grep -ioP '\[[^,]*?(\,|\])' subdir/test-logfile.log
    [groceries]
    [groceries]
    [groceries]
    [groceries]
    [groceries]
    [groceries]
    [groceries]
    [office]
    [office]
    [office]
    [office]
    [office]
    [office]

# sort and count these
$ grep -ioP '\[[^,]*?(\,|\])' subdir/test-logfile.log | sort | uniq -c
      7 [groceries]
      6 [office]
```

So we got 7 entries for the groceries task, and 6 for the office task.

We can even build a statistic on how long these tasks took:
```bash
# task took more than 2h 47 minutes:
$ grep -iP 'finished in \d{8,}' subdir/test-logfile.log
15:02:37 NORMAL   Job: 12346 [office] finished in 28800000 ms

# how many lines did tasks log, that are not work-related?
$ grep -v office subdir/test-logfile.log | wc -l
7
```

## sed (and awk)

sed is a tool I primarily use to substitute certain lines in template files, when trying to automate installation/configuration.

awk is a very versatile tool, but I don't use it very much. The reason is, that I feel more comfortable writing a small Python or Golang program when takling tasks that could be done by awk (i.e. I feel that, although it has more features than sed, when the need arises to use it, I'm better off using the 'big gun' of a full fledged programming language right away).

Very short awk tutorial: https://www.w3schools.com/bash/bash_awk.php \
In-depth awk tutorial: https://www.tutorialspoint.com/awk/index.htm

We would like to give two examples here:
1. Using sed and awk to filter a certain section of a file by a beginning and ending regex.
2. Using sed to substitute a piece of text in a file.

**Example: awk - cut out text with beginning and ending pattern**
```bash
$ awk '/15:02:25/ {if (!start) start=NR} /15:02:31/ {end=NR} {lines[NR]=$0} END {for (i=start; i<=end; i++) print lines[i]}' test-logfile.log
15:02:25 NORMAL   Job: 12345 [groceries] started job, budget=13
15:02:26 NORMAL   Job: 12345 [groceries] added sausages
15:02:27 NORMAL   Job: 12345 [groceries] added ketchup
15:02:28 NORMAL   Job: 12345 [groceries] added chips
15:02:29 NORMAL   Job: 12345 [groceries] skipped salad
15:02:30 NORMAL   Job: 12345 [groceries] vitamins unreachable 404
15:02:31 NORMAL   Job: 12345 [groceries] finished in 1234567 ms
```

That particular case, is much easier in sed:

```bash
# -n means: suppress automatic printing of pattern space
# means just print when explicitly asked to, i.e. via the 'p' command
$ sed -n '/15:02:25/,/15:02:31/p' test-logfile.log
15:02:25 NORMAL   Job: 12345 [groceries] started job, budget=13
15:02:26 NORMAL   Job: 12345 [groceries] added sausages
15:02:27 NORMAL   Job: 12345 [groceries] added ketchup
15:02:28 NORMAL   Job: 12345 [groceries] added chips
15:02:29 NORMAL   Job: 12345 [groceries] skipped salad
15:02:30 NORMAL   Job: 12345 [groceries] vitamins unreachable 404
15:02:31 NORMAL   Job: 12345 [groceries] finished in 1234567 ms
```

As we can see, awk offers quite some features:
- variables
- control structures
- each section follows the following syntax:
  - \<matching condition\> {code-block}\*
- In our case there are three:
```awk
/15:02:25/ {if (!start) start=NR}
/15:02:31/ {end=NR}
END {for (i=start; i<=end; i++) print lines[i]}
```

The downsides, imho, are:
- Implicit knowledge is needed, e.g.
  - 'END' matches the end-condition of the awkscript.
  - NR represents the current line number during execution of the script.
  - if variables are not defined, !varname could be used to initialize it.
  - commands like 'print' needed to be know.
- This is all fine, but the amount you need to know by heart imho almost reaches a programming language.
- It is easier to develop, maintain and understand a common PL, like Java, Python or Golang - back in the old unix times, when text processing was a more common task, the user base of awk was bigger, so the downsides of today did not hold back then, or were less grave.


**A program. Yes, it is much longer - but imho it could be read by more people, and the bounds, when you want to really extend it, are limitless compare compared to awk**

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const startPattern = "15:02:25"
	const endPattern = "15:02:31"

	var lines []string
	start, end := -1, -1

	scanner := bufio.NewScanner(os.Stdin)
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if strings.Contains(line, startPattern) && start == -1 {
			start = lineNum
		}
		if strings.Contains(line, endPattern) {
			end = lineNum
		}
		lineNum++
	}

	if start != -1 && end != -1 && start <= end {
		for i := start; i <= end && i < len(lines); i++ {
			fmt.Println(lines[i])
		}
	} else {
		fmt.Fprintln(os.Stderr, "Start or end pattern not found or in wrong order.")
	}
}
```

### (My) Most Common Use Case of sed

```bash
# -i means 'in-place' modification of the file
sed -i 's/10.99.133.181/hostname/' some-config.xml
```

## Vim

I would like to show you something live today, especially in comparison with another editor:

```bash
# borrowed from SO:
$ tr -dc "A-Za-z 0-9" < /dev/urandom | fold -w100|head -n 10000000 > bigfile.txt
$ ls -lh bigfile.txt
    -rwxrwxrwx 1 peterpan peterpan 964M May  1 17:47 bigfile.txt
```
