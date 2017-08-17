# logx
 - log tool,easy to use,high performance,handy,availability
 - version v1.1
# tps
```
 //OutPut to file
 //tps is : 1400000/s on windows
 //cpu i5-7600 3.5GHZ
 //8GB
 //it will be better on better platform
```
# simple to use
## install

```
go get -u github.com/kafrax/logx
```
## start
```
package main

import (
    "github.com/kafrax/logx"
)

func main(){
    logx.Debugf("test |message=%s", "logx is a lightweight log to use")
    logx.Infof("test |message=%s",  "logx is a lightweight log to use")
    logx.Errorf("test |message=%s", "logx is a lightweight log to use")
    logx.Warnf("test |message=%s",  "logx is a lightweight log to use")
    logx.Fatalf("test |message=%s", "logx is a lightweight log to use")
}
```
```
[DEBU][08-11.16.48.27.297][main.go|main.main|49] message=test |message=logx is a lightweight log to use
[INFO][08-11.16.48.27.297][main.go|main.main|50] message=test |message=logx is a lightweight log to use
[ERRO][08-11.16.48.27.297][main.go|main.main|51] message=test |message=logx is a lightweight log to use
[WARN][08-11.16.48.27.297][main.go|main.main|52] message=test |message=logx is a lightweight log to use
[FTAL][08-11.16.48.27.297][main.go|main.main|53] message=test |message=logx is a lightweight log to use

```

#  write to file
## config logx.json or config.json
- let logx.json  or config.json in your project root dir.
- will be executed by default , there is no config.json or logx.json yet.
- *notice* fileWriter use memory cache ,so must have enough time to do poller to save data to log file.
```
{
    "llevel":1,        //log level,1debug,2info,3warn,4error,5fatal
    "lmaxsize":256     //256mb
    "lout":"stdout",   //file|stdout
    "lbucketlen":1024, //memory cache size
    "lfilename":"logx",//log file name eg. logx2006-01-02.04.05.000.log
    "lfilepath":"./",  //log file path
    "lpollerinterval": //500 millisecond flush once
}
```
## start
```
package main

import (
    "github.com/kafrax/logx"
)

func main(){
    logx.Debugf("module=test |message=%s","logx is a lightweight log to use")
    var str string
    fmt.Scan(&str)
}
```

# future
 - data queue send to kafaka

# @me
 - kafrax.go@gmail.com