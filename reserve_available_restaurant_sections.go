package iiko

import (
	"github.com/google/uuid"
)

type ReserveAvailableRestaurantSectionsRequest struct {
	//Collection of terminal group ID.
	//Can be obtained by /api/1/terminal_groups operation.
	TerminalGroupIDs []uuid.UUID `json:"terminalGroupIds"`
}

type ReserveAvailableRestaurantSectionsResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// Restaurant sections.
	RestaurantSections []RestaurantSection `json:"restaurantSections"`
}

type RestaurantSection struct {
	// Restaurant section ID.
	ID uuid.UUID `json:"id"`
	// Terminal group ID.
	// Can be obtained by /api/1/terminal_groups operation.
	TerminalGroupID uuid.UUID `json:"terminalGroupId"`
	// Name
	Name string `json:"name"`
	// Tables
	Tables []RestaurantTable `json:"tables"`
}

type RestaurantTable struct {
	// Table ID.
	ID uuid.UUID `json:"id"`

	// Number of table.
	Number int `json:"number"`

	// Table name specified in the organization settings.
	Name string `json:"name"`

	// Seating capacity of the table.
	SeatingCapacity int `json:"seatingCapacity"`

	// Is table deleted.
	IsDeleted bool `json:"isDeleted"`
}

// iiko API: /api/1/reserve/available_restaurant_sections
func (c *Client) ReserveAvailableRestaurantSections(req *ReserveAvailableRestaurantSectionsRequest, opts ...Option) (*ReserveAvailableRestaurantSectionsResponse, error) {
	var response ReserveAvailableRestaurantSectionsResponse

	if err := c.post(true, "/api/1/reserve/available_restaurant_sections", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
