package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"k8s.io/klog/v2"
	"net/http"
	"time"
)

const (
	JsonContentType = "application/json"
)

type warpedResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func inputScanner() {
	_, err := fmt.Scanln(&buffer)
	if err != nil {
		klog.Fatalf("failed input scanner: %s", err.Error())
	}
}

func readWarpedResp[T any](resp *http.Response, err error) *warpedResponse[T] {
	if err != nil {
		klog.Errorf("failed to request: %s", err.Error())
		time.Sleep(time.Millisecond * 100)
		return nil
	}
	var result warpedResponse[T]
	if err = client.DecodeResponse(resp, &result); err != nil {
		klog.Errorf("failed to decode response: %s", err.Error())
		time.Sleep(time.Millisecond * 100)
		return nil
	}
	return &result
}

func readResp[T any](resp *http.Response, err error) *T {
	result := readWarpedResp[T](resp, err)
	if result == nil {
		return nil
	}
	if result.Code != 200 {
		klog.Errorf("service error: %d, %s", result.Code, result.Message)
		time.Sleep(time.Millisecond * 100)
		return nil
	}
	return &result.Data
}

func must2bytes(target any) []byte {
	b, err := json.Marshal(target)
	if err != nil {
		klog.Fatalf("must2bytes failed: %s", err.Error())
	}
	return b
}

func must2Reader(target any) io.Reader {
	return bytes.NewReader(must2bytes(target))
}
