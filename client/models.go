package client

type UserListParams struct {
	Search       string `json:"search"`
	SearchType   string `json:"searchtype"`   //user, domain, owner, ip, package.
	SearchMethod string `json:"searchmethod"` //exact ,regex , regex: contain string, exact same word
}

type UserListResponse struct {
	Metadata struct {
		Result  int    `json:"result"`
		Reason  string `json:"reason"`
		Command string `json:"command"`
		Version int    `json:"version"`
	} `json:"metadata"`
	Data struct {
		Acct []struct {
			Suspended                       int           `json:"suspended"`
			LegacyBackup                    int           `json:"legacy_backup"`
			Maxaddons                       string        `json:"maxaddons"`
			UID                             string        `json:"uid"`
			Suspendtime                     int           `json:"suspendtime"`
			Suspendreason                   string        `json:"suspendreason"`
			Partition                       string        `json:"partition"`
			OutgoingMailHold                int           `json:"outgoing_mail_hold"`
			UnixStartdate                   int           `json:"unix_startdate"`
			OutgoingMailSuspended           int           `json:"outgoing_mail_suspended"`
			Shell                           string        `json:"shell"`
			IsLocked                        int           `json:"is_locked"`
			User                            string        `json:"user"`
			Domain                          string        `json:"domain"`
			Diskused                        string        `json:"diskused"`
			Temporary                       int           `json:"temporary"`
			Theme                           string        `json:"theme"`
			MaxEmailPerHour                 string        `json:"max_email_per_hour"`
			Disklimit                       string        `json:"disklimit"`
			Maxlst                          string        `json:"maxlst"`
			HasBackup                       int           `json:"has_backup"`
			Maxpop                          string        `json:"maxpop"`
			Email                           string        `json:"email"`
			ChildNodes                      []interface{} `json:"child_nodes"`
			Plan                            string        `json:"plan"`
			Maxftp                          string        `json:"maxftp"`
			Owner                           string        `json:"owner"`
			Maxparked                       string        `json:"maxparked"`
			Inodeslimit                     string        `json:"inodeslimit"`
			IP                              string        `json:"ip"`
			MailboxFormat                   string        `json:"mailbox_format"`
			Ipv6                            []interface{} `json:"ipv6"`
			Maxsub                          string        `json:"maxsub"`
			MaxEmailacctQuota               string        `json:"max_emailacct_quota"`
			MaxDeferFailPercentage          string        `json:"max_defer_fail_percentage"`
			Maxsql                          string        `json:"maxsql"`
			Startdate                       string        `json:"startdate"`
			Inodesused                      int           `json:"inodesused"`
			MinDeferFailToTriggerProtection string        `json:"min_defer_fail_to_trigger_protection"`
			Backup                          int           `json:"backup"`
		} `json:"acct"`
	} `json:"data"`
}

type CreateAccountParams struct {
	Username               string `json:"username"`     // new account username
	Password               string `json:"password"`     // new account password
	Domain                 string `json:"domain"`       // account main domain
	BwLimit                string `json:"bwlimit"`      // account bandwidth limit in MB. Default unlimited.
	Cgi                    string `json:"cgi"`          // account cgi status. Default true.
	ContactEmail           string `json:"contactemail"` // account contact. Default empty string.
	Cpmod                  string `json:"cpmod"`        // account cpmod status. Default servers' theme. / paper_lantern
	Customip               string `json:"customip"`     // account custom ip. Default empty string.
	Dkim                   string `json:"dkim"`         // account dkim status. Default true.
	Featurelist            string `json:"featurelist"`
	Forcedns               string `json:"forcedns"`
	Frontpage              string `json:"frontpage"`
	Gid                    string `json:"gid"`
	Hasshell               string `json:"hasshell"`
	Hasuseregns            string `json:"hasuseregns"`
	Homedir                string `json:"homedir"`  // account home directory. Default /home/user.
	IP                     string `json:"ip"`       // the account has a dedicated IP address. y | n. Default n.
	Language               string `json:"language"` // account language. Default en.
	Owner                  string `json:"owner"`    // The name of the account owner.  defaults to the authenticated user.
	MailboxFormat          string `json:"mailbox_format"`
	MaxDeferFailPercentage string `json:"max_defer_fail_percentage"`
	MaxEmailPerHour        string `json:"max_email_per_hour"`
	MaxEmailacctQuota      string `json:"max_emailacct_quota"`
	Maxaddon               string `json:"maxaddon"`
	Maxftp                 string `json:"maxftp"`
	Maxlst                 string `json:"maxlst"`
	Maxpark                string `json:"maxpark"`
	Maxpop                 string `json:"maxpop"`
	Maxsql                 string `json:"maxsql"`
	Maxsub                 string `json:"maxsub"`
	Mxcheck                string `json:"mxcheck"`
	Pkgname                string `json:"pkgname"` // account plan. Default deafult plan.
	Plan                   string `json:"plan"`    // account plan. Default deafult plan.
	SavePkg                string `json:"savepkg"` // whether to save the account's settings as a new plan. Default 0.
}

