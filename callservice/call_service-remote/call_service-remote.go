// Code generated by Thrift Compiler (0.18.0). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"talkservice"
	"callservice"
)

var _ = talkservice.GoUnusedProtection__
var _ = callservice.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  CallRoute acquireCallRoute(string to, FreeCallType callType,  fromEnvInfo)")
  fmt.Fprintln(os.Stderr, "  GroupCallRoute acquireGroupCallRoute(string chatMid, GroupCallMediaType mediaType, bool isInitialHost,  capabilities)")
  fmt.Fprintln(os.Stderr, "  GroupCall getGroupCall(string chatMid)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  var cfg *thrift.TConfiguration = nil
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans = thrift.NewTSocketConf(net.JoinHostPort(host, portStr), cfg)
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransportConf(trans, cfg)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := callservice.NewCallServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "acquireCallRoute":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "AcquireCallRoute requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err := (strconv.Atoi(flag.Arg(2)))
    if err != nil {
      Usage()
     return
    }
    argvalue1 := callservice.FreeCallType(tmp1)
    value1 := argvalue1
    arg28 := flag.Arg(3)
    mbTrans29 := thrift.NewTMemoryBufferLen(len(arg28))
    defer mbTrans29.Close()
    _, err30 := mbTrans29.WriteString(arg28)
    if err30 != nil { 
      Usage()
      return
    }
    factory31 := thrift.NewTJSONProtocolFactory()
    jsProt32 := factory31.GetProtocol(mbTrans29)
    containerStruct2 := callservice.NewCallServiceAcquireCallRouteArgs()
    err33 := containerStruct2.ReadField3(context.Background(), jsProt32)
    if err33 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.FromEnvInfo
    value2 := argvalue2
    fmt.Print(client.AcquireCallRoute(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "acquireGroupCallRoute":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "AcquireGroupCallRoute requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err := (strconv.Atoi(flag.Arg(2)))
    if err != nil {
      Usage()
     return
    }
    argvalue1 := callservice.GroupCallMediaType(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3) == "true"
    value2 := argvalue2
    arg36 := flag.Arg(4)
    mbTrans37 := thrift.NewTMemoryBufferLen(len(arg36))
    defer mbTrans37.Close()
    _, err38 := mbTrans37.WriteString(arg36)
    if err38 != nil { 
      Usage()
      return
    }
    factory39 := thrift.NewTJSONProtocolFactory()
    jsProt40 := factory39.GetProtocol(mbTrans37)
    containerStruct3 := callservice.NewCallServiceAcquireGroupCallRouteArgs()
    err41 := containerStruct3.ReadField4(context.Background(), jsProt40)
    if err41 != nil {
      Usage()
      return
    }
    argvalue3 := containerStruct3.Capabilities
    value3 := argvalue3
    fmt.Print(client.AcquireGroupCallRoute(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "getGroupCall":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetGroupCall requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetGroupCall(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
