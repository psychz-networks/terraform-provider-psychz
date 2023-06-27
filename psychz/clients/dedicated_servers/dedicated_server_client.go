package psychz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.psychz.net/v1"

type Client struct {
	AccessToken    string
	AccessUsername string
}

func NewClient(accessToken, accessUsername string) *Client {
	return &Client{
		AccessToken:    accessToken,
		AccessUsername: accessUsername,
	}
}

type OrderExpress struct {
	PlanID                int
	OrderQuantity         int
	PaymentMode           int
	OSCategory            int
	OSID                  int
	DiskPartitionID       int
	SoftwareRAID          int
	Hostname              string
	Password              string
	PartnerID             int
	EnforcePasswordChange int
}

type OrderExpressResponse struct {
	Data OrderData `json:"data"`
}

type OrderData struct {
	Status               bool   `json:"status"`
	Message              string `json:"message"`
	Paid                 int    `json:"paid"`
	API                  string `json:"_api"`
	BrandID              string `json:"brand_id"`
	OrderID              string `json:"order_id"`
	Coupon               string `json:"coupon"`
	PlanID               string `json:"plan_id"`
	OrderQuantity        string `json:"order_quantity"`
	PaymentType          string `json:"payment_type"`
	CCInfo               string `json:"ccinfo"`
	InvoiceID            string `json:"invid"`
	InvoiceIDAlternative string `json:"inv_id"`
	ClientID             int    `json:"client_id"`
}

func (c *Client) CreateOrderExpress(ctx context.Context, data map[string]interface{}) (*OrderExpressResponse, error) {
	data["access_token"] = c.AccessToken
	data["access_username"] = c.AccessUsername
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/order_express", baseURL), bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("Psychz Networks API error: %s", string(bodyBytes))
	}

	var orderExpressResponse OrderExpressResponse

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bodyBytes, &orderExpressResponse); err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", orderExpressResponse)

	return &orderExpressResponse, nil
}

type OrderDetail struct {
	OrderID int `json:"data.order_id"`
	Data    Data   `json:"data"`
}

type Data struct {
	OrderTime    string `json:"order_time"`
	LastActivity string `json:"last_activity"`
	OrderStatus  string `json:"order_status"`
}

func (c *Client) GetOrderDetail(ctx context.Context, orderID int) (*OrderDetail, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/order_detail?order_id=%d&access_token=%s&access_username=%s", baseURL, orderID, c.AccessToken, c.AccessUsername), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("Psychz Networks API error: %s", string(bodyBytes))
	}

	var orderInfo OrderDetail

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bodyBytes, &orderInfo); err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", orderInfo)
	return &orderInfo, nil
}
