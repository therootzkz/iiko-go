package iiko

import "github.com/google/uuid"

type OrderCloseRequest struct {
	// OrganizationID
	// Can be obtained by /api/1/organizations operation.
	OrganizationID uuid.UUID `json:"organizationId"`
	// OrderID
	OrderID uuid.UUID `json:"orderId"`
}

type OrderCloseResponse struct {
	// Operation ID.
	CorrelationID uuid.UUID `json:"correlationId"`
}

// Close order.
//
// iiko API: /api/1/order/close
func (c *Client) OrderClose(req *OrderCloseRequest, opts ...Option) (*OrderCloseResponse, error) {
	var response OrderCloseResponse

	if err := c.post(true, "/api/1/order/close", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil

}
