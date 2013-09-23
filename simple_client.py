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

# Make an object
adfShare = Share(id=ShareID(name="adfShare"),
                 parent=ShareID(name="RootShare"))

# Talk to a server via TCP sockets, using a binary protocol
transport = TSocket.TSocket("localhost", 4242)
transport.open()
protocol = TBinaryProtocol.TBinaryProtocol(transport)

# print the shares, add a share, then print them again

service = PaneService.Client(protocol)

sharelist = service.listShares(ShareFilter())
print sharelist

response = service.newShare(adfShare)
print response

sharelist = service.listShares(ShareFilter())
print sharelist

