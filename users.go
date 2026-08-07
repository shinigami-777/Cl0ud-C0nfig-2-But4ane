package main

import (
	"strings"
)

func ConvertUsers(ccUsers []User) []ButaneUser {
	if len(ccUsers) == 0 {
		return nil
	}
	var users []ButaneUser
	for _, cu := range ccUsers {
		u := ButaneUser{
			Name:              cu.Name,
			Shell:             cu.Shell,
			SSHAuthorizedKeys: cu.SSHAuthorizedKeys,
		}

		if cu.Groups != "" {
			groups := strings.Split(cu.Groups, ",")
			for i := range groups {
				groups[i] = strings.TrimSpace(groups[i])
			}
			u.Groups = groups
		}

		if cu.LockPasswd != nil && *cu.LockPasswd {
			u.PasswordHash = "!"
		}

		users = append(users, u)
	}
	return users
}
