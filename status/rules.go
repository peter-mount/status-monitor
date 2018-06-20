package status

import (
  "time"
)

type RulesEngine struct {
  // The lookup maps
  ComponentStatus   map[string]string         `yaml:"componentStatus"`
  LambStatus        map[string]string         `yaml:"lambStatus"`
  GrafanaState      map[string]string         `yaml:"grafanaState"`
  // The components
  Components        map[string]RuleComponent  `yaml:"components"`
  Rules             map[string]Rule           `yaml:"rules"`
  // Filters
  //filters           map[string]RuleFilter     `yaml:"-"`
}

// Component definition
type RuleComponent struct {
  // Name of the component
  Name          string          `yaml:"name"`
  // Optional description
  Description   string          `yaml:"description"`
  // Default component status when ok
  Status        string          `yaml:"status"`
  // if true then delete when no active incidents
  Delete        string          `yaml:"delete"`
  // If set then the delay before deleting the component
  DeleteDelay   time.Duration   `yaml:"deleteDelay"`
}

type Rule struct {
  // Optional component to create/update
  Component     string
  // Array of rules
  Rule        []RuleDef
}

type RuleDef struct {
  // Status to set the component
  ComponentStatus string              `yaml:"componentStatus"`
  // Override status mapping
  Status          map[string]string   `yaml:"status"`
  // Filters
  Filters         map[string]int64
}

func (r *RulesEngine) Name() string {
  return "RulesEngine"
}

func (r *RulesEngine) PostInit() error {


  return nil
}
