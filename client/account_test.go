package client

import (
	"github.com/aalkan/whm-go"
	"testing"
)

func TestUserList(t *testing.T) {
	c := NewClient(&whm.Config{
		Host:     "#",
		Port:     "#",
		ApiToken: "#",
		User:     "#",
	})
	params := &UserListParams{
		Search:       "username",
		SearchMethod: "regex",
		SearchType:   "user",
	}
	response := &UserListResponse{}
	err := c.UserList(params, response)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteAccount(t *testing.T) {
	c := NewClient(&whm.Config{
		Host:     "#",
		Port:     "#",
		ApiToken: "#",
		User:     "#",
	})
	params := &DeleteAccountParams{
		User: "#",
	}
	response := &DeleteAccountResponse{}
	err := c.DeleteAccount(params, response)
	if err != nil {
		t.Error(err)
	}
}
