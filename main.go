package Foo

import (
	"Foo/auth"
	"Foo/types"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// load policy.json and parse to struct
func parsePolicy() types.Policy {
	// open policy.json
	file, err := os.Open("policy.json")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // parse json to struct
    var policy types.Policy
    if err := json.NewDecoder(file).Decode(&policy); err != nil {
        panic(err)
    }
    return policy
}

// check if user has priviledges to run this contract or not
func authorizeRequest(policy *types.Policy, role string, contract string) bool{
	if role, ok := policy.Roles[role]; ok {
		for _, c := role.Contracts{
			if(c == contract || c == "*"){
				return true
			}
		}
	}
	return false
}


func handler(w http.ResponseWriter, r *http.Request){
	// get token from header
	header := w.Header().Get("Authroization")
	if header == "" || !strings.HasPrefix(header, "Bearer "){
		http.Error(w, "Token Required", 200)
		return
	}

	token := strings.TrimPrefix(header, "Bearer")

	// check if token valid or not
	claims, err := auth.ValidateToken(token)
    if err != nil {
        http.Error(w, "Invalid Token", http.StatusUnauthorized)
        return
    }

    // read contract from request body
	bodyBytes, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Invalid Request Body", http.StatusBadRequest)
        return
    }

    contract := strings.TrimSpace(string(bodyBytes))

    // load policy.json file
    policy := parsePolicy()

    // authorize request
    if authorizeRequest(&policy, claims.Role, contract) {
        fmt.Fprintln(w, "Approved")
    } else {
        fmt.Fprintln(w, "Denied")
    }
}

func main(){
	http.HandleFunc("/iam", handler)
	http.ListenAndServe(": 8080", nil)
}