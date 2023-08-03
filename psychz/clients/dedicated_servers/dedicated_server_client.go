package psychz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

type OrderDetails struct {
	Status bool           `json:"status"`
	Data   OrderDataInner `json:"data"`
}

type OrderDataInner struct {
	OrderID        string         `json:"order_id"`
	Total          string         `json:"total"`
	OrderTime      string         `json:"order_time"`
	LastActivity   string         `json:"last_activity"`
	AddressContact AddressContact `json:"address_contact"`
	OrderInfo      OrderInfo      `json:"order_info"`
	InvoiceInfo    InvoiceInfo    `json:"invoice_info"`
	PaymentInfo    PaymentInfo    `json:"payment_info"`
	DeviceInfo     DeviceInfo     `json:"device_info"`
	IPAssignments  []IPAssignment `json:"ip_assignments"`
}

type AddressContact struct {
	First                string `json:"first"`
	Last                 string `json:"last"`
	Company              string `json:"company"`
	Address              string `json:"address"`
	Phone                string `json:"phone"`
	Fax                  string `json:"fax"`
	Email                string `json:"email"`
	ReferredBy           string `json:"referred_by"`
	InvoiceSendDate      string `json:"invoice_send_date"`
	InvoiceDueDateMethod string `json:"invoice_due_date_method"`
}
type OrderInfo struct {
	Service string  `json:"service"`
	Price   float64 `json:"price"`
	Setup   float64 `json:"setup"`
	Period  string  `json:"period"`
}

type InvoiceInfo struct {
	InvoiceID     string  `json:"invoice_id"`
	InvoiceDate   string  `json:"invoice_date"`
	InvoiceDue    string  `json:"invoice_due"`
	InvoiceAmount float64 `json:"invoice_amount"`
}

type PaymentInfo struct {
	PaymentType string `json:"payment_type"`
}

type DeviceInfo struct {
	DeviceID  string `json:"device_id"`
	ServiceID string `json:"service_id"`
}

