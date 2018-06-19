package status

import (
  "github.com/peter-mount/golib/kernel/bolt"
  "log"
  "time"
)

// An incident on the status page.
// The JSON format of this struct matches the output of the lambstatus incidents api
type Incident struct {
  Id          string            `json:"incidentID"`
  Name        string            `json:"name"`
  Created     time.Time         `json:"createdAt"`
  Updated     time.Time         `json:"updatedAt"`
  Updates   []IncidentUpdate    `json:"incidentUpdates"`
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
  //todo: Name/Title may need switching
  Name    string  `json:"-"`
  Title   string  `json:"name,omitempty"`
  Status  string  `json:"status"`
  Message string  `json:"message"`
}

// UpdateIncident updates our copy of an incident and either creates or updates
// it on the Status page
func (s *Status ) UpdateIncident( i *IncidentMessage ) error {
  return s.bolt.Update( func( tx *bolt.Tx ) error {
    bucket, err := tx.CreateBucketIfNotExists( INCIDENT_BUCKET )
    if err != nil {
      return err
    }

    response := &Incident{}

    // Lookup the incident name
    incident := &Incident{}
    if bucket.GetJSON( i.Name, incident ) {
      log.Printf( "Found incident: %s - %s", incident.Id, i.Name )

      // Update the status, so no Title
      i.Title = ""

      ok, err := s.patch( "/api/v0/incidents/" + incident.Id, i, response )
      if err != nil {
        return err
      }
      if !ok {
        log.Println( "No response" )
        return nil
      }
    } else {
      log.Printf( "New incident: %s", i.Name )

      ok, err := s.post( "/api/v0/incidents", i, response )
      if err != nil {
        return err
      }
      if !ok {
        log.Println( "No response" )
        return nil
      }
    }

    log.Printf( "Got response: %v", response )

    err = bucket.PutJSON( i.Name, response )
    if err != nil {
      return err
    }

    return nil
  } )
}
