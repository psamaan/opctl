package models

func NewOpDetailedView(
description string,
name string,
subOps []OpSummaryView,
) *OpDetailedView {

  return &OpDetailedView{
    Description:description,
    Name:name,
    SubOps:subOps,
  }

}

type OpDetailedView struct {
  Description string
  Name        string
  SubOps      []OpSummaryView
}