# Protocol Buffer Experiment

This is a playground for me to test some protocol buffer stuff :)

Links and stuff I used:  
https://protobuf.dev/
https://protobuf.dev/getting-started/gotutorial/
https://github.com/protocolbuffers/protobuf/releases/tag/v27.3
https://www.youtube.com/watch?v=qWN69yfRsVs

Helpful cmds:

```
export PATH=$PATH:/c/protoc/bin
```

```
protoc --go_out=. --go_opt=paths=source_relative model/book.proto
```

Can use hexdump to see how much leaner protobuf is than json. Thanks to this repo that allows me to use hexdump on Mingw :) 
https://github.com/wahern/hexdump

```
hexdump -C book.protobuf
00000000 08 01 12 16 54 68 65 20 52 6f 61 64 20 74 6f 20 |....The Road to |
00000010 57 69 67 61 6e 20 50 69 65 72 1a 11 08 01 12 0d |Wigan Pier......|
00000020 47 65 6f 72 67 65 20 4f 72 77 65 6c 6c |George Orwell|
```

vs

```
$ hexdump -C book.json
00000000  7b 22 49 64 22 3a 31 2c  22 54 69 74 6c 65 22 3a  |{"Id":1,"Title":|
00000010  22 54 68 65 20 52 6f 61  64 20 74 6f 20 57 69 67  |"The Road to Wig|
00000020  61 6e 20 50 69 65 72 22  2c 22 41 75 74 68 6f 72  |an Pier","Author|
00000030  73 22 3a 5b 7b 22 49 64  22 3a 31 2c 22 4e 61 6d  |s":[{"Id":1,"Nam|
00000040  65 22 3a 22 47 65 6f 72  67 65 20 4f 72 77 65 6c  |e":"George Orwel|
00000050  6c 22 7d 5d 7d   
```

Final result reading info from protobuf file:

```
$ go run main.go
book from protobuf file - ID: 1, Title: The Road to Wigan Pier, Authors: [Id:1 Name:"George Orwell"], Category: Novel
```
