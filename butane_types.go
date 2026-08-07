package main

// ButaneConfig represents the root of a Butane YAML file.
type ButaneConfig struct {
	Variant string   `yaml:"variant"`
	Version string   `yaml:"version"`
	Passwd  *Passwd  `yaml:"passwd,omitempty"`
	Storage *Storage `yaml:"storage,omitempty"`
	Systemd *Systemd `yaml:"systemd,omitempty"`
}

// Passwd contains user and group definitions.
type Passwd struct {
	Users []ButaneUser `yaml:"users,omitempty"`
}

// ButaneUser represents a user definition in Butane.
// Renamed from User to avoid conflict with cloudconfig.User in the same package
type ButaneUser struct {
	Name              string   `yaml:"name,omitempty"`
	Groups            []string `yaml:"groups,omitempty"`
	Shell             string   `yaml:"shell,omitempty"`
	SSHAuthorizedKeys []string `yaml:"ssh_authorized_keys,omitempty"`
	PasswordHash      string   `yaml:"password_hash,omitempty"`
}

// Storage contains file and directory definitions.
type Storage struct {
	Files []File `yaml:"files,omitempty"`
}

// File represents a file definition in Butane.
type File struct {
	Path     string       `yaml:"path,omitempty"`
	Mode     *int         `yaml:"mode,omitempty"`
	User     *NodeUser    `yaml:"user,omitempty"`
	Group    *NodeGroup   `yaml:"group,omitempty"`
	Contents FileContents `yaml:"contents,omitempty"`
}

// NodeUser represents a user reference by name.
type NodeUser struct {
	Name string `yaml:"name"`
}

// NodeGroup represents a group reference by name.
type NodeGroup struct {
	Name string `yaml:"name"`
}

// FileContents represents the content of a file.
type FileContents struct {
	Inline string `yaml:"inline,omitempty"`
}

// Systemd contains systemd unit definitions.
type Systemd struct {
	Units []Unit `yaml:"units,omitempty"`
}

// Unit represents a systemd unit.
type Unit struct {
	Name     string `yaml:"name,omitempty"`
	Enabled  *bool  `yaml:"enabled,omitempty"`
	Contents string `yaml:"contents,omitempty"`
}

// IntPtr is a helper function to return a pointer to an int.
func IntPtr(i int) *int {
	return &i
}

// BoolPtr is a helper function to return a pointer to a bool.
func BoolPtr(b bool) *bool {
	return &b
}
