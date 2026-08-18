# Cl0ud-C0nfig-2-But4ane

This is a prototype for Go transpiler. It converts **cloud-config** YAML files (traditionally used with cloud-init) into **Butane** YAML for provisioning Flatcar Container Linux systems. 

Because Flatcar uses Butane/Ignition for provisioning, existing infrastructure that relies heavily on cloud-config directives needs to be translated. This tool bridges the gap by translating the most common cloud-config directives into their Butane equivalents.

## Features & Supported Translations

The transpiler currently supports converting three main blocks of a cloud-config file: `users`, `write_files`, and `runcmd`.

### 1. Users
Cloud-config `users` declarations are translated into the Butane `passwd.users` node.
- **Attributes Supported**: `name`, `shell`, `ssh_authorized_keys`.
- **Groups**: Comma-separated `groups` strings are parsed and split into a YAML array for Butane.
- **Password Locking**: If `lock_passwd` is set to `true`, the Butane user's `password_hash` is explicitly set to `!` to lock the account.

### 2. Write Files
Cloud-config `write_files` blocks are translated into the Butane `storage.files` node.
- **Attributes Supported**: `path`, `content`, `permissions`, `owner`, `encoding`.
- **Base64 Decoding**: If the `encoding` is set to `base64` or `b64`, the transpiler decodes the content into plaintext before injecting it inline into Butane's `contents.inline`.
- **Permissions**: Octal permission strings (like `"0644"`) are parsed and correctly typed as integer modes in Butane.
- **Ownership**: The `owner` string (e.g., `root:root`) is parsed to populate Butane's `user.name` and `group.name` objects.

### 3. Runcmd
Butane does not natively have an equivalent to cloud-config's `runcmd` feature (which runs arbitrary shell commands late in the boot process). 
To solve this, the transpiler orchestrates a multi-step conversion:
1. **Script Generation**: It aggregates all `runcmd` commands into a single `bash` script string .
2. **File Creation**: It embeds this script into `/etc/cc-runcmd.sh` via the `storage.files` Butane node with `0755` permissions.
3. **Service Provisioning**: It creates a new systemd oneshot unit named `cc-runcmd.service` in the `systemd.units` node. This service runs the script and is configured to start after `network-online.target`.

## How to Use

We can use the tool by running main.go file. It accepts input via a file flag or standard input, and outputs to a file or standard output.

```bash
go run cmd/main.go -f testdata/example1.yaml -o butane.yaml
```

Example cloud-config and its corresponding Butane is present in [testdata](https://github.com/shinigami-777/Cl0ud-C0nfig-2-But4ane/tree/main/testdata).
### CLI Flags
- `-f`: Input cloud-config file
- `-o`: Output Butane file
