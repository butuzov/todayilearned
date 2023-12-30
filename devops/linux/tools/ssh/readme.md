<!-- toc: true -->
<!-- menu: "`ssh`" -->
# `ssh` Secure Shell

SSH, also known as Secure Shell or Secure Socket Shell, is a network protocol that gives users, particularly system administrators, a secure way to access a computer over an unsecured network.

## `ssh`

```bash
# ssh into <hostname> using <user> and its private key (-i)
ssh <user>@<hostname> -i ~/.ssh/id_rsa
# ssh and run command & exit!
ssh <user>@<hostname>  uname -a
# ssh and run command (with quotes) & exit!
ssh <user>@<hostname>  'echo "This is cool"'
# opens secret tunnel!
ssh <user>@<hostname> -Nfl 3000:localhost:8080
# bastion ssh
ssh <user>@<hostname> "ssh <user>@<internal.hostname> 'echo 1'"
```

## `ssh-keygen`

Authentication key generation, management and conversion

```bash
# Removes keys from hostname
ssh-keygen -R <hostname>
# Generates Keys
ssh-keygen -t ed25519 -C "example@example.com" -f ./infra/ssh_key
```


## `ssh-copy-id`

Copy public key onto host.

```bash
ssh-copy-id -i ~/.ssh/id_rsa.pub <user>@<hostname>
# if no ssh-copy-id - we can use this shortcut.
cat ~/.ssh/id_rsa.pub | ssh <user>@<hostname> 'cat >> ~/.ssh/authorized_keys'
```

## `ssh-agent`

The ssh-agent is a helper program that keeps track of users' identity keys and their passphrases.

```shell
# If ssh-agent is not automatically started at login, it can be started manually with the command
eval `ssh-agent`
#The easiest way to check is to check the value of the SSH_AGENT_SOCK environment variable.
# If it is set, then the agent is presumably running. It can be checked by
echo $SSH_AGENT_SOCK
```


## `ssh-add`

To add an arbitrary private key, give the path of the key file as an argument to `ssh-add`

```shell
ssh-add ~/.ssh/id_rsa
```

## `scp`

`scp` is a program for copying files between computers.

```shell
# scp [options] <user>@<src-host>:dir/file <user>@<dst-host>:dir/file
ssh <user>@<hostname> "scp <user>@<internal.hostname>:~/src.tar.gz ~/dst.tar.gz"
# The bandwidth is specified in Kbit/sec - so this is 1mb per sec (its 20 mb persec by default)
scp -l 8000 <user>@<hostname>:/home/user/* .
# copy todo.txt into /home/user
scp ~/todo.txt <user>@<hostname>:/home/user/
# recursive copy
scp -r ~/dir <user>@<hostname>:~/dir

```

### `pscp` (windows)

`pscp` is a shell command that works almost on Windows Shell the same way that scp works on Linux or Mac OS X

## `sshd`

`sshd` (OpenSSH Daemon) is the daemon program for ssh(1). Together these programs replace rlogin(1) and rsh(1), and provide secure encrypted communications between two untrusted hosts over an insecure network.

### `sshd_config`

The `sshd_config` file is an ASCII text based file where the different configuration options of the SSH server are indicated and configured with keyword/argument pairs. Arguments that contain spaces are to be enclosed in double quotes (").

```conf
Port                   2278                    # port
PermitRootLogin        no                      # no root login
PasswordAuthentication no                      # turn off - passwords
PubkeyAuthentication   yes                     # turn on  - ssh keys only
AuthorizedKeysFile     %h/.ssh/authorized_keys # Public Keys location
```


- [ssh.com: What is SSH (Secure Shell)](https://www.ssh.com/academy/ssh)
- [ssh.com: How to Configure the OpenSSH Server?](https://www.ssh.com/academy/ssh/sshd_config)


## Other Tooling

- https://mosh.org/
- https://www.putty.org/
