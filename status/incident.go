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

const (
  // Max duration since the last update on a resolved issue for us to ignore
  // reopening it
  MAX_UPDATE_AGE = time.Hour
  // Max duration since a resolved issue was created for us to ignore reopening it
  MAX_ISSUE_AGE = 3 * time.Hour
)

func (s *Status ) getIncident( bucket *bolt.Bucket, i *IncidentMessage ) (*Incident,error) {
  response := &Incident{}

  if bucket.GetJSON( i.Name, response ) {
    log.Printf( "Found incident: %s - %s", response.Id, i.Name )

    ok, err := s.get( "/api/v0/incidents/" + response.Id , response )
    if !ok || err != nil {
      return nil, err
    }

    // Get the most recent status
    status := ""
    if len(response.Updates) > 0 {
      se := response.Updates[0]
      for _, s := range response.Updates {
        if s.Created.After( se.Created ) || s.Updated.After( se.Updated ) {
          se = s
        }
      }
      status = se.Status
    }

    if status == "Resolved" {
      // If the issue is Resolved and the last update was yesterday, or it was
      // created yesterday with it being resolved today then delete it locally
      // so we will create a new issue.
      now := time.Now().UTC()
      updated := response.Updated.UTC()
      created := response.Created.UTC()

      if now.Sub( updated ) >= MAX_UPDATE_AGE || now.Sub( created ) >= MAX_ISSUE_AGE {
        log.Printf( "Ignoring incident %s as too old", response.Id )
        err = bucket.Delete( i.Name )
        return nil, err
      }
    }

    // Active incident
    return response, nil
  } else {
    return nil, nil
  }
}

// UpdateIncident updates our copy of an incident and either creates or updates
// it on the Status page
func (s *Status ) UpdateIncident( i *IncidentMessage ) error {
  return s.bolt.Update( func( tx *bolt.Tx ) error {
    bucket, err := tx.CreateBucketIfNotExists( INCIDENT_BUCKET )
    if err != nil {
      return err
    }

    response, err := s.getIncident( bucket, i )
    if err != nil {
      return err
    }

    if response != nil {
      log.Printf( "Updating incident: %s - %s", response.Id, i.Name )

      // Update the status, so no Title
      i.Title = ""

      ok, err := s.patch( "/api/v0/incidents/" + response.Id, i, response )
      if err != nil {
        return err
      }
      if !ok {
        log.Println( "No response" )
        return nil
      }
    } else {
      log.Printf( "New incident: %s", i.Name )

      response = &Incident{}

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
