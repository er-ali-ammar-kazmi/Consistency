package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"practise/applications/protobuf"

	"google.golang.org/protobuf/encoding/protojson"
)

type RestApp struct {
	calci Calculator
}

func NewRestApp() *RestApp {
	return &RestApp{calci: NewCalculator()}
}

func (restService RestApp) GetAddition(w http.ResponseWriter, req *http.Request) {

	bodyByte, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Can't able to read request Body %v", err.Error())
		return
	}

	var reqPayload protobuf.Operators
	if err := protojson.Unmarshal(bodyByte, &reqPayload); err != nil {
		fmt.Fprintf(w, "Can't able to Unmarshal request Body %v", err.Error())
		return
	}

	if reqPayload.A == 0 || reqPayload.B == 0 {
		fmt.Fprintf(w, "missing required fields in Request Body")
		return
	}

	output := restService.calci.Addition(reqPayload.A, reqPayload.B)
	result := fmt.Sprintf("Addition of %f and %f is %f", reqPayload.A, reqPayload.B, output)

	ans := &protobuf.Msg{}
	ans.Value = result
	data, err := protojson.Marshal(ans)
	if err != nil {
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
		fmt.Fprintf(w, "Can't able to read request Body %v", err.Error())
		return
	}

	var reqPayload protobuf.Operators
	if err := protojson.Unmarshal(bodyByte, &reqPayload); err != nil {
		fmt.Fprintf(w, "Can't able to Unmarshal request Body %v", err.Error())
		return
	}

	if reqPayload.A == 0 || reqPayload.B == 0 {
		fmt.Fprintf(w, "missing required fields in Request Body")
		return
	}

	output := restService.calci.Subtraction(reqPayload.A, reqPayload.B)
	result := fmt.Sprintf("Subtraction of %f and %f is %f", reqPayload.A, reqPayload.B, output)

	ans := &protobuf.Msg{}
	ans.Value = result
	w.Header().Set("Content-Type", "application/json")
	data, err := protojson.Marshal(ans)
	if err != nil {
		fmt.Fprintf(w, "Can't able to Marshal request Body %v", err.Error())
		return
	}
	log.Println("Request Processed:", result)
	w.Write(data)
}

func (restService RestApp) GetMultiplication(w http.ResponseWriter, req *http.Request) {

	bodyByte, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Can't able to read request Body %v", err.Error())
		return
	}

	var reqPayload protobuf.Operators
	if err := protojson.Unmarshal(bodyByte, &reqPayload); err != nil {
		fmt.Fprintf(w, "Can't able to Unmarshal request Body %v", err.Error())
		return
	}

	if reqPayload.A == 0 || reqPayload.B == 0 {
		fmt.Fprintf(w, "missing required fields in Request Body")
		return
	}

	output := restService.calci.Multiplication(reqPayload.A, reqPayload.B)
	result := fmt.Sprintf("Multiplication of %f and %f is %f", reqPayload.A, reqPayload.B, output)

	ans := &protobuf.Msg{}
	ans.Value = result
	w.Header().Set("Content-Type", "application/json")
	data, err := protojson.Marshal(ans)
	if err != nil {
		fmt.Fprintf(w, "Can't able to Marshal request Body %v", err.Error())
		return
	}
	log.Println("Request Processed:", result)
	w.Write(data)
}

func (restService RestApp) GetDivision(w http.ResponseWriter, req *http.Request) {

	bodyByte, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Can't able to read request Body %v", err.Error())
		return
	}

	var reqPayload protobuf.Operators
	if err := protojson.Unmarshal(bodyByte, &reqPayload); err != nil {
		fmt.Fprintf(w, "Can't able to Unmarshal request Body %v", err.Error())
		return
	}

	if reqPayload.A == 0 || reqPayload.B == 0 {
		fmt.Fprintf(w, "missing required fields in Request Body")
		return
	}

	output := restService.calci.Division(reqPayload.A, reqPayload.B)
	result := fmt.Sprintf("Division of %f and %f is %f", reqPayload.A, reqPayload.B, output)

	ans := &protobuf.Msg{}
	ans.Value = result
	w.Header().Set("Content-Type", "application/json")
	data, err := protojson.Marshal(ans)
	if err != nil {
		fmt.Fprintf(w, "Can't able to Marshal request Body %v", err.Error())
		return
	}
	log.Println("Request Processed:", result)
	w.Write(data)
}
