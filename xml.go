package main

import "encoding/xml"

type StartTime struct {
	XMLName xml.Name `xml:"starttime"`
	Year    int      `xml:"year"`
	Month   int      `xml:"month"`
	Day     int      `xml:"day"`
	Hour    int      `xml:"hour"`
	Minute  int      `xml:"minute"`
	Second  int      `xml:"second"`
}

type Static struct {
	XMLName  xml.Name `xml:"static"`
	Duration int      `xml:"duration"`
	File     string   `xml:"file"`
}

type Transition struct {
	XMLName  xml.Name `xml:"transition"`
	Duration float64  `xml:"duration"`
	From     string   `xml:"from"`
	To       string   `xml:"to"`
}

type Background struct {
	XMLName   xml.Name  `xml:"background"`
	StartTime StartTime `xml:"starttime"`
	Comment   string    `xml:",comment"`
	Items     []interface{}
}
