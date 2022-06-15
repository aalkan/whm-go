# whmcs-go

## Installation

```
go get github.com/aalkan/whm-go
```

## Usage

````
        c := NewClient(&whm.Config{
		Host:     "#",
		Port:     "#",
		ApiToken: "#",
		User:     "#",
	})
	params := &UserListParams{
		Search: "username",
		SearchMethod: "regex",
		SearchType: "user",
	}
	response := &UserListResponse{}
	err := c.UserList(params, response)
	if err != nil {
		fmt.Println(err)
	}
````