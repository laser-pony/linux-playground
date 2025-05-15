# Linux Playground Chapter 4

In this chapter we are going to cover the basics of systemd and the Vim text editor.

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

- systemd is the first 'process' that is started on most modern linux systems.
  - I say 'most systems' because there is an old-(fashioned) flame war, if systemd is the right way to go compared to initd or rc.d
- Most people interact with systemd with the systemctl command, but 
  - a family of other programs is related to systemd,
  - and systemd starts and stops so called 'units', not all of them are processes.

### Examples


Query the status of a service:

```bash
systemctl status some.service
```

Start some service:

```bash
systemctl start some.service
```

Stop some service:

```bash
systemctl stop some.service
```

Add some service from the ones who should be available / start automatically:

```bash
systemctl enable some.service
```

Remove some service from the ones who should be available / start automatically:

```bash
systemctl disable some.service
```

Query 'enabled-status' of a service:

```bash
systemctl is-enabled update-on-boot.service
```

List units of type service; regex filter the results:

```bash
systemctl list-units --type=service | grep -E 'cups|docker'
```

## Units that usually correspond to processes / do not correspond

- service — these manage long-running daemons or one-shot tasks.
- socket - not directly a process, but it can spawn one (via socket activation).
- timer - schedules a service to run later — the timer isn't a process, but it triggers one.
- scope — wraps an external process into systemd's supervision.
- busname - activates a service when a D-Bus name is requested — indirectly leads to a process.

### Units that do not correspond to processes

- target - purely logical groupings — like runlevels. No process.
- mount - mounts a filesystem — managed by the kernel. No persistent process.
- automount - configures kernel-based on-demand mounts. No process unless activated.
- swap - manages swap devices/files. Kernel-level; no daemon.
- device - represents a device in /sys via udev. Not a process.
- path - watches a file or dir for changes. Uses inotify/fanotify internally, not a full process.
- slice - Organizes units into cgroup slices for resource control — again, not a process itself.


We will 'kinda' work towards this status, and use this tutorial to do so: https://linuxhandbook.com/create-systemd-services/

```bash
$ cat /root/.scripts/sys-update.sh
#!/usr/bin/env bash
if [ ${EUID} -ne 0 ]
then
        exit 1 # this is meant to be run as root
fi
echo 'i was running' >> /root/logs/sys-update.log
```
```bash
$ cat /etc/systemd/system/update-on-boot.service
[Unit]
Description=Keeping my sources fresher than Arch Linux
After=multi-user.target
# Requires=, Wants=, After=

[Service]
ExecStart=/usr/bin/bash /root/.scripts/sys-update.sh
Type=simple

[Install]
WantedBy=multi-user.target
```

### Difference between 'requires', 'after' and 'wants':


>Requires=

Declares a strong dependency. \
If the required unit fails to start or is missing, the current unit will also fail.

>Wants=

Declares a weak dependency. \
The wanted unit will be started, but not required to succeed.
Failure to start the other unit does not stop this one.

>After=

Controls start-up order only, not dependency. \
Ensures this unit starts after the listed one — doesn't imply starting it.

systemctl is-enabled update-on-boot.service


>WantedBy=

Tells systemctl enable where to symlink this unit, to make it part of a boot target.

We want to make systemd aware of our service file:

```bash
systemctl daemon-reload
```
---

```
STD-1247:/etc/systemd # pwd
/etc/systemd
STD-1247:/etc/systemd # tree -L 3
.
├── journald.conf
├── logind.conf
├── network
├── pstore.conf
├── sleep.conf
├── system
│   ├── dbus-org.opensuse.Network.AUTO4.service -> /usr/lib/systemd/system/wickedd-auto4.service
│   ├── dbus-org.opensuse.Network.DHCP4.service -> /usr/lib/systemd/system/wickedd-dhcp4.service
│   ├── dbus-org.opensuse.Network.DHCP6.service -> /usr/lib/systemd/system/wickedd-dhcp6.service
│   ├── dbus-org.opensuse.Network.Nanny.service -> /usr/lib/systemd/system/wickedd-nanny.service
│   ├── default.target.wants
│   │   └── ca-certificates.path -> /usr/lib/systemd/system/ca-certificates.path
│   ├── getty.target.wants
│   │   ├── container-getty@1.service -> /usr/lib/systemd/system/container-getty@.service
│   │   ├── container-getty@2.service -> /usr/lib/systemd/system/container-getty@.service
│   │   └── getty@tty1.service -> /usr/lib/systemd/system/getty@.service
│   ├── multi-user.target.wants
...
│   │   ├── cups.path -> /usr/lib/systemd/system/cups.path
│   │   ├── docker.service -> /usr/lib/systemd/system/docker.service
│   │   ├── kbdsettings.service -> /usr/lib/systemd/system/kbdsettings.service
│   │   ├── some.service -> /usr/lib/systemd/system/some.service
│   │   ├── some_stop.service -> /usr/lib/systemd/system/some_stop.service
│   │   ├── pre_startup_some.service -> /usr/lib/systemd/system/pre_startup_some.service
│   │   ├── smb.service -> /usr/lib/systemd/system/smb.service
│   │   ├── sshd.service -> /usr/lib/systemd/system/sshd.service
│   │   ├── startup_some.service -> /usr/lib/systemd/system/startup_some.service
...
│   ├── network-online.target.wants
│   │   └── wicked.service -> /usr/lib/systemd/system/wicked.service
│   ├── network.service -> /usr/lib/systemd/system/wicked.service
│   ├── printer.target.wants
...
├── system.conf
├── system-generators
│   └── lxc
├── timesyncd.conf
├── user
│   ├── basic.target.wants
│   │   └── systemd-tmpfiles-setup.service -> /usr/lib/systemd/user/systemd-tmpfiles-setup.service
│   └── timers.target.wants
│       └── systemd-tmpfiles-clean.timer -> /usr/lib/systemd/user/systemd-tmpfiles-clean.timer
└── user.conf

13 directories, 40 files
```



