// This file is auto-generated, don't edit it. Thanks.
/**
 * This is for OpenApi SDK
 */
package client

import (
	"errors"
	"net/http"

	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CommonClient interface {
	ProcessCommonRequest(request *requests.CommonRequest) (*responses.CommonResponse, error)
}

type Client struct {
	RegionId        string
	AccessKeyId     string
	AccessKeySecret string
	Domain          string
	Scheme          string
	PathPattern     string
	Method          string
}

func NewClientWithAccessKey(regionid, accessKeyId, accessKeySecret string) (*Client, error) {
	client := new(Client)
	if regionid == "" || accessKeyId == "" || accessKeySecret == "" {
		return nil, errors.New("regionid, accessKeyId or accessKeySecret can't be unset")
	}
	client.AccessKeyId = accessKeyId
	client.RegionId = regionid
	client.AccessKeySecret = accessKeySecret
	return client, nil
}

func (client *Client) ProcessCommonRequest(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
	path := defaultString(client.PathPattern, request.PathPattern, "/api/acs/openapi")
	request_ := tea.NewRequest()
	request_.Protocol = tea.String(defaultString(client.Scheme, request.Scheme, "HTTP"))
	request_.Method = tea.String(defaultString(client.Method, request.Method, "GET"))
	request_.Pathname = tea.String(path)
	request_.Query = map[string]*string{
		"Product":  tea.String(request.Product),
		"RegionId": tea.String(request.RegionId),
		"Action":   tea.String(request.ApiName),
		"Version":  tea.String(request.Version),
	}

	if len(request.QueryParams) > 0 {
		if request.QueryParams["IdToken"] != "" {
			request_.Query["IdToken"] = tea.String(request.QueryParams["IdToken"])
			delete(request.QueryParams, "IdToken")
		}

		if request.QueryParams["AliUid"] != "" {
			request_.Query["AliUid"] = tea.String(request.QueryParams["AliUid"])
			delete(request.QueryParams, "AliUid")
		}

		if request.QueryParams["RiskCode"] != "" {
			request_.Query["RiskCode"] = tea.String(request.QueryParams["RiskCode"])
			delete(request.QueryParams, "RiskCode")
		}

		if request.QueryParams["TraceId"] != "" {
			request_.Query["TraceId"] = tea.String(request.QueryParams["TraceId"])
			delete(request.QueryParams, "TraceId")
		}
	}
	request_.Query["Params"] = util.ToJSONString(request.QueryParams)

	request_.Headers = map[string]*string{
		"host":       tea.String(defaultString(client.Domain, request.Domain, "work-cn-hangzhou.aliyuncs.com")),
		"user-agent": client.getUserAgent(),
		//"x-acs-version": tea.String(request.Version),
		//"x-acs-action":  tea.String(request.ApiName),
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		request_.Body = tea.ToReader(request.Content)
	} else if !tea.BoolValue(util.IsUnset(request.FormParams)) {
		request_.Body = tea.ToReader(util.ToJSONString(request.FormParams))
	}
	request_.Headers["content-type"] = tea.String("application/json")
	request_.Query["SignatureMethod"] = tea.String("HMAC-SHA1")
	request_.Query["SignatureVersion"] = tea.String("1.0")
	request_.Query["AccessKeyId"] = tea.String(client.AccessKeyId)

	signedParam := make(map[string]*string)
	for k, v := range request.QueryParams {
		signedParam[k] = tea.String(v)
	}

	for k, v := range request.FormParams {
		signedParam[k] = tea.String(v)
	}

	request_.Query["Signature"] = openapiutil.GetRPCSignature(signedParam, request_.Method, tea.String(client.AccessKeySecret))

	response_, _err := tea.DoRequest(request_, nil)
	if _err != nil {
		return nil, _err
	}
	if tea.BoolValue(util.Is4xx(response_.StatusCode)) || tea.BoolValue(util.Is5xx(response_.StatusCode)) {
		_res, _err := util.ReadAsJSON(response_.Body)
		if _err != nil {
			return nil, _err
		}

		errObj := util.AssertAsMap(_res)
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    tea.ToString(errObj["code"]),
			"message": "code: " + tea.ToString(tea.IntValue(response_.StatusCode)) + ", " + tea.ToString(errObj["message"]) + " request id: " + tea.ToString(errObj["requestId"]),
			"data":    err,
		})
		return nil, _err
	}

	response = &responses.CommonResponse{
		BaseResponse: &responses.BaseResponse{},
	}

	header := http.Header{}
	for k, v := range response_.Headers {
		header.Add(k, tea.StringValue(v))
	}

	httpResponse := &http.Response{
		StatusCode: tea.IntValue(response_.StatusCode),
		Header:     header,
		Status:     tea.StringValue(response_.StatusMessage),
		Body:       response_.Body,
	}

	err = responses.Unmarshal(response, httpResponse, "JSON")
	if err != nil {
		return response, tea.NewSDKError(map[string]interface{}{
			"code":    "SDK.JsonUnmarshalError",
			"message": err.Error(),
		})
	}

	return response, nil
}

/**
 * Get user agent
 * @return user agent
 */
func (client *Client) getUserAgent() *string {
	userAgent := util.GetUserAgent(tea.String(""))
	return userAgent
}

/**
 * If inputValue is not null, return it or return defaultValue
 * @param inputValue  users input value
 * @param defaultValue default value
 * @return the final result
 */
func DefaultAny(inputValue interface{}, defaultValue interface{}) (_result interface{}) {
	if tea.BoolValue(util.IsUnset(inputValue)) {
		_result = defaultValue
		return _result
	}

	_result = inputValue
	return _result
}

func defaultString(a, b, defaultValue string) string {
	if a != "" {
		return a
	} else if b != "" {
		return b
	} else {
		return defaultValue
	}
}
