package models

func NewAddSubOpReq(
projectUrl string,
opName string,
subOpUrl string,
precedingSubOpUrl string,
) *AddSubOpReq {

  return &AddSubOpReq{
    ProjectUrl:projectUrl,
    OpName :opName,
    SubOpUrl :subOpUrl,
    PrecedingSubOpUrl :precedingSubOpUrl,
  }

}

type AddSubOpReq struct {
  ProjectUrl        string `json:"-"`
  OpName            string `json:"-"`
  SubOpUrl          string
  PrecedingSubOpUrl string
}
