// Thrift protocol for the Particpatory Networking API
// http://pane.cs.brown.edu
//
// Author: Andrew Ferguson (adf@cs.brown.edu)
//

namespace cpp pane
namespace java edu.brown.cs.pane

////////////////////////////////////////////////////////////////////////////////
// 
// Extensible wrappers for describing IP addresses, transport protocols, and
// transport ports
//
////////////////////////////////////////////////////////////////////////////////

// TODO(adf): IP prefixes and binary representations
enum IPAddrType {
  UNKNOWN = 0,
  IPV4_STRING = 1,
  IPV6_STRING = 2
}

struct IPAddress {
  1: optional IPAddrType type;  // required
  2: optional string address;
}

enum TransportProto {
  UNKNOWN = 0,
  TCP = 1,
  UDP = 2
}

// Leave extensible to support port ranges, port prefixes (already supported
// in Open vSwitch), etc.
struct Port {
  2: optional i32 number;
}

////////////////////////////////////////////////////////////////////////////////
// 
// Time Handling
//
// To express either a fixed time (seconds since Unix epoch), or a time
// relative to when the message arrives at the server.
// For example "now" is type:  TT_RELATIVE, time: 0
//
////////////////////////////////////////////////////////////////////////////////

enum TimeType {
  UNKNOWN = 1,
  RELATIVE = 2,
  ABSOLUTE = 3
}

struct Time {
  1: optional TimeType type,  // required
  2: optional i32 time  // required
}

////////////////////////////////////////////////////////////////////////////////
// 
// PANE Shares
//
// A share describes *who* can say *what* about *which* flows in the network
//   who - set of Principals
//   what - set of Privileges
//   which - set of Flows
//
////////////////////////////////////////////////////////////////////////////////

//////////////
// Principals
//////////////

struct Principal {
  1: optional string user;  // empty means all users
  2: optional string host;  // empty means all hosts
  3: optional string application;  // empty means all applications
}

//////////////
// Flows
//////////////

// With Flows, any field left empty is treated as a wildcard
struct Flow {
  1: optional IPAddress src_ip;
  2: optional IPAddress dst_ip;
  3: optional TransportProto transport_proto;
  4: optional Port src_port;
  5: optional Port dst_port;
}

//////////////
// Privileges: cover Requests, Hints, Queries
//////////////

enum RequestType {
  UNKNOWN = 0,
  ALLOW = 1,
  DENY = 2,
  RESERVE = 3,
  RATE_LIMIT = 4,
  WAYPOINT = 5,
  AVOID = 6
}

// Limits for Request privileges
struct RequestPrivilege {
  1: optional RequestType type;

  // can be used by all request types
  2: optional i32 time_limit;
  // can be used by RESERVE, RATE_LIMIT
  3: optional i32 bandwidth_limit;
  // can be used by WAYPOINT, RESERVE
  4: optional list<IPAddress> ip_limit;
}

enum HintType {
  UNKNOWN = 0,
  DURATION = 1
}

// Limits for Hint privileges
struct HintPrivilege {
  1: optional HintType type;
  // can be used by DURATION
  2: optional i32 time_limit;
}

enum QueryType {
  UNKNOWN = 0,
  TRAFFIC = 1
}

// Limits for Query privileges
struct QueryPrivilege {
  1: optional QueryType type;
  // can by used by TRAFFIC
  2: optional list<IPAddress> src_ip;
  3: optional list<IPAddress> dst_ip;
}

// TODO(adf): administrative privileges? (add user, create subshare, etc.)
enum PrivilegeType {
  UNKNOWN = 0,
  REQUEST = 1,
  HINT = 2,
  QUERY = 3
}

struct Privilege {
  1: optional PrivilegeType type;  // required

  2: optional RequestPrivilege request;
  3: optional HintPrivilege hint;
  4: optional QueryPrivilege query;
}

//////////////
// Share
//////////////

struct ShareID {
  1: optional string name;  // required
}

struct Share {
   1: optional ShareID id;  // required
   // the following reference Share.id
   2: optional ShareID parent;  // required, except for RootShare

  // skip to 10 to leave room for other metadata
  10: optional list<Principal> principal;  // empty means all principals
                                           // TODO(adf): sure that's a good idea?
  11: optional list<Flow>      flow;       // empty means all flows
  12: optional list<Privilege> privilege;  // empty means NO privileges
}

////////////////////////////////////////////////////////////////////////////////
// 
// PANE Verbs: Requests, Hints, Queries
//
////////////////////////////////////////////////////////////////////////////////

struct Request {
   1: optional RequestType type;  // required
   2: optional Principal principal;  // required
   3: optional list<Flow> flow;  // empty means all flows
   // references Share.id
   4: optional ShareID share;  // required
   5: optional bool strict; // default = true

   6: optional Time fromTime;
   7: optional Time untilTime;

  // required for RESERVE, RATE_LIMIT
  20: optional i32 bandwidth;
  // required for WAYPOINT, AVOID
  21: optional i32 ip;
}

struct Query {
   1: optional QueryType type;  // required

  // required for TRAFFIC
  20: optional IPAddress src_ip;
  21: optional IPAddress dst_ip;
}

struct Hint {
   1: optional HintType type;  // required

  // required for DURATION
  20: optional Time duration;
}

////////////////////////////////////////////////////////////////////////////////
// 
// RPC service stub declarations
//
////////////////////////////////////////////////////////////////////////////////

enum Result {
  UNKNOWN_RESULT = 0,
  SUCCESS = 1,
  ACCEPTED = 2,  // for hints
  INVALID_REQUEST = 20,
  INVALID_PERMISSION = 21,
  INSUFFICIENT_RESOURCES = 22,
  OTHER_FAILURE = 100
}

struct GenericResponse {
  1: optional Result result;  // required
  2: optional string msg;
}

struct Grant {
  // References Share.id
  1: optional ShareID share_id;  // required
  2: optional Principal to_principal;  // required
}

struct ShareFilter {
  1: optional Principal principal;
  2: optional Flow flow;
}

struct ShareListResponse {
  1: optional Result result;
  // references Share.id
  10: optional list<ShareID> share_id;
}

struct ShareResponse {
   1: optional Result result;  // required
  10: optional Share share;  // required
}

struct RequestResponse {
   1: optional Result result;  // required

  // For non-strict requests, return what was achieved
  10: optional i32 bandwidth;
  11: optional list<Flow> flow;
}

struct QueryResponse {
   1: optional Result result;  // required
  10: optional i32 traffic;  // for TRAFFIC
}

service PaneService {
  GenericResponse authenticate(1: Principal principal);

  GenericResponse grantShare(1: Grant grant);
  GenericResponse newShare(1: Share share);
// TODO(adf):  ... getSchedule (...);
  ShareListResponse listShares(1: ShareFilter share_filter);
  ShareResponse viewShare(1: ShareID share_id);

  RequestResponse makeRequest(1: Request request);
  QueryResponse issueQuery(1: Query query);
  GenericResponse provideHint(1: Hint hint);
}
