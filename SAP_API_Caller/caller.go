package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-business-area-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetBusinessArea(businessArea, language, businessAreaName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "BusinessArea":
			func() {
				c.BusinessArea(businessArea)
				wg.Done()
			}()
		case "Text":
			func() {
				c.Text(language, businessAreaName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) BusinessArea(businessArea string) {
	businessAreaData, err := c.callBusinessAreaSrvAPIRequirementBusinessArea("A_BusinessArea", businessArea)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(businessAreaData)
	}

	textData, err := c.callToText(businessAreaData[0].ToText)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(textData)
	}
	return
}

func (c *SAPAPICaller) callBusinessAreaSrvAPIRequirementBusinessArea(api, businessArea string) ([]sap_api_output_formatter.BusinessArea, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESSAREA_SRV", api}, "/")
	param := c.getQueryWithBusinessArea(map[string]string{}, businessArea)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBusinessArea(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToText(url string) ([]sap_api_output_formatter.Text, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Text(language, businessAreaName string) {
	data, err := c.callBusinessAreaSrvAPIRequirementText("A_BusinessAreaText", language, businessAreaName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBusinessAreaSrvAPIRequirementText(api, language, businessAreaName string) ([]sap_api_output_formatter.Text, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESSAREA_SRV", api}, "/")

	param := c.getQueryWithText(map[string]string{}, language, businessAreaName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithBusinessArea(params map[string]string, businessArea string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BusinessArea eq '%s'", businessArea)
	return params
}

func (c *SAPAPICaller) getQueryWithText(params map[string]string, language, businessAreaName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Language eq '%s' and substringof('%s', BusinessAreaName)", language, businessAreaName)
	return params
}
