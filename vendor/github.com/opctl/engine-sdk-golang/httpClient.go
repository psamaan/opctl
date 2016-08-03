package opctlengine

import (
  goHttp "net/http"
  "io"
  "encoding/json"
  "io/ioutil"
  "errors"
)

type httpClient interface {
  Do(
  req *goHttp.Request,
  v interface{},
  ) (resp *goHttp.Response, err error)
}

func newHttpClient(
goHttpClient *goHttp.Client,
) httpClient {

  return &_httpClient{
    goHttpClient:goHttpClient,
  }

}

type _httpClient struct {
  goHttpClient *goHttp.Client
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (this _httpClient) Do(
httpRequest *goHttp.Request,
v interface{},
) (httpResponse *goHttp.Response, err error) {

  httpResponse, err = this.goHttpClient.Do(httpRequest)
  if err != nil {
    return
  }

  if (httpResponse.StatusCode > 300) {
    var bodyBytes []byte
    bodyBytes, err = ioutil.ReadAll(httpResponse.Body)
    err = errors.New(string(bodyBytes))
    return
  }

  if v != nil {

    if w, ok := v.(io.Writer); ok {
      io.Copy(w, httpResponse.Body)
    } else {
      err = json.NewDecoder(httpResponse.Body).Decode(v)
      if err == io.EOF {
        err = nil // ignore EOF errors caused by empty response body
      }
    }

  }

  defer httpResponse.Body.Close()

  return

}
