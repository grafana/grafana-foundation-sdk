package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	redDashboard, err := red(redConfig{
		DashboardTitle: "RED method",
		ServiceIDs:     []string{"sample-service", "payments", "front-gateway"},
	}).Build()
	if err != nil {
		panic(err)
	}
	dashboardJson, err := json.MarshalIndent(redDashboard, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dashboardJson))
}
