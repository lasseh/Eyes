package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const url = "https://api.thousandeyes.com/alerts"

// Alarms is a helper struct for getting data from Thousandeyes
// (http://developer.thousandeyes.com/alerts/)
type Alarms struct {
	Alerts []*Alert `json:"alert"`
}

// Alert is an event from Thousandeyes, triggered by blarh
type Alert struct {
	Active         int        `json:"active"`
	DateStart      string     `json:"dateStart"`
	RuleExpression string     `json:"ruleExpression"`
	RuleName       string     `json:"ruleName"`
	Type           string     `json:"type"`
	TestName       string     `json:"testName"`
	ViolationCount int        `json:"violationCount"`
	WarningLevel   string     `json:"warninglevel"`
	WarningIcon    string     `json:"warningicon"`
	Message        string     `json:"message"`
	Monitors       []*Monitor `json:"monitors"`
}

// Monitor
type Monitor struct {
	Active         int    `json:"active"`
	DateStart      string `json:"dateStart"`
	MetricsAtEnd   string `json:"metricsAtEnd"`
	MetricsAtStart string `json:"metricsAtStart"`
	MonitorName    string `json:"monitorName"`
	Network        string `json:"network"`
	Prefix         string `json:"prefix"`
}

func GetStatus() *Alarms {
	var username = Config.Username
	var key = Config.Key
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Auth Header
	req.SetBasicAuth(username, key)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	alarms := Alarms{}

	err = decoder.Decode(&alarms)
	if err != nil {
		log.Fatal(err)
	}

	// Set properties based on the warning
	for _, value := range alarms.Alerts {
		// Set the border color
		if value.RuleName == "BGP Outage" {
			value.WarningLevel = "danger"
		} else {
			value.WarningLevel = "warning"
		}

		// Set status message
		if value.RuleName == "BGP Outage" {
			value.Message = fmt.Sprintf("Outage from: %s", value.Monitors[0].Network)
		} else {
			value.Message = fmt.Sprintf("Degraded routing from: %s", value.Monitors[0].Network)
		}

		// Set the icon
		if value.RuleName == "BGP Outage" {
			value.WarningIcon = "fa-chain-broken"
		} else {
			value.WarningIcon = "fa-random"
		}
	}

	return &alarms
}
