package user

import (
    "errors"
    "encoding/json"
)

type User struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Age      int    `json:"age"`
}

func (u *User) ToJSON() (string, error) {
    data, err := json.Marshal(u)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func FromJSON(data string) (*User, error) {
    var user User
    err := json.Unmarshal([]byte(data), &user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (u *User) Validate() error {
    if u.ID == "" {
        return errors.New("user ID cannot be empty")
    }
    if u.Name == "" {
        return errors.New("user name cannot be empty")
    }
    if u.Email == "" {
        return errors.New("user email cannot be empty")
    }
    if u.Age < 0 {
        return errors.New("user age cannot be negative or empty")
    }

    return nil
}