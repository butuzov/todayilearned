# wget

[wget](https://www.gnu.org/software/wget/) - &nbsp;retrieving files using HTTP, HTTPS, FTP and FTPS.

```shell
# Pipe Content Into stdout
wget http://example.com --output-document - | head -n4

# Continue Partial Download
wget --continue https://example.com/linux-distro.iso

# Download Sequence (expansion?)
wget http://example.com/file_{1..4}.webp

# Mirror
wget --mirror --convert-links --adjust-extension --page-requisites \
    --no-parent http://example.org

wget -mkEpnp http://example.org # Same command but shorter

# Only Specific files
wget -e robots=off -r -l1 -A ".pdf" http://example.org

# Headers
wget --debug example.com
wget --debug --header="User-Agent: Mozilla/5.0" http://example.com

# Follow Redirect
wget --max-redirect 0 http://iana.org

# Expand SHort URL
wget --max-redirect 0 "https://bit.ly/2yDyS4T"

# INternal Docker Helthcheck
wget --no-verbose --tries=1 --spider '0.0.0.0:9090/-/healthy' || exit 1
```
