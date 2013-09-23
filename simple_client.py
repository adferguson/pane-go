#!/usr/bin/env python

import sys
sys.path.append('gen-py/')
 
from pane import PaneService 
from pane.ttypes import *
from pane.constants import *
 
from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol

try:
  # Make an object
  adfShare = Share(id=ShareID(name="adfShare"),
                   parent=ShareID(name="RootShare"))

  # Talk to a server via buffered TCP sockets, using a binary protocol
  transport = TSocket.TSocket("localhost", 4242)
  transport = TTransport.TBufferedTransport(transport)
  protocol = TBinaryProtocol.TBinaryProtocol(transport)

  client = PaneService.Client(protocol)
  transport.open()

  # print the shares, add a share, then print them again

  auth = client.authenticate(Principal(user="root"))
  print auth

  sharelist = client.listShares(auth.nonce, ShareFilter())
  print sharelist

  response = client.newShare(auth.nonce, adfShare)
  print response

  sharelist = client.listShares(auth.nonce, ShareFilter())
  print sharelist

  transport.close()

except Thrift.TException, tx:
    print "Exception: %s" % (tx.message)

