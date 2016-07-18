## On Linux

```shell
curl -L -o /usr/local/bin/opctl https://github.com/opctl/cli/releases/download/0.1.5/opctl-linux-x86_64 && \
chmod +x /usr/local/bin/opctl
```

## On OSX

```shell
curl -L -o /usr/local/bin/opctl https://github.com/opctl/cli/releases/download/0.1.5/opctl-darwin-x86_64 && \
chmod +x /usr/local/bin/opctl
```

## On Windows with git bash:

```shell
if [[ ! -d "$HOME/bin" ]]; then mkdir -p "$HOME/bin"; fi && \
curl -L https://github.com/opctl/cli/releases/download/0.1.5/opctl-windows-x86_64.exe > "$HOME/bin/opctl.exe" && \
chmod +x "$HOME/bin/opctl.exe"
```
