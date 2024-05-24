package seeder

import (
	"go-fiber/pkg/db"

	"gorm.io/gorm"
)

func AddressSeeder(g *gorm.DB) {
	var contacts []*db.Contact

	addresses := []*db.Address{
		{Street: "street 1", City: "Bandung", Province: "Jawa Barat", Country: "Indonesia", PostalCode: "41234"},
		{Street: "street 2", City: "Surabaya", Province: "Jawa Timur", Country: "Indonesia", PostalCode: "41235"},
	}

	g.Find(&contacts)

	for i, contact := range contacts {
		addresses[i].ContactId = contact.Id
	}

	g.Create(&addresses)
}
