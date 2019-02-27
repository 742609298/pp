package test

import "log"

func  main()  {
	type RequestBody struct {
		Email string `json:"email"`
	}
	requestBody := RequestBody{}
	requestBody.Email="1"


	requestBody2 := &RequestBody{}
	requestBody2.Email="2"

	log.Print("requestBody:",requestBody.Email)
	log.Print("requestBody2:",requestBody2.Email)
}
