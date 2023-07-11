# GO DNS Fun

This is a simple DNS server project written in Go. It uses the [miekg/dns](github.com/miekg/dns) library to handle the DNS protocol.
I wrote this to learn more about DNS and Go.

This consists of 2 parts; 
1. A DNS server that responds to queries to `about` with a TXT record containing some information about me, like a resume. 
2. A Simple webserver that serves a static page with simple documentation how to use this.

### How to use
1. Clone this repo
2. Run `go build` to build the binary
3. Run `go run .` to start the DNS server
This command starts a DNS server on port `5030`
4. To make a query locally, run `dig @localhost -p 5030 about`


