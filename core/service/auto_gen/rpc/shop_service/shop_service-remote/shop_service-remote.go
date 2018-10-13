// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

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
        "git.apache.org/thrift.git/lib/go/thrift"
	"go2o/core/service/auto_gen/rpc/ttype"
        "go2o/core/service/auto_gen/rpc/shop_service"
)

var _ = ttype.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  SStore GetStore(i32 venderId)")
  fmt.Fprintln(os.Stderr, "  SStore GetStoreById(i32 shopId)")
  fmt.Fprintln(os.Stderr, "  Result TurnShop(i32 shopId, bool on, string reason)")
  fmt.Fprintln(os.Stderr, "  Result OpenShop(i32 shopId, bool opening, string reason)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
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
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
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
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := shop_service.NewShopServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "GetStore":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetStore requires 1 args")
      flag.Usage()
    }
    tmp0, err12 := (strconv.Atoi(flag.Arg(1)))
    if err12 != nil {
      Usage()
      return
    }
    argvalue0 := int32(tmp0)
    value0 := argvalue0
    fmt.Print(client.GetStore(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetStoreById":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetStoreById requires 1 args")
      flag.Usage()
    }
    tmp0, err13 := (strconv.Atoi(flag.Arg(1)))
    if err13 != nil {
      Usage()
      return
    }
    argvalue0 := int32(tmp0)
    value0 := argvalue0
    fmt.Print(client.GetStoreById(context.Background(), value0))
    fmt.Print("\n")
    break
  case "TurnShop":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "TurnShop requires 3 args")
      flag.Usage()
    }
    tmp0, err14 := (strconv.Atoi(flag.Arg(1)))
    if err14 != nil {
      Usage()
      return
    }
    argvalue0 := int32(tmp0)
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.TurnShop(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "OpenShop":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "OpenShop requires 3 args")
      flag.Usage()
    }
    tmp0, err17 := (strconv.Atoi(flag.Arg(1)))
    if err17 != nil {
      Usage()
      return
    }
    argvalue0 := int32(tmp0)
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.OpenShop(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
