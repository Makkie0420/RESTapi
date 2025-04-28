package controllers

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "github.com/revel/revel"
)

type OracleUsers struct {
    *revel.Controller
}

const (
    oracleAPIBaseURL = "https://identity.<your-region>.oraclecloud.com/admin/v1"
    accessToken      = "<your-access-token-here>" // <<< Put your actual token here
)

// Helper: prepare authorized HTTP request
func prepareRequest(method, url string, body []byte) (*http.Request, error) {
    req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer "+accessToken)
    req.Header.Set("Content-Type", "application/json")
    return req, nil
}

// List Users
func (c OracleUsers) ListUsers() revel.Result {
    url := oracleAPIBaseURL + "/Users"
    req, err := prepareRequest("GET", url, nil)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    var result map[string]interface{}
    json.Unmarshal(body, &result)

    return c.RenderJSON(result)
}

// Get User by ID
func (c OracleUsers) GetUser(id string) revel.Result {
    url := oracleAPIBaseURL + "/Users/" + id
    req, err := prepareRequest("GET", url, nil)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    var result map[string]interface{}
    json.Unmarshal(body, &result)

    return c.RenderJSON(result)
}

// Create User
func (c OracleUsers) CreateUser() revel.Result {
    var user map[string]interface{}
    err := c.Params.BindJSON(&user)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": "Invalid JSON"})
    }

    jsonBody, _ := json.Marshal(user)
    url := oracleAPIBaseURL + "/Users"
    req, err := prepareRequest("POST", url, jsonBody)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    var result map[string]interface{}
    json.Unmarshal(body, &result)

    return c.RenderJSON(result)
}

// Update User
func (c OracleUsers) UpdateUser(id string) revel.Result {
    var updatedUser map[string]interface{}
    err := c.Params.BindJSON(&updatedUser)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": "Invalid JSON"})
    }

    jsonBody, _ := json.Marshal(updatedUser)
    url := oracleAPIBaseURL + "/Users/" + id
    req, err := prepareRequest("PUT", url, jsonBody)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    var result map[string]interface{}
    json.Unmarshal(body, &result)

    return c.RenderJSON(result)
}

// Delete User
func (c OracleUsers) DeleteUser(id string) revel.Result {
    url := oracleAPIBaseURL + "/Users/" + id
    req, err := prepareRequest("DELETE", url, nil)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }
    defer resp.Body.Close()

    if resp.StatusCode == 204 {
        return c.RenderJSON(map[string]string{"message": "User deleted"})
    }

    body, _ := ioutil.ReadAll(resp.Body)
    var result map[string]interface{}
    json.Unmarshal(body, &result)

    return c.RenderJSON(result)
}
