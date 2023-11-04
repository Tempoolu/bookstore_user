package discovery

import (
    "time"
    "fmt"
    "context"
    "strings"

    "go.etcd.io/etcd/clientv3"
)

var (
	client *clientv3.Client
)

func InitDiscovery(addr string) error {
    var err error
    client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return err
	}
    return nil
}

func Register(name string, addr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    _, err := client.Put(ctx, name + "_" + addr, "1")
	cancel()
	if err != nil {
		return err
	}
	return nil
}

func Unregister(name string, addr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    _, err := client.Delete(ctx, name + "_" + addr)
	cancel()
	if err != nil {
		return err
	}
	return nil
}

func Discovery(name string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    list, err := client.Get(ctx, name + "_", clientv3.WithPrefix())
	cancel()
	if err != nil {
		return nil, err
	}
	var resp []string
    for _, kv := range list.Kvs {
        resp = append(resp, strings.TrimPrefix(string(kv.Key), name+"_"))
    }
	return resp, nil
}
