## On Linux

```shell
sudo curl -L -o /usr/local/bin/doctl https://github.com/doctl/cli/releases/download/0.1.0/doctl-linux-x86_64 && \
sudo chmod +x /usr/local/bin/doctl
```

## On OSX

```shell
sudo curl -L -o /usr/local/bin/doctl https://github.com/doctl/cli/releases/download/0.1.0/doctl-osx-x86_64 && \
sudo chmod +x /usr/local/bin/doctl
```

## On Windows with git bash:

```shell
if [[ ! -d "$HOME/bin" ]]; then mkdir -p "$HOME/bin"; fi && \
curl -L https://github.com/doctl/cli/releases/download/0.1.0/doctl-windows-x86_64.exe > "$HOME/bin/doctl.exe" && \
chmod +x "$HOME/bin/doctl.exe"
```
