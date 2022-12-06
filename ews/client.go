package ews

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const contentTypeApplicationUrlEncoded = "application/x-www-form-urlencoded"
const contentTypeApplicationJson = "application/json"
const contentTypeApplicationZip = "application/zip"

const durationOfRetriesInSeconds = 30

// Client represents an internal client that brokers calls to the Ews API
type Client struct {
	config          *Config
	httpClient      *http.Client
	providerVersion string
}

// NewClient creates a new client with the provided configuration
func NewClient(config *Config) *Client {
	client := &http.Client{}

	return &Client{config: config, httpClient: client, providerVersion: "1.2.2"}
}

func (c *Client) CreateFormDataBody(bodyMap map[string]interface{}) ([]byte, string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, value := range bodyMap {
		switch value.(type) {
		case string:
			fw, err := writer.CreateFormField(key)
			if err != nil {
				log.Printf("failed to create %s formdata field", key)
			}
			_, err = io.Copy(fw, strings.NewReader(fmt.Sprintf("%v", value)))
			break
		case []byte:
			fw, err := writer.CreateFormFile(key, filepath.Base(key+".pfx"))
			if err != nil {
				log.Printf("failed to create %s formdata field", key)
			}
			fw.Write(value.([]byte))
			break
		default:
			//throw error
		}
	}
	writer.Close()

	return body.Bytes(), writer.FormDataContentType()
}

func (c *Client) PostFormWithHeaders(url string, data url.Values, operation string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("Error preparing request: %s", err)
	}

	SetHeaders(c, req, contentTypeApplicationUrlEncoded, operation, nil)
	return c.executeRequest(req)
}

func (c *Client) DoJsonRequestWithCustomHeaders(method string, url string, data []byte, headers map[string]string, operation string) (*http.Response, error) {
	req, err := PrepareJsonRequest(method, url, data)
	if err != nil {
		return nil, fmt.Errorf("Error preparing request: %s", err)
	}

	SetHeaders(c, req, contentTypeApplicationJson, operation, headers)

	return c.executeRequest(req)
}

func (c *Client) DoJsonRequestWithHeaders(method string, url string, data []byte, operation string) (*http.Response, error) {
	return c.DoJsonRequestWithCustomHeaders(method, url, data, nil, operation)
}

func (c *Client) DoJsonAndQueryParamsRequestWithHeaders(method string, url string, data []byte, params map[string]string, operation string) (*http.Response, error) {
	req, err := PrepareJsonRequest(method, url, data)
	if err != nil {
		return nil, fmt.Errorf("Error preparing request: %s", err)
	}
	q := req.URL.Query()
	for name, value := range params {
		q.Add(name, value)
	}
	req.URL.RawQuery = q.Encode()

	SetHeaders(c, req, contentTypeApplicationJson, operation, nil)

	return c.executeRequest(req)
}

// GetRequestParamsWithCaid Use this function if you want to add caid to your request as a query param.
// you need to send caid if you want to preform action on resources belong to child account (example: reseller -> account)
func GetRequestParamsWithCaid(accountId int) map[string]string {
	var params = map[string]string{}
	if accountId != 0 {
		params["caid"] = strconv.Itoa(accountId)
	}

	return params
}

func (c *Client) DoFormDataRequestWithHeaders(method, url string, data []byte, contentType, wasmId, filterPath string, deploy bool) (*http.Response, error) {
	req, err := PrepareJsonRequest(method, url, data)
	if err != nil {
		return nil, fmt.Errorf("Error preparing request: %s", err)
	}

	//SetHeaders(c, req, contentType, operation, nil)
	SetHeadersNew(c, req, contentType, wasmId, filterPath, deploy)
	log.Printf("[DEBUG] DoFormDataRequestWithHeaders request: %v\n", *req)
	return c.executeRequest(req)
}

func (c *Client) sendBinaryRequest(method, url, lambdaPath, wasmId, filterPath string, deploy bool) (*http.Response, error) {

	wasmZipFile, err := ioutil.ReadFile(lambdaPath)
	if err != nil {
		panic(err)
	}

	payload := strings.NewReader(string(wasmZipFile))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, fmt.Errorf("Error preparing request: %s", err)
	}

	SetHeadersNew(c, req, contentTypeApplicationZip, wasmId, filterPath, deploy)

	responseOnRequest, errorOnRequest := client.Do(req)
	if errorOnRequest != nil {
		return nil, fmt.Errorf("Error on request: %s", errorOnRequest)
	}

	return responseOnRequest, errorOnRequest
}

func PrepareJsonRequest(method string, url string, data []byte) (*http.Request, error) {
	if data == nil {
		return http.NewRequest(method, url, nil)
	}

	return http.NewRequest(method, url, bytes.NewReader(data))
}

func SetHeadersNew(c *Client, req *http.Request, contentType, wasmId, filterPath string, deploy bool) {
	req.Header.Set("x-wasm-id", wasmId)
	if deploy {
		// Deploy case, we need filter-path
		req.Header.Set("x-filter-path", filterPath)
	} else {
		// Compile case, we need Content-Type (ZIP)
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("x-api-id", c.config.APIID)
	req.Header.Set("x-api-key", c.config.APIKey)
}

func SetHeaders(c *Client, req *http.Request, contentType string, operation string, customHeaders map[string]string) {
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("x-api-id", c.config.APIID)
	req.Header.Set("x-api-key", c.config.APIKey)
	req.Header.Set("x-tf-provider-ver", c.providerVersion)
	req.Header.Set("x-tf-operation", operation)

	if customHeaders != nil {
		for name, value := range customHeaders {
			req.Header.Set(name, value)
		}
	}
}

func (c *Client) executeRequest(req *http.Request) (*http.Response, error) {
	//if "read" action then we want to allow retries in case of timeout from ews service
	operation := req.Header.Get("x-tf-operation")
	if req.Method == http.MethodGet || (req.Method == http.MethodPost && strings.HasPrefix(strings.ToLower(operation), "read")) {
		var responseOnRequest *http.Response
		var errorOnRequest error
		resource.Retry(durationOfRetriesInSeconds*time.Second, func() *resource.RetryError {
			responseOnRequest, errorOnRequest = c.httpClient.Do(req)
			if errorOnRequest != nil {
				log.Printf("[ERROR] Error from Ews service when reading resource")
				return resource.NonRetryableError(errorOnRequest)
			}
			if responseOnRequest.StatusCode == 502 {
				log.Printf("[WARN] Error from Ews service when reading resource, performing retry")
				return resource.RetryableError(fmt.Errorf("error code 502 from ews service when reading resource, performing retry"))
			}
			return nil
		})
		return responseOnRequest, errorOnRequest
	}
	//if not a "read" request  - don't do retries (retires for updates are risky and result could be non-deterministic)
	return c.httpClient.Do(req)
}
