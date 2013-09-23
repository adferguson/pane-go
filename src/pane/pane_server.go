// Participatory Networking. Copyright (C) 2012-2013 Brown University
//
// Author: Andrew Ferguson <adf@cs.brown.edu>
//

package pane

import (
  "sync"
)

/***************************************************************************************************
 *
 * PANE Server
 *
 **************************************************************************************************/

type ShareIndex string  // to facilitiate switching to uuid, or hash(name), etc.

type PaneServer struct {
  // map holding shares in the ShareTree. each share already contains the name of its parent
  // but we keep a second map (parent -> children) for convenience. both are protected by stLock.
  shareTree        map[ShareIndex]*Share
  subShares        map[ShareIndex][]ShareIndex
  stLock           sync.RWMutex;

  // map from a share to accepted requests. protected by requestsLock.
  acceptedRequests map[ShareIndex]*Request
  requestsLock     sync.RWMutex;
}

func (server *PaneServer) Init() {
  server.shareTree = make(map[ShareIndex]*Share)
  server.subShares = make(map[ShareIndex][]ShareIndex)
  server.acceptedRequests = make(map[ShareIndex]*Request)

  // initialize the RootShare

  rootShare := &Share {
    Id: &ShareID{ Name: ThriftString("RootShare"), },
    Principal: []*Principal{ &Principal{ User: ThriftString("Root") }},
    // TODO(adf): add all privileges
  }

  server.shareTree[rootShare.GetIndex()] = rootShare
}

func (share *Share) GetIndex() ShareIndex {
  return (ShareIndex)(*share.Id.Name)
}

func (share_id *ShareID) GetIndex() ShareIndex {
  return (ShareIndex)(*share_id.Name)
}

/***************************************************************************************************
 *
 * Administrative commands
 *
 **************************************************************************************************/

func (server *PaneServer) Authenticate(principal *Principal) (*GenericResponse, error) {
  return nil, nil
}

func (server *PaneServer) GrantShare(grant *Grant) (*GenericResponse, error) {
  return nil, nil
}

// TODO(adf): needs error handling
func (server *PaneServer) NewShare(share *Share) (*GenericResponse, error) {
  sid := share.GetIndex()
  pid := share.Parent.GetIndex()

  server.stLock.Lock()
  server.shareTree[sid] = share
  server.subShares[pid] = append(server.subShares[pid], sid)
  server.stLock.Unlock()  // TODO(adf): use defer?

  rv := GenericResponse { Result: ResultSuccess.Enum(), }

  return &rv, nil
}

// TODO(adf): actually do filtering
// TODO(adf): needs error handling
func (server *PaneServer) ListShares(share_filter *ShareFilter) (*ShareListResponse, error) {
  var rv *ShareListResponse

  server.stLock.RLock()

  rv = &ShareListResponse { Result: ResultSuccess.Enum() }
  for _, v := range server.shareTree {
    rv.ShareId = append(rv.ShareId, v.Id)
  }

  server.stLock.RUnlock()

  return rv, nil
}

func (server *PaneServer) ViewShare(share_id *ShareID) (*ShareResponse, error) {
  var rv *ShareResponse

  server.stLock.RLock()
  val, present := server.shareTree[share_id.GetIndex()]
  server.stLock.RUnlock()  // TODO(adf): use defer?

  if present {
    rv = &ShareResponse { Result: ResultSuccess.Enum(), Share: val }
  } else {
    rv = &ShareResponse { Result: ResultInvalidRequest.Enum() }
  }

  return rv, nil
}

/***************************************************************************************************
 *
 * Verb commands
 *
 **************************************************************************************************/

func (server *PaneServer) MakeRequest(request *Request) (*RequestResponse, error) {
  return nil, nil
}

func (server *PaneServer) ProvideHint(hint *Hint) (*GenericResponse, error) {
  return nil, nil
}


func (server *PaneServer) IssueQuery(query *Query) (*QueryResponse, error) {
  return nil, nil
}
