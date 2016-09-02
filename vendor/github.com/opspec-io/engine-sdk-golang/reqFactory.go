package opctlengine

import (
  "net/url"
  netHttp "net/http"
  "io"
  "bytes"
  "encoding/json"
  "github.com/opspec-io/engine-sdk-golang/ports"
  "fmt"
)

type reqFactory interface {
  Construct(
  method,
  relUrl string,
  body interface{},
  ) (req *netHttp.Request, err error)
}

func newReqFactory(
host ports.Host,
) reqFactory {

  return &_reqFactory{
    host:host,
  }

}

type _reqFactory struct {
  host ports.Host
}
//  If specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (this *_reqFactory) Construct(
method,
relUrlStr string,
body interface{},
) (req *netHttp.Request, err error) {

  hostname, err := this.host.GetHostname()
  if (nil != err) {
    return
  }

  var u *url.URL
  u, err = url.Parse(fmt.Sprintf("http://%v:42224/%v", hostname, relUrlStr))
  if err != nil {
    return
  }

  var buf io.ReadWriter
  if body != nil {
    buf = new(bytes.Buffer)
    err = json.NewEncoder(buf).Encode(body)
    if (nil != err) {
      return
    }
  }

  req, err = netHttp.NewRequest(method, u.String(), buf)

  return
}
