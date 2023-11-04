package clients

import (
    "time"
    "encoding/json"
    "net"
    "net/http"
    "io/ioutil"
    "fmt"
    "context"

	"github.com/Tempoolu/bookstore_user/discovery"
)

func Rpc(name string, method string, path string) []byte {
    ips, _ := discovery.Discovery(name)
    ret, _ := json.Marshal(ips)
    fmt.Printf("%s", string(ret))

    d := net.Dialer{
        Timeout:   30 * time.Second,
        KeepAlive: 30 * time.Second,
    }

    tr := &http.Transport{
        DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
            addr = ips[0]
            return d.DialContext(ctx, network, addr)
        },
    }

    webclient := &http.Client{Transport: tr}

    // Use NewRequest so we can change the UserAgent string in the header
    req, err := http.NewRequest(method, "http://fakehost.com"+path, nil)
    if err != nil {
        panic(err)
    }

    res, err := webclient.Do(req)
    if err != nil {
        panic(err)
    }

    defer res.Body.Close()

    content, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }
    return content
}
