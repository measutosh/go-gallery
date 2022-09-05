
![](4_Simple_Load_Balancer_SS.gif)


# Simple Load Balancer

This project is a Simple Load Balancer that has 3 servers to balance the load.


## Little More
- Uses the given pacakges :- "fmt", "net/http", "net/http/httputil", "net/url", "os"
- It creates simpleServer, LoadBalancer, reverseProxy using the struct
- The loadbalancer checks if one of the servers are alive or not.
- If alive then redirects the address to that alive server or else keep searching among the 3 servers until it finds a suitable one.


