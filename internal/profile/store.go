package profile

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".config", "req")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "profiles.json"), nil
}

func LoadAll() (map[string]Profile, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return make(map[string]Profile), nil
	}
	if err != nil {
		return nil, err
	}
	var profiles map[string]Profile
	if err := json.Unmarshal(data, &profiles); err != nil {
		return nil, err
	}
	return profiles, nil
}

func Save(p Profile) error {
	profiles, err := LoadAll()
	if err != nil {
		return err
	}
	profiles[p.Name] = p
	return writeAll(profiles)
}

func Get(name string) (Profile, error) {
	profiles, err := LoadAll()
	if err != nil {
		return Profile{}, err
	}
	p, ok := profiles[name]
	if !ok {
		return Profile{}, fmt.Errorf("profile %q not found", name)
	}
	return p, nil
}

func Delete(name string) error {
	profiles, err := LoadAll()
	if err != nil {
		return err
	}
	if _, ok := profiles[name]; !ok {
		return fmt.Errorf("profile %q not found", name)
	}
	delete(profiles, name)
	return writeAll(profiles)
}

func writeAll(profiles map[string]Profile) error {
	path, err := configPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(profiles, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
