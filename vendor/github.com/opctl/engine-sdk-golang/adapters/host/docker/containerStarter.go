package docker

import (
  "fmt"
  "github.com/mitchellh/go-homedir"
  "os/exec"
)

type containerStarter interface {
  ContainerStart(
  image string,
  ) (err error)
}

func newContainerStarter(
pathNormalizer pathNormalizer,
) (containerStarter containerStarter) {

  containerStarter = _containerStarter{
    pathNormalizer:pathNormalizer,
  }

  return

}

type _containerStarter struct {
  pathNormalizer pathNormalizer
}

func (this _containerStarter) ContainerStart(
image string,
) (err error) {

  usersDir, err := homedir.Dir()
  if (nil != err) {
    return
  }

  normalizedUsersDir := this.pathNormalizer.Normalize(usersDir)

  dockerRunCmd :=
    exec.Command(
      "docker",
      "run",
      "-d",
      "-p",
      "42224:42224",
      "-v",
      fmt.Sprintf("%v:%v", normalizedUsersDir, normalizedUsersDir),
      "-v",
      fmt.Sprintf("%v/.docker/config.json:/root/.docker/config.json", normalizedUsersDir),
      "-v",
      "/var/run/docker.sock:/var/run/docker.sock",
      "--name",
      containerName,
      image,
    )

  _, err = dockerRunCmd.Output()

  return
}
