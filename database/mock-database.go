package database

import (
	"hotel-check-in/entities"

	"github.com/google/uuid"
)

var MockHotelDatabase = []entities.Hotel{
	{
		ID:   uuid.NewString(),
		Name: "Ibis Hotel",
		Slug: "ibis-hotel",
	},
	{
		ID:   uuid.NewString(),
		Name: "Hotel Trivago",
		Slug: "hotel-trivago",
	},
	{
		ID:   uuid.NewString(),
		Name: "Hotel Chinzansyo Tokyo",
		Slug: "hotel-chinzansyo",
	},
	{
		ID:   uuid.NewString(),
		Name: "Ibis Hotel",
		Slug: "ibis-hotel",
	},
	{
		ID:   uuid.NewString(),
		Name: "Hotel Bardo",
		Slug: "hotel-bardo",
	},
}
