package main
import (
	"testing"
	"time"
	"net/http"
	"errors"
	"bytes"
	//"fmt"
	"strings"
	"io/ioutil"
	"github.com/emicklei/go-restful"
	"github.com/liuliuzi/datacollecter/cmd/app"
)


func waitForServerUp(serverURL string) error {
	for start := time.Now(); time.Since(start) < time.Minute; time.Sleep(1 * time.Second) {
		_, err := http.Get(serverURL + "/")
		if err == nil {
			return nil
		}
	}
	return errors.New("waiting for server timed out")
}

func TestServer(t *testing.T) {
	serverURL := "http://0.0.0.0:8060/metrics"
	go func() {
		if err := app.Run("0.0.0.0","8060","5s"); err != nil {
			t.Errorf("%v", err)
			}
	}()
	if err := waitForServerUp(serverURL); err != nil {
		t.Errorf("%v", err)
	}


	// Send a POST request for cpu.
	var jsonStr = []byte(`{"MetricValue": "1221","Timestamp":"`+time.Now().Format(time.RFC3339)+`"}`)
	req, err := http.NewRequest("POST", serverURL+"/cpu/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", restful.MIME_JSON)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("unexpected error in sending req: %v", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated{
		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
	}

	// Send a POST request for multi metrics .
	jsonStr = []byte(`{"MetricValue": "1221","Timestamp":"`+time.Now().Format(time.RFC3339)+`"}`)
	req, err = http.NewRequest("POST", serverURL+"/cpu/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", restful.MIME_JSON)

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("unexpected error in sending req: %v", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated{
		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
	}

	jsonStr = []byte(`{"MetricValue": "12.21","Timestamp":"`+time.Now().Format(time.RFC3339)+`"}`)
	req, err = http.NewRequest("POST", serverURL+"/mem/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", restful.MIME_JSON)

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("unexpected error in sending req: %v", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated{
		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
	}

	req, err = http.NewRequest("GET", serverURL, nil)
	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("unexpected error in sending req: %v", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated{
		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if !strings.Contains(string(body), "cpu 1221"){
		t.Errorf("not get expect metric value")
	}
	if !strings.Contains(string(body), "mem 12.21"){
		t.Errorf("not get expect metric value")
	}
	if !strings.Contains(string(body), "# TYPE cpu gauge"){
		t.Errorf("not get expect metric value")
	}
	if !strings.Contains(string(body), "# TYPE mem gauge"){
		t.Errorf("not get expect metric value")
	}


	// test metric live time .
	jsonStr = []byte(`{"MetricValue": "1221","Timestamp":"`+time.Now().Format(time.RFC3339)+`"}`)
	req, err = http.NewRequest("POST", serverURL+"/cpu/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", restful.MIME_JSON)

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("unexpected error in sending req: %v", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated{
		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
	}

	time.Sleep(6 * time.Second)

	req, err = http.NewRequest("GET", serverURL, nil)
	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("unexpected error in sending req: %v", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated{
		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if len(body)!=0{
		t.Errorf("metric value live time not effect")
	}
}