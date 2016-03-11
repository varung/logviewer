# logviewer
helps visualize the logs of multiple programs that are interacting with each other

assume each line in the log file looks like
00:12:39.598400 stuff

We print out something like:

$ go run main.go a b
```
                    a                                                 b                                                 
00:12:39.598400     >RPC 10.0.0.1 AppendEntriesRequest
00:12:39.598450                                                       <RPC 10.0.0.0 AppendEntriesRequest 
00:12:39.598470                                                       >RPC 10.0.0.1 AppendEntriesResponse
00:12:39.598500     <RPC 10.0.0.0 AppendEntriesResponse
```
