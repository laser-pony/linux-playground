# Linux Playground Chapter 4

In this chapter we are going to cover the basics of systemd and the Vim text editor.

- Know important resources to look up infos about systemd
- Understand how to set up a hello-world like service

## systemd

Official systemd resources: \
https://github.com/systemd/systemd \
https://systemd.io/ 

Imho the wikipedia article is good: \
https://en.wikipedia.org/wiki/Systemd

Get help about systemd:

```bash
man systemd.service
man systemd.unit
man systemd.exec
```

### Good to know 

- systemd is the first 'process' that is started on most modern linux systems.
  - I say 'most systems' because there is an old-(fashioned) flame war, if systemd is the right way to go compared to init.d, OpenRC (some Linux distros) or rc.d (BSD)
- Most people interact with systemd with the systemctl command, but a family of other programs is related to systemd.
- **systemd starts and stops so called 'units', not all of them are processes.**
- Every **unit** has its own unit description file, which obeys a certain syntax, and determines 
  - at what stage of the boot/start process, certain units are started or stopped,
  - and what dependencies relate different units.

### Will be clearer later in this chapter, but is a usefull table anyways

| SysV Runlevel | systemd Target      |
| ------------- | ------------------- |
| 0             | `poweroff.target`   |
| 1             | `rescue.target`     |
| 3             | `multi-user.target` |
| 5             | `graphical.target`  |
| 6             | `reboot.target`     |

I.e. this maps commands like `init <some-number>` to systemd _targets_.


### Examples using systemctl

```bash
# Query the status of a service:
systemctl status some.service

# Start some service:
systemctl start some.service

# Stop some service:
systemctl stop some.service

# Add some service to the ones which should start automatically on 'system boot':
# (the 'runlevel' is determined by the service/unit description file)
systemctl enable some.service

# Remove some service from the ones who should be available / start automatically:
systemctl disable some.service

# Query 'enabled-status' of a service:
systemctl is-enabled update-on-boot.service

#List units of type service; regex filter the results:
systemctl list-units --type=service | grep -E 'cups|docker'
```


## Units that usually correspond to processes

- service — these manage long-running daemons or one-shot tasks.
- socket - not directly a process, but it can spawn one (via socket activation).
- timer - schedules a service to run later — the timer isn't a process, but it triggers one.
- scope — wraps an external process into systemd's supervision.
- busname - activates a service when a D-Bus name is requested — indirectly leads to a process.

### Units that do not correspond to processes

- target - purely logical groupings, usually of services — like runlevels. No process.
- mount - mounts a filesystem — managed by the kernel. No persistent process.
- automount - configures kernel-based on-demand mounts. No process unless activated.
- swap - manages swap devices/files. Kernel-level; no daemon.
- device - represents a device in /sys via udev. Not a process.
- path - watches a file or dir for changes. Uses inotify/fanotify internally, not a full process.
- slice - Organizes units into cgroup slices for resource control — again, not a process itself.

---

** We are only going to deal with services in this tutorial **

---

We will 'kinda' work towards this status, and use this tutorial to do so: https://linuxhandbook.com/create-systemd-services/

```bash
# Create a log- directory and file:
mkdir -p /root/logs && touch "$_"/sys-hello.sh

# This is how the actual 'program'/script looks like:
$ cat /root/.scripts/sys-hello.sh
#!/usr/bin/env bash
if [ ${EUID} -ne 0 ]
then
        exit 1 # this is meant to be run as root
fi
echo "I was running at $(date)" >> /root/logs/sys-hello.log

# This is how the unit file should look like: 
$ cat /etc/systemd/system/hello-on-boot.service
[Unit]
Description=Saying hello on system boot
After=multi-user.target
# Requires=, Wants=, After=

[Service]
ExecStart=/usr/bin/bash /root/.scripts/sys-hello.sh
Type=simple

[Install]
WantedBy=multi-user.target
```

### Difference between 'requires', 'after' and 'wants':


>Requires= \
>RequiredBy= 

Declares a strong dependency. \
If the required unit fails to start or is missing, the current unit will also fail.

>Wants= \
>WantedBy=  

Declares a weak dependency. \
The wanted unit will be started, but not required to succeed in starting.
I.e. failure to start the other unit does not stop this one.

>After= \
>Before= 

Controls start-up order only, not dependency. \
Ensures this unit starts after the listed one — doesn't imply starting it. \
Different resources on the web quote that 'order' is an orthogonal property to
'dependency', so the relevant situation is if and only if, both units need to 
be started. Usually that need is defined by some dependency reason, possibly to
a third unit or target.

** Where to put these? **

| Directive             | Valid in `[Unit]`? | `[Service]`? | `[Install]`? |
| --------------------- | ------------------ | ------------ | ------------ |
| `After=`, `Before=`   | ✅ Yes              | ❌ No         | ❌ No         |
| `Requires=`, `Wants=` | ✅ Yes              | ❌ No         | ❌ No         |
| `WantedBy=`           | ❌ No               | ❌ No         | ✅ Yes        |
| `RequiredBy=`         | ❌ No               | ❌ No         | ✅ Yes        |
| `Conflicts=`          | ✅ Yes              | ❌ No         | ❌ No         |
| `PartOf=`, etc.       | ✅ Yes              | ❌ No         | ❌ No         |


---

```bash
# make the systemd aware of our unit file
systemctl daemon-reload

# query if that unit file is already 'enabled' (= installed)
systemctl is-enabled update-on-boot.service

# enable it
systemctl enable update-on-boot.service

# successfully enabled
systemctl is-enabled update-on-boot.service

# query status
systemctl status update-on-boot.service

cat ~/logs/sys-hello.log
...
systemctl start update-on-boot.service
...
cat ~/logs/sys-hello.log

cd /etc/systemd
pwd
...
:/etc/systemd # tree -L 3

systemctl start multi-user.target
...
cat ~/logs/sys-hello.log
```

## vim

-> still a todo - walk through parts of vimtutor
