package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/miekg/dns"
	"github.com/parnurzeal/gorequest"
)

type BitcoinPriceResponse struct {
	Bitcoin struct {
		Usd float64 `json:"usd"`
	} `json:"bitcoin"`
}

func bitcoin(w dns.ResponseWriter, r *dns.Msg) {
	remoteAddr := w.RemoteAddr().String()
	log.Printf("Received request for bitcoin from %s\n", remoteAddr)

	message := new(dns.Msg)
	message.SetReply(r)

	var response BitcoinPriceResponse
	_, _, errs := gorequest.New().Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd").
		EndStruct(&response)
	if errs != nil {
		log.Println("Error fetching BTC price:", errs)
		return
	}

	btcPrice := strconv.FormatFloat(response.Bitcoin.Usd, 'f', 2, 64)
	message.Answer = append(message.Answer, &dns.TXT{
		Hdr: dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0},
		Txt: []string{fmt.Sprintf("BTC price: $%s", btcPrice)},
	})

	log.Println("Sending response for bitcoin:", btcPrice)

	w.WriteMsg(message)
}

func about(w dns.ResponseWriter, r *dns.Msg) {
	remoteAddr := w.RemoteAddr().String()
	log.Printf("Received request for about from %s\n", remoteAddr)
	
	message := new(dns.Msg)
	message.SetReply(r)

	var lines_of_text []string
	lines_of_text = append(lines_of_text, "Hi, I'm Alec.")
	lines_of_text = append(lines_of_text, "I'm a Cloud Engineer & Information Science student from NL.")
	lines_of_text = append(lines_of_text, "My fields of interest are Cloud & Security.")
	lines_of_text = append(lines_of_text, "Lets connect! linkedin.com/in/alecsiemerink")

	for _, line := range lines_of_text {
		message.Answer = append(message.Answer, &dns.TXT{
			Hdr: dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0},
			Txt: []string{line},
		},
		)
	}

	log.Println("Sending response for about")

	w.WriteMsg(message)
}

func main() {
	handler := dns.NewServeMux() 
	handler.HandleFunc("about.", about)
	handler.HandleFunc("bitcoin.", bitcoin)

	server := &dns.Server{
		Addr: ":5003", 
		Net: "udp", 
		Handler: handler,
	}

	log.Println("Starting server on port 5003")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}

	log.Println("Server shutdown")
}
