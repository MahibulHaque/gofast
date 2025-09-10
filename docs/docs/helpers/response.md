---
sidebar_position: 2
sidebar_label: Response Package
---

# Response Package
The `response` package provides utility functions for sending JSON responses in a consistent and structured way across your Go applications. It abstracts the repetitive boilerplate of marshaling data, setting headers, and writing responses, making your HTTP handlers cleaner and easier to maintain.

The main functions are:

- `JSON` – Sends a JSON response with the given HTTP status code and data payload.

- `JSONWithHeaders` – Extends JSON by allowing custom HTTP headers to be set in the response.

Both functions use `json.MarshalIndent` to format the JSON output with indentation, making responses human-readable by default. Additionally, they ensure the `Content-Type` header is correctly set to `application/json`.

This package is especially useful for APIs where you want consistent, reusable patterns for returning structured JSON responses.

## Code implementation

```go
package response

import (
    "encoding/json"
    "net/http"
)

func JSON(w http.ResponseWriter, status int, data any) error {
    return JSONWithHeaders(w, status, data, nil)
}

func JSONWithHeaders(w http.ResponseWriter, status int, data any, headers http.Header) error {
    js, err := json.MarshalIndent(data, "", "\t")
    if err != nil {
        return err
    }

    js = append(js, '\n')

    for key, values := range headers {
        for _, value := range values {
            w.Header().Add(key, value)
        }
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(js)

    return nil
}
```

## Usage

Here is a small usage example of the `response` helper package:

```go
func createUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := request.DecodeJSON(w, r, &req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    user := processUser(req) // your business logic
    
    if err := response.JSON(w, http.StatusCreated, user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
```
