package status

import (
  "github.com/peter-mount/status-monitor/status/bolt"
  "time"
)

// An incident on the status page.
// The JSON format of this struct matches the output of the lambstatus incidents api
type Incident struct {
  Id          string            `json:"incidentID"`
  Name        string            `json:"name"`
  Created     time.Time         `json:"createdAt"`
  Updated     time.Time         `json:"updatedAt"`
  updates   []IncidentUpdate    `json:"incidentUpdates"`
}

type IncidentUpdate struct {
  Id        string      `json:"incidentID"`
  UpdateId  string      `json:"incidentUpdateID"`
  Status    string      `json:"incidentStatus"`
  Message   string      `json:"message"`
  Created   time.Time   `json:"createdAt"`
  Updated   time.Time   `json:"updatedAt"`
}

// An internal incident message
type IncidentMessage struct {
  Name    string
  Title   string
  Status  string
  Message string
}

// UpdateIncident updates our copy of an incident and either creates or updates
// it on the Status page
func (s *Status ) UpdateIncident( i *IncidentMessage ) error {
  return s.bolt.Update( func( tx *bolt.Tx ) error {
    return nil
  } )
}
