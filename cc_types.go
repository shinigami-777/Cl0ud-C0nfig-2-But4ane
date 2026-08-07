package main

// CloudConfig represents the root of a cloud-config YAML file.
type CloudConfig struct {
	Users      []User      `yaml:"users,omitempty"`
	WriteFiles []WriteFile `yaml:"write_files,omitempty"`
	Runcmd     []any       `yaml:"runcmd,omitempty"` // array of string or []string
}

// User represents a user definition in cloud-config.
type User struct {
	Name              string   `yaml:"name,omitempty"`
	Groups            string   `yaml:"groups,omitempty"` // Comma-separated list of groups
	Shell             string   `yaml:"shell,omitempty"`
	SSHAuthorizedKeys []string `yaml:"ssh_authorized_keys,omitempty"`
	LockPasswd        *bool    `yaml:"lock_passwd,omitempty"`
}

// WriteFile represents a file to be written in cloud-config.
type WriteFile struct {
	Path        string `yaml:"path,omitempty"`
	Content     string `yaml:"content,omitempty"`
	Encoding    string `yaml:"encoding,omitempty"`    // base64, gz, etc.
	Permissions string `yaml:"permissions,omitempty"` // e.g. "0640"
	Owner       string `yaml:"owner,omitempty"`       // e.g. "root:root"
}
