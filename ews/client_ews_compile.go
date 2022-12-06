package ews

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Step 1: Compile Rust source code to WebAssembly
//curl -X POST https://ews-management.abp-monsters.com/compile?accountId=-1 -H "x-wasm-id: leaked-redirector" -H "Content-Type: application/zip" --data-binary "@./leaked-redirector/leaked-redirector.zip"
//
//Step 2: Deploy the compiled WebAssembly code to the path '/login' for a site hackathon.abp-monsters.com
//curl -X POST https://ews-management.abp-monsters.com/deploy?accountId=-1 -H "x-wasm-id: leaked-redirector" -H "x-filter-path: /login"

// Endpoints (unexported consts)
const endpointWASMCompile = "compile"

type WASMStruct struct {
	// Define as string for now
	Lambda string `json:"lambda"`
}

type ApiStatus struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// EwsApiDTO - Same DTO for: GET response, POST request, and POST response
type EwsApiDTO struct {
	Status []ApiStatus `json:"status"`
	Data   WASMStruct  `json:"data"`
}

func (c *Client) CompileWebAssembly(accountID, lambdaName, lambdaPath string) error {
	log.Printf("[INFO] Compile WebAssembly for accountID: %s\n", accountID)
	reqURL := fmt.Sprintf("%s/%s?accountId=%s", c.config.baseURLEWS, endpointWASMCompile, accountID)
	log.Printf("[INFO]  reqURL: %v\n", reqURL)

	//todo - Raphy - update logic here
	resp, err := c.DoFormDataRequestWithHeaders(http.MethodPost, reqURL, wasmJSON, contentTypeApplicationZip, lambdaName, "", false)

	if err != nil {
		return fmt.Errorf("Error executing Compile WebAssembly request for accountID %s: %s", accountID, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] EWS Compile WebAssembly JSON response: %s\n", string(responseBody))

	if string(responseBody) != "success" {
		return fmt.Errorf("failed to compile lambda")
	}

	return nil
}
