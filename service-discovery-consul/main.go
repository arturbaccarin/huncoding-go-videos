// https://youtu.be/xkYdLB9UiMo

package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/hashicorp/consul/api"
)

const (
	serviceName = "user-service"
	servicePort = 8080
)

func main() {
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	registerService(consulClient, serviceName, servicePort)

	defer deregisterService(consulClient, serviceName)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})

	go func() {
		http.ListenAndServe(fmt.Sprintf(":%d", servicePort), nil)
	}()

	key := "huncoding"
	value := "Hello World"

	if err := putKey(consulClient, key, value); err != nil {
		panic(err)
	}

	v, err := getKey(consulClient, key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value: %s\n", v)
}

func putKey(client *api.Client, key string, value string) error {
	kv := client.KV()
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	return err
}

func getKey(client *api.Client, key string) (string, error) {
	kv := client.KV()
	p, _, err := kv.Get(key, nil)
	if err != nil {
		return "", err
	}

	return string(p.Value), nil
}

func registerService(client *api.Client, serviceName string, servicePort int) error {
	registration := &api.AgentServiceRegistration{
		Name: serviceName,
		ID:   serviceName,
		Port: servicePort,
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/health", getOutboundIP(), servicePort),
			Interval: "10s",
		},
	}

	return client.Agent().ServiceRegister(registration)
}

func deregisterService(client *api.Client, serviceName string) error {
	return client.Agent().ServiceDeregister(serviceName)
}

func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
