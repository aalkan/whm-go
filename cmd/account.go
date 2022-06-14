package cmd

func (c *client) UserList(params *UserListParams, resonse *UserListResponse) error {
	if err := c.request("GET", "/listaccts", params, resonse); err != nil {
		return err
	}
	return nil
}

func (c *client) CreateAccount(params *CreateAccountParams, response *CreateAccountResponse) error {
	if err := c.request("POST", "/createacct", params, response); err != nil {
		return err
	}
	return nil
}

func (c *client) DeleteAccount(params *DeleteAccountParams, response *DeleteAccountResponse) error {
	if err := c.request("POST", "/deleteacct", params, response); err != nil {
		return err
	}
	return nil
}

func (c *client) SuspendAccount(params *SuspendAccountParams, response *SuspendMetadataResponse) error {
	if err := c.request("POST", "/suspendacct", params, response); err != nil {
		return err
	}
	return nil
}

func (c *client) UnsuspendAccount(params *UnsuspendAccountParams, response *SuspendMetadataResponse) error {
	if err := c.request("POST", "/unsuspendacct", params, response); err != nil {
		return err
	}
	return nil
}

func (c *client) ChangeAccountPassword(params *ChangePasswordParams, response *ChangePasswordResponse) error {
	if err := c.request("POST", "/passwd", params, response); err != nil {
		return err
	}
	return nil
}

func (c *client) ChangeAccountPackage(params *ChangePackageParams, response *ChangePackageResponse) error {
	if err := c.request("POST", "/changepackage", params, response); err != nil {
		return err
	}
	return nil
}
