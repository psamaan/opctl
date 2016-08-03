package models

func NewRunOpReq(
args map[string]string,
opUrl string,
) *RunOpReq {

  return &RunOpReq{
    Args:args,
    OpUrl:opUrl,
  }

}

type RunOpReq struct {
  Args  map[string]string `json:"args"`
  OpUrl string `json:"opUrl"`
}
