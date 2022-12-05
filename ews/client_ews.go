package ews

import (
	"encoding/json"
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
const endpointWASMDeploy = "deploy"

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

func (c *Client) CompileWebAssembly(accountID, lambdaName string, requestDTO EwsApiDTO) (*EwsApiDTO, error) {
	log.Printf("[INFO] Compile WebAssembly for accountID: %s\n", accountID)

	log.Printf("[INFO]  requestDTO: %+v\n", requestDTO)

	wasmJSON, err := json.Marshal(requestDTO)
	log.Printf("[INFO]  wasmJSON: %v\n", string(wasmJSON))
	reqURL := fmt.Sprintf("%s/%s?accountId=%s", c.config.baseURLEWS, endpointWASMCompile, accountID)
	log.Printf("[INFO]  reqURL: %v\n", reqURL)

	//resp, err := c.DoJsonRequestWithHeaders(http.MethodPost, reqURL, wasmJSON, CompileWASM)
	resp, err := c.DoFormDataRequestWithHeaders(http.MethodPost, reqURL, wasmJSON, contentTypeApplicationZip, lambdaName, "", false)

	if err != nil {
		return nil, fmt.Errorf("Error executing Compile WebAssembly request for accountID %s: %s", accountID, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] EWS Compile WebAssembly JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var responseDTO EwsApiDTO
	err = json.Unmarshal([]byte(responseBody), &responseDTO)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Compile WebAssembly JSON response for accountID %s: %s\nresponse: %s", accountID, err, string(responseBody))
	}

	return &responseDTO, nil
}

func (c *Client) DeployWebAssembly(accountID, lambdaName, filterPath string) (*EwsApiDTO, error) {
	log.Printf("[INFO] Deploy WebAssembly for accountID: %s\n", accountID)

	reqURL := fmt.Sprintf("%s/%s?accountId=%s", c.config.baseURLEWS, endpointWASMDeploy, accountID)

	//resp, err := c.PostFormWithHeaders(reqURL, values, DeployWASM)
	resp, err := c.DoFormDataRequestWithHeaders(http.MethodPost, reqURL, nil, contentTypeApplicationZip, lambdaName, filterPath, true)
	if err != nil {
		return nil, fmt.Errorf("Error executing Deploy WebAssembly request for accountID %s: %s", accountID, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] EWS Deploy WebAssembly JSON response: %s\n", string(responseBody))

	// Parse the JSON
	var responseDTO EwsApiDTO
	err = json.Unmarshal([]byte(responseBody), &responseDTO)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Deploy WebAssembly JSON response for accountID %s: %s\nresponse: %s", accountID, err, string(responseBody))
	}

	return &responseDTO, nil
}
