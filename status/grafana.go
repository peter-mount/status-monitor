package status

import (
  "github.com/peter-mount/golib/rest"
  "log"
)

type GrafanaMessage struct {
  Title     string                `json:"title"`
  Id        int                   `json:"ruleId"`
  RuleName  string                `json:"ruleName"`
  RuleUrl   string                `json:"ruleUrl"`
  State     string                `json:"state"`
  ImageUrl  string            `json:"imageUrl"`
  Message   string            `json:"message"`
  Matches   []GrafanaMatches  `json:"evalMatches"`
}

type GrafanaMatches struct {
  Metric  string              `json:"metric"`
  Tags    map[string]string   `json:"tags"`
  Value   int64               `json:"value"`
}

// grafanaHandler handles the inbound webhook from Grafana
func (s *Status) grafanaHandler( r *rest.Rest ) error {
  msg := &GrafanaMessage{}

  err := r.Body( msg )
  if err != nil {
    return err
  }

  imsg := &IncidentMessage{
    Name: msg.RuleName,
    Title: msg.Title,
    Message: msg.Message,
  }
  switch msg.State {
    case "ok":
      imsg.Status = "Resolved"
    case "paused":
      imsg.Status = "Investigating"
    case "alerting":
      imsg.Status = "Monitoring"
    case "pending":
      imsg.Status = "Monitoring"
    case "no_data":
      imsg.Status = "Monitoring"
    default:
      imsg.Status = "Monitoring"
  }

  err = s.UpdateIncident( imsg )
  if err != nil {
    log.Printf( "ERROR: %s", err )
    return err
  }

  r.Status( 200 )

  return nil
}
