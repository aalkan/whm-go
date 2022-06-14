package cmd

import (
	"fmt"
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
	params := &UserListParams{}
	response := &UserListResponse{}
	err := c.UserList(params, response)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Succeed", response)
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
	fmt.Println("Succeed", response)
}