type DeleteAccountParams struct {
	User    string `json:"user"`    // account username
	Keepdns string `json:"keepdns"` // whether to keep the account's DNS records. Default false.
}

type DeleteAccountResponse struct {
	Metadata struct {
		Reason  string `json:"reason"`
		Version int    `json:"version"`
		Command string `json:"command"`
		Output  struct {
			Raw string `json:"raw"`
		} `json:"output"`
		Result int `json:"result"`
	} `json:"metadata"`
}

type CreateAccountResponse struct {
	Data struct {
		Nameserver4      string      `json:"nameserver4"`
		Nameserverentry2 interface{} `json:"nameserverentry2"`
		Nameserver2      string      `json:"nameserver2"`
		Nameserver3      string      `json:"nameserver3"`
		Nameservera      interface{} `json:"nameservera"`
		Nameserverentry  interface{} `json:"nameserverentry"`
		Nameserverentry3 interface{} `json:"nameserverentry3"`
		Nameservera2     interface{} `json:"nameservera2"`
		IP               string      `json:"ip"`
		Package          string      `json:"package"`
		Nameservera4     interface{} `json:"nameservera4"`
		Nameserver       string      `json:"nameserver"`
		Nameserverentry4 interface{} `json:"nameserverentry4"`
		Nameservera3     interface{} `json:"nameservera3"`
	} `json:"data"`
	Metadata struct {
		Version int `json:"version"`
		Output  struct {
			Raw string `json:"raw"`
		} `json:"output"`
		Command string `json:"command"`
		Reason  string `json:"reason"`
		Result  int    `json:"result"`
	} `json:"metadata"`
}

type ChangePackageParams struct {
	User string `json:"user"` // the username of a cPanel account on the server.
	Pkg  string `json:"pkg"`  // an existing plan name on the server
}

type ChangePackageResponse struct {
	Metadata struct {
		Result  int `json:"result"`
		Version int `json:"version"`
		Output  struct {
			Raw string `json:"raw"`
		} `json:"output"`
		Command string `json:"command"`
		Reason  string `json:"reason"`
	} `json:"metadata"`
}

type ChangePasswordParams struct {
	User         string `json:"user"`           // the username of a cPanel account on the server.
	Password     string `json:"password"`       // a new secure password.
	Enabledigest bool   `json:"enabledigest"`   // Whether to enable Digest Authentication for the account. Default current digest.
	DbPassUpdate bool   `json:"db_pass_update"` //Whether to also change the account's MySQL?? password. This parameter defaults to 0.
}
type ChangePasswordResponse struct {
	App []string `json:"app"` //An array that contains the name of a service for which the function changed the account's password.
}

type SuspendAccountParams struct {
	User                 string `json:"user"`                    // the username of a cPanel account on the server.
	Reason               string `json:"reason"`                  // a reason for the account suspension.
	Disallowun           string `json:"disallowun"`              // Whether to allow only the root user to unsuspend the account. This parameter defaults to 0.
	LeaveFtpAcctsEnabled string `json:"leave-ftp-accts-enabled"` // Whether to skip suspension of the account's FTP accounts. This parameter defaults to 0.
}

type UnsuspendAccountParams struct {
	User                 string `json:"user"`                   // the username of a cPanel account on the server.
	RetainServiceProxies bool   `json:"retain-service-proxies"` // Whether to retain any service proxies on an account. This parameter defaults to 0.
}
type SuspendMetadataResponse struct {
	Metadata struct {
		Result  int `json:"result"`
		Version int `json:"version"`
		Output  struct {
			Raw string `json:"raw"`
		} `json:"output"`
		Command string `json:"command"`
		Reason  string `json:"reason"`
	} `json:"metadata"`
}
