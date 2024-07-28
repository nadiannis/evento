package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/domain/request"
	"github.com/nadiannis/evento/internal/domain/response"
)

func main() {
	const numOfRequests = 100

	var wg sync.WaitGroup
	wg.Add(numOfRequests)

	customerJSON, err := createCustomer(request.CustomerRequest{Username: "ron", Balance: 1000000})
	if err != nil {
		fmt.Println(err)
		return
	}

	customer, _, err := unmarshalJSON[domain.Customer](customerJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	ticketsJSON, err := getTickets()
	if err != nil {
		fmt.Println(err)
		return
	}

	tickets, _, err := unmarshalJSON[[]domain.Ticket](ticketsJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	var ticket domain.Ticket
	for _, t := range *tickets {
		if t.Type == domain.TicketTypeCAT1 {
			ticket = t
			break
		}
	}

	responses := make(chan string, numOfRequests)

	for i := 0; i < numOfRequests; i++ {
		go func() {
			defer wg.Done()
			orderTicket(request.OrderRequest{
				CustomerID: customer.ID,
				TicketID:   ticket.ID,
				Quantity:   1,
			}, responses)
		}()
	}

	wg.Wait()
	close(responses)

	for res := range responses {
		fmt.Println(res)
	}
}

func createCustomer(input request.CustomerRequest) (string, error) {
	url := "http://localhost:8080/api/customers"
	jsonBytes, _ := json.Marshal(input)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("create customer response status:", res.Status)
	return string(body), nil
}

func getTickets() (string, error) {
	url := "http://localhost:8080/api/tickets"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("get tickets response status:", res.Status)
	return string(body), nil
}

func orderTicket(input request.OrderRequest, responses chan<- string) {
	url := "http://localhost:8080/api/orders"
	jsonBytes, _ := json.Marshal(input)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		responses <- fmt.Sprintf("error creating request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		responses <- fmt.Sprintf("error sending request: %v", err)
		return
	}
	defer res.Body.Close()

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	responses <- fmt.Sprintf("error reading response: %v", err)
	// 	return
	// }

	responses <- fmt.Sprintf("order response status: %v", res.Status)
}

func unmarshalJSON[T any](jsonString string) (*T, *response.ErrorResponse, error) {
	var genericResponse map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &genericResponse)
	if err != nil {
		return nil, nil, err
	}

	status, ok := genericResponse["status"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("response does not contain a valid status")
	}

	if status == string(response.Success) {
		var successRes response.SuccessResponse
		err = json.Unmarshal([]byte(jsonString), &successRes)
		if err != nil {
			return nil, nil, err
		}

		// Convert the Data field to the appropriate type
		dataBytes, err := json.Marshal(successRes.Data)
		if err != nil {
			return nil, nil, err
		}
		var data T
		err = json.Unmarshal(dataBytes, &data)
		if err != nil {
			return nil, nil, err
		}

		return &data, nil, nil
	} else {
		var errorRes response.ErrorResponse
		err = json.Unmarshal([]byte(jsonString), &errorRes)
		if err != nil {
			return nil, nil, err
		}

		return nil, &errorRes, nil
	}
}
