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
const endpointWASMDeploy = "deploy"

func (c *Client) DeployWebAssembly(accountID, lambdaName, filterPath string) error {
	log.Printf("[INFO] Deploy WebAssembly for accountID: %s\n", accountID)
	reqURL := fmt.Sprintf("%s/%s?accountId=%s", c.config.baseURLEWS, endpointWASMDeploy, accountID)

	resp, err := c.DoFormDataRequestWithHeaders(http.MethodPost, reqURL, nil, contentTypeApplicationZip, lambdaName, filterPath, true)
	if err != nil {
		return fmt.Errorf("Error executing Deploy WebAssembly request for accountID %s: %s", accountID, err)
	}

	// Read the body
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	// Dump JSON
	log.Printf("[DEBUG] EWS Deploy WebAssembly JSON response: %s\n", string(responseBody))
	if string(responseBody) != "success" {
		return fmt.Errorf("failed to deploy lambda")
	}
	return nil
}
