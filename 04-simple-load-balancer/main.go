package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// creating the Server interface
type Server interface {
  Address() string
  IsAlive() bool
  Serve(rw http.ResponseWriter, r *http.Request)
}

// creating struct for Reverse Proxy
type simpleServer struct{
  addr  string
  proxy *httputil.ReverseProxy
}

// this returns a new instance of a server
func newSimpleServer(addr string) *simpleServer {
  serverURL, err := url.Parse(addr)
  // creating a function handle the error
  handleError(err)
  // returning the new instance
  return &simpleServer {
    addr:  addr,
    proxy: httputil.NewSingleHostReverseProxy(serverURL),
  }
}

// creating struct for LoadBalancer
type LoadBalancer struct {
  port            string
  roundRobinCount int
  servers         []Server
}

// creating the function to create a new instance of LoadBalancer
func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
  return &LoadBalancer {
    roundRobinCount: 0,
    port:            port,
    servers:         servers,
  }
}

func handleError(err error){
  if err != nil {
    fmt.Printf("The error is : %v\n", err)
    os.Exit(1)
  }
}



// method
func (s *simpleServer) Address() string { return s.addr }

// method
func (s *simpleServer) IsAlive() bool { return true   }




// method
func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request){
  // s is the simpleserver which is a proxy
  // serverHTTP comes from httputil package
  // it serves the same site but in dfferent proxies at the same time
  s.proxy.ServeHTTP(rw, req)
}

// Method using roundrobin to find next server
func (lb *LoadBalancer) getNextAvailableServer() Server{
  // check which of the servers is alive
  server := lb.servers[lb.roundRobinCount % len(lb.servers)]
  for !server.IsAlive(){
    lb.roundRobinCount++
    server = lb.servers[lb.roundRobinCount % len(lb.servers)]
  }
  lb.roundRobinCount++
  // return the available server
  return server
}

// method accessible by the Loadbalancer
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request){
  // gets the next available server
  targetServer := lb.getNextAvailableServer()
  // this will serve all the addresses that has been provided
  fmt.Printf("forwarding request to address %q\n", targetServer.Address())
  targetServer.Serve(rw, req)
}


func main() {
  // ceating slice of multiple servers, each representing a different address
  servers:= []Server{
    newSimpleServer("https://twitter.com/"),
    newSimpleServer("https://google.com/"),
    newSimpleServer("https://duckduckgo.com/"),
  }

  // creating a loadbalancer and sending all the servers to it
  lb := NewLoadBalancer("8000", servers)
  // this function will be used to redirect the requests
  handleRedirect := func(rw http.ResponseWriter, req *http.Request){
    // to access methods "lb." is being used
    lb.serveProxy(rw, req)
  }
  // handles when / is hit at 8000, calls the serveproxy
  http.HandleFunc("/", handleRedirect)
  
	fmt.Printf("Serving requests at localhost : %s\n", lb.port)
  http.ListenAndServe(":" + lb.port, nil)
}
