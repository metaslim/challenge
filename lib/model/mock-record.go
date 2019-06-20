package model

import (
	"github.com/metaslim/challenge/lib/schema"
)

var MockOrganizations = Organizations{
	BaseRecords: BaseRecords{
		Size: 2,
	},
	Items: []schema.Organization{
		schema.Organization{
			ID:         1,
			URL:        "http://localhost/organization/1",
			ExternalID: "external-id-1",
			Name:       "Shield",
			DomainNames: []string{
				"shield.gov.us",
			},
			CreatedAt:     "2016-05-21T11:10:28 -10:00",
			Details:       "Shield",
			SharedTickets: true,
			Tags: []string{
				"shield",
			},
		},
		schema.Organization{
			ID:         2,
			URL:        "http://localhost/organization/2",
			ExternalID: "external-id-2",
			Name:       "Avengers",
			DomainNames: []string{
				"avengers.world",
			},
			CreatedAt:     "2016-05-21T11:10:28 -10:00",
			Details:       "Avengers",
			SharedTickets: true,
			Tags: []string{
				"avengers",
			},
		},
	},
}

var MockTickets = Tickets{
	BaseRecords: BaseRecords{
		Size: 2,
	},
	Items: []schema.Ticket{
		schema.Ticket{
			ID:             "1",
			URL:            "http://localhost/ticket/1",
			ExternalID:     "external-id-1",
			CreatedAt:      "2016-05-21T11:10:28 -10:00",
			TicketType:     "incident",
			Subject:        "Thanos attack, Infinity Wars",
			Description:    "",
			Priority:       "high",
			Status:         "pending",
			SubmitterID:    1,
			AssigneeID:     2,
			OrganizationID: 1,
			Tags: []string{
				"thanos",
			},
			HasIncidents: true,
			DueAt:        "2019-05-21T11:10:28 -10:00",
			Via:          "web",
		},
		schema.Ticket{
			ID:             "2",
			URL:            "http://localhost/ticket/2",
			ExternalID:     "external-id-2",
			CreatedAt:      "2016-05-21T11:10:28 -10:00",
			TicketType:     "incident",
			Subject:        "Loki attack",
			Description:    "Loki attack, Puny God",
			Priority:       "low",
			Status:         "solved",
			SubmitterID:    2,
			AssigneeID:     1,
			OrganizationID: 2,
			Tags: []string{
				"loki",
			},
			HasIncidents: true,
			DueAt:        "2019-05-21T11:10:28 -10:00",
			Via:          "web",
		},
	},
}
var MockUsers = Users{
	BaseRecords: BaseRecords{
		Size: 2,
	},
	Items: []schema.User{
		schema.User{
			ID:             1,
			URL:            "http://localhost/user/1",
			ExternalID:     "external-id-1",
			Name:           "Nick Fury",
			Alias:          "Scorpio",
			CreatedAt:      "2016-05-21T11:10:28 -10:00",
			Active:         true,
			Verified:       true,
			Shared:         true,
			Locale:         "en-US",
			Timezone:       "New York",
			LastLoginAt:    "2016-05-21T11:10:28 -10:00",
			Email:          "nick.fury@shield.gov.us",
			Phone:          "123456789",
			Signature:      "I still believe in heroes",
			OrganizationID: 1,
			Tags: []string{
				"nickfury",
				"heroes",
			},
			Suspended: false,
			Role:      "admin",
		},
		schema.User{
			ID:             2,
			URL:            "http://localhost/user/2",
			ExternalID:     "external-id-2",
			Name:           "Steve Rogers",
			Alias:          "Captain America",
			CreatedAt:      "2016-05-21T11:10:28 -10:00",
			Active:         true,
			Verified:       true,
			Shared:         true,
			Locale:         "en-US",
			Timezone:       "New York",
			LastLoginAt:    "2016-05-21T11:10:28 -10:00",
			Email:          "steve.rogers@avengers.world",
			Phone:          "987654321",
			Signature:      "I can do this all day",
			OrganizationID: 2,
			Tags: []string{
				"steverogers",
				"heroes",
			},
			Suspended: false,
			Role:      "admin",
		},
	},
}

var MockDataSet = DataSet{
	Organizations: MockOrganizations,
	Tickets:       MockTickets,
	Users:         MockUsers,
}
