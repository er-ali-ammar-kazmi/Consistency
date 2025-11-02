package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RestApp struct {
	calci Calculator
}

func NewRestApp() *RestApp {
	return &RestApp{calci: NewCalculator()}
}

type operatorsModel struct {
	A float32 `json:"a"`
	B float32 `json:"b"`
}

type msgModel struct {
	Value string `json:"value"`
}

func (restService RestApp) GetAddition(w http.ResponseWriter, req *http.Request) {

	bodyByte, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Can't able to read request Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to read request Body %v", err.Error())
		return
	}

	var reqPayload operatorsModel
	if err := json.Unmarshal(bodyByte, &reqPayload); err != nil {
		log.Printf("Can't able to Unmarshal request Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to Unmarshal request Body %v", err.Error())
		return
	}

	if reqPayload.A == 0 || reqPayload.B == 0 {
		log.Println("missing required fields in Request Body")
		fmt.Fprintf(w, "missing required fields in Request Body")
		return
	}

	output := restService.calci.Addition(reqPayload.A, reqPayload.B)
	result := fmt.Sprintf("Addition of %f and %f is %f", reqPayload.A, reqPayload.B, output)

	ans := msgModel{}
	ans.Value = result
	data, err := json.Marshal(ans)
	if err != nil {
		log.Printf("Can't able to Marshal response Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to Marshal response Body %v", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	log.Println("Request Processed:", result)
	w.Write(data)
}

func (restService RestApp) GetSubtraction(w http.ResponseWriter, req *http.Request) {

	bodyByte, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Can't able to read request Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to read request Body %v", err.Error())
		return
	}

	var reqPayload operatorsModel
	if err := json.Unmarshal(bodyByte, &reqPayload); err != nil {
		log.Printf("Can't able to Unmarshal request Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to Unmarshal request Body %v", err.Error())
		return
	}

	if reqPayload.A == 0 || reqPayload.B == 0 {
		log.Println("missing required fields in Request Body")
		fmt.Fprintf(w, "missing required fields in Request Body")
		return
	}

	output := restService.calci.Subtraction(reqPayload.A, reqPayload.B)
	result := fmt.Sprintf("Subtraction of %f and %f is %f", reqPayload.A, reqPayload.B, output)

	ans := &msgModel{}
	ans.Value = result
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(ans)
	if err != nil {
		log.Printf("Can't able to Marshal request Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to Marshal request Body %v", err.Error())
		return
	}
	log.Println("Request Processed:", result)
	w.Write(data)
}

func (restService RestApp) GetMultiplication(w http.ResponseWriter, req *http.Request) {

	bodyByte, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Can't able to read request Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to read request Body %v", err.Error())
		return
	}

	var reqPayload operatorsModel
	if err := json.Unmarshal(bodyByte, &reqPayload); err != nil {
		log.Printf("Can't able to Unmarshal request Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to Unmarshal request Body %v", err.Error())
		return
	}

	if reqPayload.A == 0 || reqPayload.B == 0 {
		log.Println("missing required fields in Request Body")
		fmt.Fprintf(w, "missing required fields in Request Body")
		return
	}

	output := restService.calci.Multiplication(reqPayload.A, reqPayload.B)
	result := fmt.Sprintf("Multiplication of %f and %f is %f", reqPayload.A, reqPayload.B, output)

	ans := &msgModel{}
	ans.Value = result
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(ans)
	if err != nil {
		log.Printf("Can't able to Marshal request Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to Marshal request Body %v", err.Error())
		return
	}
	log.Println("Request Processed:", result)
	w.Write(data)
}

func (restService RestApp) GetDivision(w http.ResponseWriter, req *http.Request) {

	bodyByte, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Can't able to read request Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to read request Body %v", err.Error())
		return
	}

	var reqPayload operatorsModel
	if err := json.Unmarshal(bodyByte, &reqPayload); err != nil {
		log.Printf("Can't able to Unmarshal request Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to Unmarshal request Body %v", err.Error())
		return
	}

	if reqPayload.A == 0 || reqPayload.B == 0 {
		log.Println("missing required fields in Request Body")
		fmt.Fprintf(w, "missing required fields in Request Body")
		return
	}

	output := restService.calci.Division(reqPayload.A, reqPayload.B)
	result := fmt.Sprintf("Division of %f and %f is %f", reqPayload.A, reqPayload.B, output)

	ans := &msgModel{}
	ans.Value = result
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(ans)
	if err != nil {
		log.Printf("Can't able to Marshal request Body %v", err.Error())
		fmt.Fprintf(w, "Can't able to Marshal request Body %v", err.Error())
		return
	}
	log.Println("Request Processed:", result)
	w.Write(data)
}
