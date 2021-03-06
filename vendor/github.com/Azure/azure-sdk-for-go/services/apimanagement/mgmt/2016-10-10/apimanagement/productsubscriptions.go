package apimanagement

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ProductSubscriptionsClient is the apiManagement Client
type ProductSubscriptionsClient struct {
	BaseClient
}

// NewProductSubscriptionsClient creates an instance of the ProductSubscriptionsClient client.
func NewProductSubscriptionsClient(subscriptionID string) ProductSubscriptionsClient {
	return NewProductSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProductSubscriptionsClientWithBaseURI creates an instance of the ProductSubscriptionsClient client.
func NewProductSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) ProductSubscriptionsClient {
	return ProductSubscriptionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByProducts lists the collection of subscriptions to the specified product.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// productID - product identifier. Must be unique in the current API Management service instance.
// filter - | Field        | Supported operators    | Supported functions                         |
// |--------------|------------------------|---------------------------------------------|
// | id           | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | name         | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | stateComment | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | userId       | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | productId    | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | state        | eq                     |                                             |
// top - number of records to return.
// skip - number of records to skip.
func (client ProductSubscriptionsClient) ListByProducts(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result SubscriptionCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductSubscriptionsClient.ListByProducts")
		defer func() {
			sc := -1
			if result.sc.Response.Response != nil {
				sc = result.sc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: productID,
			Constraints: []validation.Constraint{{Target: "productID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "productID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "productID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.ProductSubscriptionsClient", "ListByProducts", err.Error())
	}

	result.fn = client.listByProductsNextResults
	req, err := client.ListByProductsPreparer(ctx, resourceGroupName, serviceName, productID, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductSubscriptionsClient", "ListByProducts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByProductsSender(req)
	if err != nil {
		result.sc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ProductSubscriptionsClient", "ListByProducts", resp, "Failure sending request")
		return
	}

	result.sc, err = client.ListByProductsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductSubscriptionsClient", "ListByProducts", resp, "Failure responding to request")
	}

	return
}

// ListByProductsPreparer prepares the ListByProducts request.
func (client ProductSubscriptionsClient) ListByProductsPreparer(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"productId":         autorest.Encode("path", productID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/subscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByProductsSender sends the ListByProducts request. The method will close the
// http.Response Body if it receives an error.
func (client ProductSubscriptionsClient) ListByProductsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByProductsResponder handles the response to the ListByProducts request. The method always
// closes the http.Response Body.
func (client ProductSubscriptionsClient) ListByProductsResponder(resp *http.Response) (result SubscriptionCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByProductsNextResults retrieves the next set of results, if any.
func (client ProductSubscriptionsClient) listByProductsNextResults(ctx context.Context, lastResults SubscriptionCollection) (result SubscriptionCollection, err error) {
	req, err := lastResults.subscriptionCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductSubscriptionsClient", "listByProductsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByProductsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductSubscriptionsClient", "listByProductsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByProductsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductSubscriptionsClient", "listByProductsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByProductsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProductSubscriptionsClient) ListByProductsComplete(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result SubscriptionCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductSubscriptionsClient.ListByProducts")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByProducts(ctx, resourceGroupName, serviceName, productID, filter, top, skip)
	return
}
