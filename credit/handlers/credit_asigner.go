package handlers

import (
    "errors"
    "encoding/json"
    "net/http"
)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type CreditAssignerImpl struct{}

type Credit struct {
    Investment int32 `json:"investment"`
}

func AssigmentInvestment(w http.ResponseWriter, r *http.Request) {
    creditAsigner := &CreditAssignerImpl{}

    var credit Credit

    err := json.NewDecoder(r.Body).Decode(&credit)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    credit300, credit500, credit700, err := creditAsigner.Assign(credit.Investment)

    if err != nil {
        http.Error(w, "", http.StatusBadRequest)
        return
    }

    credits := map[string]int32{
        "credit_type_300": credit300,
        "credit_type_500": credit500,
        "credit_type_700": credit700,
    }

    jsonResponse, err := json.Marshal(credits)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func (ca *CreditAssignerImpl) Assign(investment int32) (int32, int32, int32, error) {
	if investment < 300 || investment%100 != 0 {
		return 0, 0, 0, errors.New("invalid investment amount")
	}

	for c1 := int32(1); c1 <= (investment / 300) / 2; c1++ {
		for c2 := int32(1); c2 < (investment / 500) / 2; c2++ {
			for c3 := int32(1); c3 <= (investment / 700) / 2; c3++ {
                sum := (c1 * 300) + (c2 * 500) + (c3 * 700)
				if sum == investment {
					return c1, c2, c3, nil
				}
			}
		}
	}

	return 0, 0, 0, errors.New("no valid assignment found")
}
