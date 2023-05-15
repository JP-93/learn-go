package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type HttpReportPCMPayload struct {
	FapiInteractionId string `json:"fapiInteractionId"`
	Endpoint          string `json:"endpoint"`
	StatusCode        int64  `json:"statusCode"`
	HttpMethod        string `json:"httpMethod"`
	ClientOrgId       string `json:"clientOrgId"`
	ServerOrgId       string `json:"serverOrgId"`
	CorrelationId     string `json:",omitempty"`
	Timestamp         string `json:"timestamp"`
	ProcessTimespan   int64  `json:"processTimespan"`
	ClientSSId        string `json:"clientSSId"`
	EndpointUriPrefix string `json:"endpointUriPrefix"`
	AdditionalInfo    `json:"additionalInfo"`
}

// AdditionalInfo represents aditional information used on pcm payload
type AdditionalInfo struct {
	ConsentId       string `json:"consentId"`
	PersonType      string `json:"personType"`
	LocalInstrument string `json:"localInstrument"`
	Status          string `json:"status"`
	RejectReason    string `json:"rejectReason"`
}

type Response struct {
	ReportID      string
	Status        string
	CorrelationId string
	Message       string
	HttpReportPCMPayload
}

func main() {
	var httpReport []HttpReportPCMPayload
	json.NewDecoder(strings.NewReader(j)).Decode(&httpReport)

	var arr []Response

	for _, res := range httpReport {
		ruuId := "qwerrtt"
		resp := Response{
			ReportID:      ruuId,
			Status:        "accepted",
			CorrelationId: res.CorrelationId,
		}
		arr = append(arr, resp)
	}

	ret, _ := json.Marshal(arr)
	b := bytes.NewBuffer(ret).String()

	fmt.Println(b)
}

var j = `[{
	"fapiInteractionId": "1",
	"endpoint": "AS",
	"statusCode": "D",
	"httpMethod": "F",
	"correlationId": "G",
	"additionalInfo": {
	"personType": "PF",
	"consentId": "urn:123",
	"status": "G",
	"rejectReason": "G"
	},
	"timestamp": "f",
	"processTimespan": 123,
	"clientSSId": "g",
	"clientOrgId": "h",
	"serverOrgId": "a",
	"endpointUriPrefix": "a",
	"role": "CLIENT"

},{
	"fapiInteractionId": "2",
	"endpoint": "12",
	"statusCode": "3",
	"httpMethod": "4",
	"correlationId": "d",
	"additionalInfo": {
	"personType": "PF",
	"consentId": "urn:123",
	"status": "g",
	"rejectReason": "j"
	},
	"timestamp": "f",
	"processTimespan": 123,
	"clientSSId": "n",
	"clientOrgId": "m",
	"serverOrgId": "n",
	"endpointUriPrefix": "f",
	"role": "CLIENT"
}]`
