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
		log.Println(fmt.Sprintf("Can't able to read request Body %v", err.Error()))
		fmt.Fprintf(w, "Can't able to read request Body %v", err.Error())
		return
	}

	var reqPayload operatorsModel
	if err := json.Unmarshal(bodyByte, &reqPayload); err != nil {
		log.Println(fmt.Sprintf("Can't able to Unmarshal request Body %v", err.Error()))
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

	ans := &msgModel{}
	ans.Value = result
	w.Header().Set("Content-Type", "application/x-protobuf")
	data, err := json.Marshal(ans)
	if err != nil {
		log.Println(fmt.Sprintf("Can't able to Unmarshal request Body %v", err.Error()))
		fmt.Fprintf(w, "Can't able to Unmarshal request Body %v", err.Error())
		return
	}
	log.Println("Request Processed:", result)
	w.Write(data)
}
