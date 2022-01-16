# gosocket

## About:
- A Simple Go Socket Framework 

## Installation:
```
$ go get -u github.com/golangpros/gosocket
```

## Tutorial:
```golang
import "github.com/golangpros/gosocket"

func main() {
 	myhost := "127.0.0.1:8080"

 	ss, err := gosocket.NewSocketService(myhost)
	if err != nil {
		return
	}

	ss.SetHeartBeat(5*time.Second, 30*time.Second)

	ss.RegMessageHandler(HandleMessage)
	ss.RegConnectHandler(HandleConnect)
	ss.RegDisconnectHandler(HandleDisconnect)

	ss.Serv()
}

```

## Thanks for contributors:
- [krishpranav](https://github.com/krishpranav)
- [NukeWilliams](https://github.com/NukeWilliams)