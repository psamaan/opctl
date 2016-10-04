>Note: these instructions work equally well for upgrading.

## On Linux

```shell
curl -L -o /usr/local/bin/opctl https://github.com/opspec-io/cli/releases/download/0.1.12/opctl-linux-x86_64 && \
chmod +x /usr/local/bin/opctl
```

## On OSX

> Docker for mac bind mount behavior (at time of writing) is
> [known to be unreliable](https://forums.docker.com/t/file-access-in-mounted-volumes-extremely-slow-cpu-bound/8076).
> Until Docker gets this resolved, the recommended approach is to use
> docker-machine and
> [docker-machine-nfs](https://github.com/adlogix/docker-machine-nfs)

```shell
curl -L -o /usr/local/bin/opctl https://github.com/opspec-io/cli/releases/download/0.1.12/opctl-darwin-x86_64 && \
chmod +x /usr/local/bin/opctl
```

## On Windows with git bash:

```shell
if [[ ! -d "$HOME/bin" ]]; then mkdir -p "$HOME/bin"; fi && \
curl -L https://github.com/opspec-io/cli/releases/download/0.1.12/opctl-windows-x86_64.exe > "$HOME/bin/opctl.exe" && \
chmod +x "$HOME/bin/opctl.exe"
```

