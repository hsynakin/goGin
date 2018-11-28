package main

import "time"

var employees = map[string]Employee{
	"962134": Employee{
		Id:        962134,
		FirstName: "Jennifer",
		LastName:  "Watson",
		Position:  "CEO",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  30,
	},
	"176158": Employee{
		Id:        176158,
		FirstName: "Allison",
		LastName:  "Jane",
		Position:  "CEO",
		StartDate: time.Now().Add(-4 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  30,
	},
	"423313": Employee{
		Id:        423313,
		FirstName: "das",
		LastName:  "dasg",
		Position:  "CTO",
		StartDate: time.Now().Add(-6 * time.Hour * 24 * 365),
		TotalPTO:  20,
	},
	"765234": Employee{
		Id:        765234,
		FirstName: "ytr",
		LastName:  "dasıug",
		Position:  "Worker Bee",
		StartDate: time.Now().Add(-12 * time.Hour * 24 * 365),
		TotalPTO:  25,
	},
}

var TimesOff = map[string][]TimeOff{
	"962134": []TimeOff{
		{
			Type:      "Holiday",
			Amount:    8.,
			StartDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    "Taken",
		}, {
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 8, 1, 0, 0, 0, 0, time.UTC),
			Status:    "Scheduled",
		}, {
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 12, 1, 0, 0, 0, 0, time.UTC),
			Status:    "Requested",
		},
	},
}

type Employee struct {
	Id        uint
	FirstName string
	LastName  string
	StartDate time.Time
	Position  string
	TotalPTO  float32
	Status    string
	TimesOff  []TimeOff
}

type TimeOff struct {
	Type      string
	Amount    float32
	StartDate time.Time
	Status    string
}
