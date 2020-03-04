package main

import (
        "io/ioutil"

        "gopkg.in/yaml.v2"
)

type Owner struct {
        GroupByEmail []string `yaml:"group_by_email,omitempty"`
        UserByEmail  []string `yaml:"user_by_email,omitempty"`
        SpecialGroup []string `yaml:"special_group,omitempty"`
}

type Writer struct {
        GroupByEmail []string `yaml:"group_by_email,omitempty"`
        UserByEmail  []string `yaml:"user_by_email,omitempty"`
        SpecialGroup []string `yaml:"special_group,omitempty"`
}

type Reader struct {
        GroupByEmail []string `yaml:"group_by_email,omitempty"`
        UserByEmail  []string `yaml:"user_by_email,omitempty"`
        SpecialGroup []string `yaml:"special_group,omitempty"`
}

type View struct {
        DatasetID string `yaml:"dataset_id,omitempty"`
        ViewID    string `yaml:"view_id,omitempty"`
}

// Config object to be loaded with configuration from YAML file
type Config struct {
        Project  string `yaml:"project,omitempty"`
        Datasets []struct {
                Name   string `yaml:"name,omitempty"`
                Owner  Owner  `yaml:"owner,omitempty"`
                Writer Writer `yaml:"writer,omitempty"`
                Reader Reader `yaml:"reader,omitempty"`
                View   []View `yaml:"view,omitempty"`
        } `yaml:"datasets,omitempty"`
}

// LoadFromFile return Config object according to a YAML file
func (conf *Config) LoadFromFile(file string) error {
        yamlFile, err := ioutil.ReadFile(file)
        if err == nil {
                err = yaml.Unmarshal(yamlFile, conf)
        }
        return err
}