# bitcoin-grpc
gRPC interface for bitcoind JSON-RPC

# Development

*Requirements*

* Install docker `sh scripts/install_docker.sh`
* Pull and run a  regtest `bitcoind` cluster `sh scripts/run_bitcoind_regtest.sh`
* Compile protobuf definitions `sh scripts/compile_proto.sh`

*Build & Run*

* Build the grpc server `go build`
* Run the server `./bitcoin-grpc`

* Build the grpc client `cd client && go build`
* Run the test client request `./client`


