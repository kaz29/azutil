// Package powerbiembedded implements the Azure ARM Powerbiembedded service API
// version 2016-01-29.
//
// Client to manage your Power BI Embedded workspace collections and retrieve
// workspaces.
package powerbiembedded

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Powerbiembedded
	DefaultBaseURI = "https://management.azure.com"
)

// ManagementClient is the base client for Powerbiembedded.
type ManagementClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the ManagementClient client.
func New(subscriptionID string) ManagementClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the ManagementClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return ManagementClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// GetAvailableOperations indicates which operations can be performed by the
// Power BI Resource Provider.
func (client ManagementClient) GetAvailableOperations() (result OperationList, err error) {
	req, err := client.GetAvailableOperationsPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "powerbiembedded.ManagementClient", "GetAvailableOperations", nil, "Failure preparing request")
	}

	resp, err := client.GetAvailableOperationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "powerbiembedded.ManagementClient", "GetAvailableOperations", resp, "Failure sending request")
	}

	result, err = client.GetAvailableOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.ManagementClient", "GetAvailableOperations", resp, "Failure responding to request")
	}

	return
}

// GetAvailableOperationsPreparer prepares the GetAvailableOperations request.
func (client ManagementClient) GetAvailableOperationsPreparer() (*http.Request, error) {
	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.PowerBI/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetAvailableOperationsSender sends the GetAvailableOperations request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetAvailableOperationsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetAvailableOperationsResponder handles the response to the GetAvailableOperations request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetAvailableOperationsResponder(resp *http.Response) (result OperationList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
