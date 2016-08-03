package ports

//go:generate counterfeiter -o ../adapters/host/fake/host.go --fake-name FakeHost ./ Host

type Host interface {
  EnsureRunning(
  ) (err error)

  GetHostname(
  ) (
  hostname string,
  err error,
  )
}
