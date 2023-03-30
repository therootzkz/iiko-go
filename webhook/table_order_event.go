package webhook

import (
	"github.com/google/uuid"
	"github.com/iiko-go/iiko-go"
	"time"
)

type TableOrderEventRequest struct {
	Notification
	EventInfo TableOrderEventInfo `json:"eventInfo"`
}

type TableOrderEventInfo struct {
	OrderID        uuid.UUID                 `json:"id"`
	PosID          string                    `json:"posId,omitempty"`
	ExternalNumber string                    `json:"externalNumber,omitempty"`
	OrganizationID uuid.UUID                 `json:"organizationId"`
	Timestamp      int64                     `json:"timestamp"`
	CreationStatus OrderCreationStatus       `json:"creationStatus"`
	ErrorInfo      *TableOrderEventErrorInfo `json:"errorInfo,omitempty"`
	Order          *OrderCreationDetails     `json:"order,omitempty"`
}

type OrderCreationDetails struct {
	// TableIDs.
	// Can be obtained by /api/1/reserve/available_restaurant_sections operation
	TableIDs        []uuid.UUID        `json:"tableIds"`
	Guest           *iiko.Customer     `json:"guest,omitempty"`
	Phone           string             `json:"phone,omitempty"`
	Status          iiko.OrderStatus   `json:"status"`
	WhenCreated     *time.Time         `json:"whenCreated,omitempty"`
	Waiter          *iiko.Waiter       `json:"waiter,omitempty"`
	TabName         *string            `json:"tabName,omitempty"`
	Sum             float64            `json:"sum"`
	Number          int                `json:"number"`
	SourceKey       *string            `json:"sourceKey,omitempty"`
	WhenBillPrinted *time.Time         `json:"whenBillPrinted,omitempty"`
	WhenClosed      *time.Time         `json:"whenClosed,omitempty"`
	GuestsInfo      *iiko.Guests       `json:"guestsInfo"`
	Items           []CreatedOrderItem `json:"items"`
	OrderType       CreatedOrderType   `json:"orderType"`
	TerminalGroupID uuid.UUID          `json:"terminalGroupId"`
}

type CreatedOrderType struct {
	// Order type ID.
	ID uuid.UUID `json:"id"`
	// Order type name.
	Name string `json:"name"`
	// Order type.
	OrderServiceType OrderServiceType `json:"orderServiceType"`
}

type CreatedOrderItem struct {
	// Item.
	Product CreatedOrderProduct `json:"product"`
	// Price per item unit. Can be sent different from the price in the base menu.
	Price float64 `json:"price"`
	// Total cost per item without tax, discounts/surcharges.
	Cost float64 `json:"cost"`
	// Whether price is predefined.
	PricePredefined bool `json:"pricePredefined"`
	// Unique identifier of the item in the order and for the whole system.
	PositionID uuid.UUID `json:"positionId,omitempty"`
	// Tax rate
	TaxPercent *float64             `json:"taxPercent,omitempty"`
	Type       CreatedOrderItemType `json:"type"`
	// Item cooking status.
	Status CookingStatus `json:"status"`
	// Quantity
	Amount float64 `json:"amount"`
	// Comment
	Comment string `json:"comment"`
}

type CreatedOrderProduct struct {
	// Product ID.
	ID uuid.UUID `json:"id"`
	// Product name.
	Name string `json:"name"`
}

type TableOrderEventErrorInfo struct {
	Code iiko.ErrorInfoCode `json:"code"`
	// Nonlocalized message.
	Message string `json:"message,omitempty"`
	// Localized message.
	Description string `json:"description,omitempty"`
	// Additional information.
	AdditionalData string `json:"additionalData,omitempty"`
}

type OrderServiceType string

const (
	OrderServiceTypeDeliveryByClient  OrderServiceType = "DeliveryByClient"
	OrderServiceTypeDeliveryByCourier OrderServiceType = "DeliveryByCourier"
	OrderServiceTypeCommon            OrderServiceType = "Common"
)

type CookingStatus string

const (
	CookingStatusCreated           CookingStatus = "Added"
	CookingStatusPrintedNotCooking CookingStatus = "PrintedNotCooking"
	CookingStatusStarted           CookingStatus = "CookingStarted"
	CookingStatusCookingCompleted  CookingStatus = "CookingCompleted"
	CookingStatusServed            CookingStatus = "Served"
)

type CreatedOrderItemType string

const (
	CreatedOrderItemTypeProduct  CreatedOrderItemType = "Product"
	CreatedOrderItemTypeCompound CreatedOrderItemType = "Compound"
	CreatedOrderItemTypeService  CreatedOrderItemType = "Service"
)

type OrderCreationStatus string

const (
	OrderCreationStatusSuccess    OrderCreationStatus = "Success"
	OrderCreationStatusError      OrderCreationStatus = "Error"
	OrderCreationStatusInProgress OrderCreationStatus = "InProgress"
)
