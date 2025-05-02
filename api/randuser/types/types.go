package types

import (
	"time"
)

type Result struct {
	Results []User `json:"results"`
	Info    Info   `json:"info"`
}

type Info struct {
	Seed    string  `json:"seed"`
	Results int64   `json:"results"`
	Page    int64   `json:"page"`
	Version float64 `json:"version"`
}

type User struct {
	Gender     string   `json:"gender"`
	Name       Name     `json:"name"`
	Location   Location `json:"location"`
	Email      string   `json:"email"`
	Login      Login    `json:"login"`
	Dob        Dob      `json:"dob"`
	Registered Dob      `json:"registered"`
	Phone      string   `json:"phone"`
	Cell       string   `json:"cell"`
	ID         ID       `json:"id"`
	Picture    Picture  `json:"picture"`
	Nat        string   `json:"nat"`
}

type Dob struct {
	Date time.Time `json:"date"`
	Age  int64     `json:"age"`
}

type ID struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Location struct {
	Street      Street      `json:"street"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	Postcode    int64       `json:"postcode"`
	Coordinates Coordinates `json:"coordinates"`
	Timezone    Timezone    `json:"timezone"`
}

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Street struct {
	Number int64  `json:"number"`
	Name   string `json:"name"`
}

type Timezone struct {
	Offset      string `json:"offset"`
	Description string `json:"description"`
}

type Login struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Md5      string `json:"md5"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}
