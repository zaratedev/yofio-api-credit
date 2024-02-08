package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "go.mongodb.org/mongo-driver/mongo"
)

var fakeCollection *mongo.Collection

func TestAssigmentInvestment(t *testing.T) {
    t.Skip("Test not implemented yet")
    requestBody := map[string]int32{"investment": 3000}
    bodyBytes, _ := json.Marshal(requestBody)
    bodyReader := bytes.NewReader(bodyBytes)
    
    req := httptest.NewRequest("POST", "/credit-assigment", bodyReader)
    w := httptest.NewRecorder()

    AssigmentInvestment(w, req, fakeCollection)

    if w.Code != http.StatusOK {
        t.Errorf("Waiting http status %d but received %d", http.StatusOK, w.Code)
    }

    expectedResponse := `{"credit_type_300":2,"credit_type_500":2,"credit_type_700":2}`
    if w.Body.String() != expectedResponse {
        t.Errorf("Waiting response %s but received %s", expectedResponse, w.Body.String())
    }
}

func TestBadRequest(t *testing.T) {
    t.Skip("Test not implemented yet")
    req := httptest.NewRequest("POST", "/credit-assigment", nil)
    w := httptest.NewRecorder()

    AssigmentInvestment(w, req, fakeCollection)

    if w.Code != http.StatusBadRequest {
        t.Errorf("Waiting http status %d but received %d", http.StatusBadRequest, w.Code)
    }
}

func TestAssigmentInvestmentWithInvalidAmount(t *testing.T) {
    t.Skip("Test not implemented yet")
    requestBody := map[string]int32{"investment": 400}
    bodyBytes, _ := json.Marshal(requestBody)
    bodyReader := bytes.NewReader(bodyBytes)

    req := httptest.NewRequest("POST", "/credit-assigment", bodyReader)
    w := httptest.NewRecorder()

    AssigmentInvestment(w, req, fakeCollection)

    if w.Code != http.StatusBadRequest {
        t.Errorf("Waiting http status %d but received %d", http.StatusBadRequest, w.Code)
    }
}

func TestAsigmentInvestmentWithValidAmount(t *testing.T) {
    creditAssigner := &CreditAssignerImpl{}
    investment := int32(6700)

    c300, c500, c700, _ := creditAssigner.Assign(investment)

    if c300 != 7 {
        t.Errorf("Expected %d but received %d", 7, c300)
    }

    if c500 != 5 {
        t.Errorf("Expected %d but received %d", 5, c500)
    }

    if c700 != 3 {
        t.Errorf("Expected %d but received %d", 5, c700)
    }
}

func TestAssignWithInvalitAmount(t *testing.T) {
    creditAssigner := &CreditAssignerImpl{}
    investment := int32(250)

    _, _, _, err := creditAssigner.Assign(investment)

    if err == nil {
        t.Errorf("An error was expected, but none was received.")
    }

    expectedError := "invalid investment amount"

    if err.Error() != expectedError {
        t.Errorf("exptected error %q, but received %q", expectedError, err.Error())
    }
}

func TestAssignInvalid(t *testing.T) {
    creditAssigner := &CreditAssignerImpl{}
    investment := int32(400)

    _, _, _, err := creditAssigner.Assign(investment)

    if err == nil {
        t.Errorf("An error was expected, but none was received.")
    }

    expectedError := "no valid assignment found"

    if err.Error() != expectedError {
        t.Errorf("exptected error %q, but received %q", expectedError, err.Error())
    }
}
