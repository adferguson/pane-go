// Participatory Networking. Copyright (C) 2012-2013 Brown University
//
// Author: Andrew Ferguson <adf@cs.brown.edu>
//

package pane

/***************************************************************************************************
 *
 * helper functions to mirror proto.String(), proto.Int32(), etc.
 *
 **************************************************************************************************/

func ThriftString(s string) *string {
  p := new(string)
  *p = s
  return p
}

func ThriftInt32(i int32) *int32 {
  p := new(int32)
  *p = i
  return p
}

/***************************************************************************************************
 *
 * in an idea world, the Thrift generator would make these for us based on the generatorEnum
 * code in goprotobuf/protoc-gen-go/generator/generator.go
 *
 **************************************************************************************************/

func (e HintType) Enum() *HintType {
  p := new(HintType)
  *p = e
  return p
}

func (e IPAddrType) Enum() *IPAddrType {
  p := new(IPAddrType)
  *p = e
  return p
}

func (e PrivilegeType) Enum() *PrivilegeType {
  p := new(PrivilegeType)
  *p = e
  return p
}

func (e QueryType) Enum() *QueryType {
  p := new(QueryType)
  *p = e
  return p
}

func (e RequestType) Enum() *RequestType {
  p := new(RequestType)
  *p = e
  return p
}

func (e Result) Enum() *Result {
  p := new(Result)
  *p = e
  return p
}

func (e TimeType) Enum() *TimeType {
  p := new(TimeType)
  *p = e
  return p
}

func (e TransportProto) Enum() *TransportProto {
  p := new(TransportProto)
  *p = e
  return p
}
