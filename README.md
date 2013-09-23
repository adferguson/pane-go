<pre>
     _______  ____     ____   ___ _______
    /  __  / / __ \    |   \  | | | _____|
   /  /_/ / / /__\ \   | |\ \ | | | |___
  / _____/ / ______ \  | | \ \| | |  __|
 / /      / /      \ \ | |  \ \ | | |____
/_/      /_/        \_\|_|   \__| |______|

    A PROTOTYPE PARTICIPATORY NETWORK

</pre>

PANE is a prototype OpenFlow controller which implements Participatory
Networking, an API for end-users, hosts and applications to take part in network
management. PANE allows these principals to directly contact the network
control-plane to  place requests for resources, provide hints about future
traffic, or query the state of the network. PANE divides and delegates authority
for network management using a hierarchy of "shares," which are also managed by
interacting with the PANE server.

Code Layout
-------------------------
`/src/pane/`       (Go source code)

`main.go`          (Execution entry point)

`pane.thrift`      (Thrift file defining the API)


Building PANE 
-------------------------

PANE requires:
  * Go compiler (tested with go 1.1.1)
  * Brown's Go-OpenFlow library (GoOF)
  * Go packages: go-thrift (which also provides `generator`)

Detailed instructions for building:
<pre>
go get github.com/samuel/go-thrift
make
</pre>

You will also need `$GOPATH` set somewhere sensible, and for `$GOPATH/bin` to be
in your `$PATH`.

Lastly, make sure GoOF is in your `$GOPATH`, or do `go get github.com/brownsys/goof`.
[Not actually working yet, I think.]

Optionallly, you may use Apache Thrift (tested with 0.9.0) and the Python bindings
for Thrift to run the simple Python client. The easiest way to compile pane.thrift
is: `thrift --gen py pane.thrift`.



Research
-------------------------
PANE is part of the [participatory networking project](http://pane.cs.brown.edu) at
[Brown University](http://www.cs.brown.edu).


Gofmt
-------------------------
This looks nice: `gofmt -tabwidth=2 -tabs=false`
