package models

func NewOpSummaryView(
name string,
_type string,
) *OpSummaryView {

  return &OpSummaryView{
    Name:name,
    Type:_type,
  }

}

type OpSummaryView struct {
  Name string
  Type string
}