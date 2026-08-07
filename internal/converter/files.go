package converter

import (
	"encoding/base64"
	"strconv"
	"strings"
)

func ConvertFiles(ccFiles []WriteFile) []File {
	if len(ccFiles) == 0 {
		return nil
	}
	var files []File
	for _, cf := range ccFiles {
		f := File{
			Path: cf.Path,
		}

		if cf.Permissions != "" {
			// Parse octal string like "0640" or "640"
			modeStr := cf.Permissions
			// In case they didn't put a leading 0, assume it's octal if length is 3 or 4
			mode, err := strconv.ParseInt(modeStr, 8, 32)
			if err == nil {
				f.Mode = IntPtr(int(mode))
			}
		}

		if cf.Owner != "" {
			parts := strings.SplitN(cf.Owner, ":", 2)
			f.User = &NodeUser{Name: strings.TrimSpace(parts[0])}
			if len(parts) > 1 {
				f.Group = &NodeGroup{Name: strings.TrimSpace(parts[1])}
			}
		}

		content := cf.Content
		if cf.Encoding == "base64" || cf.Encoding == "b64" {
			// Remove whitespace/newlines before decoding
			cleaned := strings.ReplaceAll(content, "\n", "")
			cleaned = strings.ReplaceAll(cleaned, "\r", "")
			cleaned = strings.ReplaceAll(cleaned, " ", "")
			decoded, err := base64.StdEncoding.DecodeString(cleaned)
			if err == nil {
				content = string(decoded)
			}
		}

		f.Contents = FileContents{
			Inline: content,
		}

		files = append(files, f)
	}
	return files
}