func (c *Client) GetOrderDetail(ctx context.Context, orderID int) (OrderDetails, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/order_detail?order_id=%d&access_token=%s&access_username=%s", baseURL, orderID, c.AccessToken, c.AccessUsername), nil)
	if err != nil {
		return OrderDetails{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return OrderDetails{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return OrderDetails{}, fmt.Errorf("Psychz Networks API error: %s", string(bodyBytes))
	}

	var orderInfo OrderDetails

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return OrderDetails{}, err
	}

	if err := json.Unmarshal(bodyBytes, &orderInfo); err != nil {
		return OrderDetails{}, err
	}

	// Check if ip_assignments is not an empty array
	if len(orderInfo.Data.IPAssignments) == 0 {
		// If the ip_assignments is an empty array, set IPAssignments to nil
		orderInfo.Data.IPAssignments = nil
	} else {
		// If the ip_assignments is a JSON array, keep it as it is
	}

	fmt.Printf("%+v\n", orderInfo)

	if orderInfo.Data.DeviceInfo.ServiceID != "0" {
		serviceID, err := strconv.Atoi(orderInfo.Data.DeviceInfo.ServiceID)
		if err != nil {
			return OrderDetails{}, err
		}
		serviceDetail, err := c.GetServiceDetail(ctx, serviceID)
		if err != nil {
			return OrderDetails{}, err
		}

		orderInfo.Data.IPAssignments = serviceDetail
	}

	return orderInfo, nil
}

type ServiceDetail struct {
	Status bool             `json:"status"`
	Data   ServiceDataInner `json:"data"`
}

type ServiceDataInner struct {
	ServiceID        string                  `json:"service_id"`
	DesServ          string                  `json:"desserv"`
	Price            string                  `json:"price"`
	Period           string                  `json:"period"`
	PeriodCode       string                  `json:"period_code"`
	Discount         string                  `json:"discount"`
	Cost             string                  `json:"cost"`
	ParentPack       string                  `json:"parentpack"`
	Active           string                  `json:"active"`
	AutoBill         string                  `json:"auto_bill"`
	Quantity         string                  `json:"quantity"`
	Billed           string                  `json:"billed"`
	OrderID          string                  `json:"order_id"`
	ServiceStart     string                  `json:"service_start"`
	ServiceCreated   string                  `json:"service_created"`
	ServiceLastRenew string                  `json:"service_last_renew"`
	ServiceRenewDate string                  `json:"service_renew_date"`
	IPAssignments    map[string]IPAssignment `json:"-"`
	IPAssignmentsArr []IPAssignment          `json:"ip_assignments"`
}

type IPAssignment struct {
	Address           string `json:"address"`
	AssignDescription string `json:"assign_description"`
}

func (s *ServiceDataInner) UnmarshalJSON(data []byte) error {
	type TempServiceDataInner ServiceDataInner
	temp := struct {
		*TempServiceDataInner
		IPAssignments json.RawMessage `json:"ip_assignments"`
	}{
		TempServiceDataInner: (*TempServiceDataInner)(s),
	}

	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	if len(temp.IPAssignments) > 0 && temp.IPAssignments[0] == '[' {
		err = json.Unmarshal(temp.IPAssignments, &s.IPAssignmentsArr)
	} else {
		s.IPAssignments = make(map[string]IPAssignment)
		err = json.Unmarshal(temp.IPAssignments, &s.IPAssignments)
	}

	return err
}

func (c *Client) GetServiceDetail(ctx context.Context, orderID int) ([]IPAssignment, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/services_details?service_id=%d&access_token=%s&access_username=%s&ip_assignments=1", baseURL, orderID, c.AccessToken, c.AccessUsername), nil)
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
		return []IPAssignment{}, nil
	}

	var serviceDetail ServiceDetail

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bodyBytes, &serviceDetail); err != nil {
		return nil, err
	}

	if len(serviceDetail.Data.IPAssignmentsArr) > 0 {
		return serviceDetail.Data.IPAssignmentsArr, nil
	}

	ipAssignments := make([]IPAssignment, 0, len(serviceDetail.Data.IPAssignments))
	for _, v := range serviceDetail.Data.IPAssignments {
		ipAssignments = append(ipAssignments, v)
	}

	return ipAssignments, nil
}

type OrderPlan struct {
	CategoryName  string `json:"category_name"`
	CategoryID    int    `json:"category_id"`
	BillingType   string `json:"billing_type"`
	PlanName      string `json:"plan_name"`
	PlanID        string `json:"plan_id"`
	BasePrice     string `json:"base_price"`
	ResellerPrice string `json:"reseller_price"`
	StandardNote  string `json:"standard_note"`
	LocationName  string `json:"location_name"`
}

type OrderPlanType struct {
	// BackupServers               map[string]OrderPlan `json:"backup_servers"`
	DedicatedServers map[string]OrderPlan `json:"dedicated_servers"`
	// RemoteDdosProtectionServers map[string]OrderPlan `json:"remote_ddos_protection_servers"`
	// CdnServers                  map[string]OrderPlan `json:"cdn_servers"`
}

type OrderPlanResponse struct {
	Status bool          `json:"status"`
	Data   OrderPlanType `json:"data"`
}

func (c *Client) GetOrderPlans(ctx context.Context, params map[string]string) (OrderPlanResponse, error) {

	url := fmt.Sprintf("%s/order_plans", baseURL)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return OrderPlanResponse{}, err
	}

	params["access_token"] = c.AccessToken
	params["access_username"] = c.AccessUsername

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return OrderPlanResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return OrderPlanResponse{}, fmt.Errorf("Psychz Networks API error: %s", string(bodyBytes))
	}

	var orderPlansResponse OrderPlanResponse
	if err := json.NewDecoder(resp.Body).Decode(&orderPlansResponse); err != nil {
		return OrderPlanResponse{}, err
	}

	fmt.Printf("%+v\n", orderPlansResponse)

	return orderPlansResponse, nil
}
