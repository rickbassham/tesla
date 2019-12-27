package tesla

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type debugTransport struct {
	r http.RoundTripper
}

func (d *debugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("****REQUEST****\n%s\n", dump)
	resp, err := d.r.RoundTrip(req)
	dump, _ = httputil.DumpResponse(resp, true)
	fmt.Printf("****RESPONSE****\n%s\n****************\n\n", dump)
	return resp, err
}
