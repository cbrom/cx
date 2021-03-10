package playground

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prashantv/gostub"
)

func TestGetExampleFileList(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// mock
	stubs := gostub.Stub(&exampleNames, []string{"aa", "dd"})
	defer stubs.Reset()
	exampleNameList, _ := json.Marshal(exampleNames)
	{
		handler := http.HandlerFunc(GetExampleFileList)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if rr.Body.String() != string(exampleNameList) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), exampleNameList)
		}
	}
}

func TestGetExampleFileContent(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	mockSuccesBody := `{"examplename": "test.cx"}`
	req, err := http.NewRequest("POST", "/", strings.NewReader(mockSuccesBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	mockSuccesRespone := `package main; 

	func main() {
			str.print("Hello, world example1!");
	`
	stubs := gostub.Stub(&getExampleFileContent, func(name string) (string, error) {
		return mockSuccesRespone, nil
	})
	fmt.Println(getExampleFileContent(""))
	defer stubs.Reset()
	handler := http.HandlerFunc(GetExampleFileContent)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if rr.Body.String() != string(mockSuccesRespone) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), mockSuccesRespone)
	}

}
