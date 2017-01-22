# Electron + Go + gRPC
This is just a proof of concept for Electron + Go + gRPC. You can also check out
a video walkthrough of it on [YouTube](https://www.youtube.com/watch?v=jl2M1SfsoWk).

## Dependencies

### Protobuf
1. Make sure you have Protobuf install. If on OS X, you should be okay with just
   doing `brew install protobuf`.

### Electron
1. `npm install` from the `client` folder and just run `electron main.js`.

###  Go
1. Make sure you have gRPC installed. Follow the instructions
   [here](https://github.com/grpc/grpc-go).

## Running the program

### Generating the Server
1. From the `demo` folder, run the following command. `protoc -I . demo.proto
   --go_out=plugins=grpc:.`. This will generate the necessary PB package you
   need.

### Generating the Client
1. I actually could not get the regular client generation working, so I just
   gave up on this and just had it dynamically generated at runtime.


### Tying them together
1. First run the server by doing `go run server/main.go`.
2. Then run the client doing `electron client/main.js`.
3. You're done.

## Common Issues
If you are running into compile time errors, you might have to rebuild the
dependencies. I suggest running `./client/rebuild.sh` and seeing if that fixes
your issues.

## Contact
Questions? Comments? Concerns? Hit me up on
[Twitter](https://twitter.com/kwuchu)
