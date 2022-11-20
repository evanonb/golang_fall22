package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
)

type Welcome struct {
	Name string
	Company string
	Date string
}

// type Welcome struct {
// 	Name string
// 	Time string
// }

type ShippingInfo struct {
	Shipper Shipper `json:"Shipper"`
	Recipient Recipient `json:"Recipient"`
}

type Shipper struct {
	CompanyName string `json:"shipperCompanyName"`
	Address string `json:"shipperAddress"`
	City string `json:"shipperCity"`
	State string `json:"shipperState"`
	Phone string `json:"shipperPhone"`
}

type Recipient struct {
	FirstName string `json:"recipientFirstName"`
	LastName string `json:"recipientLastName"`
	Address string `json:"shipperAddress"`
	City string `json:"shipperCity"`
	State string `json:"shipperState"`
	Phone string `json:"shipperPhone"`
}
// type JsonResponse struct {
// 	Value1 string `json:"key1"`
// 	Value2 string `json:"key2"`
// 	JsonNested JsonNested `json:"jsonNested"`
// }

// type JsonNested struct {
// 	NestedValue1 string `json:"nestedKey1"`
// 	NestedValue2 string `json:"nestedKey2"`
// }

func main() {
	// welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	welcome := Welcome{"Jimmy John", "Amazon", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/shipping-info-template.html"))

	shipper := Shipper{
		CompanyName: "Amazon",
		Address: "280 Bridgeport Blvd",
		City: "Newnan",
		State: "GA",
		Phone: "770-463-9837",
	}

	recipient := Recipient{
		FirstName: "Jimmy",
		LastName: "John",
		Address: "154 Macon Road",
		City: "Columbus",
		State: "GA",
		Phone: "407-069-4123",
	}

	shipinfo := ShippingInfo{
		Shipper: shipper,
		Recipient: recipient,
	}
	// nested := JsonNested{
	// 	NestedValue1: "first nested value",
	// 	NestedValue2: "second nested value",
	// }

	// jsonResp := JsonResponse{
	// 	Value1: "some Data",
	// 	Value2: "other Data",
	// 	JsonNested: nested,
	// }

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "shipping-info-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// third path, get/fetch, return a json object like an API, include 2 nested objects
	// {firstname:"", lastname:"", address:"", city}

	http.HandleFunc("/jsonShip", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(shipinfo)
	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
