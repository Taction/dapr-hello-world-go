/**
 * @Author: zhangchao
 * @Description:
 * @Date: 2021/3/11 9:09 下午
 */
package main

import (
	"context"
	"encoding/json"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	stateStoreName = `statestore`
	daprPort       = "3500"
)

var (
	port   string
	client dapr.Client
)

type resp struct {
	Response string `json:"response"`
	Success  bool   `json:"success"`
}

func init() {
	if port = os.Getenv("DAPR_GRPC_PORT"); len(port) == 0 {
		port = daprPort
	}
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := strings.ToLower(r.Method)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	result, ok := serve(path, method, body)
	res := resp{
		Response: result,
		Success:  ok,
	}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {
	var err error
	// create the client
	client, err = dapr.NewClientWithPort(port)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":8001", nil)
}

func getState(storeName string) (value string, metadata map[string]string, errStr string) {
	item, err := client.GetState(context.Background(), storeName, "order")
	if err != nil {
		return "", nil, fmt.Sprintf("failed to get state: %s", err)
	}
	if len(item.Value) > 0 {
		return string(item.Value), item.Metadata, ""
	} else {
		return "", nil, "order Not Found"
	}
}

func putState(storeName, orderID string) (errStr string) {
	err := client.SaveState(context.Background(), storeName, "order", []byte(orderID))
	if err != nil {
		return fmt.Sprintf("Failed to persist state: %v", err)
	} else {
		return ""
	}
}

func delState(storeName string) (errStr string) {
	err := client.DeleteState(context.Background(), storeName, "order")
	if err != nil {
		return fmt.Sprintf("Failed to delete state: %v", err)
	} else {
		return ""
	}
}
