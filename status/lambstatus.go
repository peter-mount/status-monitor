package status

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "net/http"
)

func (s *Status) get( path string, result interface{} ) ( bool, error ) {
  b, err := s.sendToLamb( "GET", path, nil, result )
  return b, err
}

func (s *Status) patch( path string, payload interface{}, result interface{} ) ( bool, error ) {
  b, err := s.sendToLamb( "PATCH", path, payload, result )
  return b, err
}

func (s *Status) post( path string, payload interface{}, result interface{} ) ( bool, error ) {
  b, err := s.sendToLamb( "POST", path, payload, result )
  return b, err
}

func (s *Status) sendToLamb( method, path string, payload interface{}, result interface{} ) ( bool, error ) {
  var req *http.Request
  var err error

  if payload == nil {
    req, err = http.NewRequest( method, s.url + path, nil )
    if err != nil {
      return false, err
    }
  } else {
    post, err := json.Marshal( payload )
    if err != nil {
      return false, err
    }

    req, err = http.NewRequest( method, s.url + path, bytes.NewReader( post ) )
    if err != nil {
      return false, err
    }
  }
  req.Header.Add( "x-api-key", s.key )
  req.Header.Add( "Content-Type", "application/json" )

  resp, err := http.DefaultClient.Do( req )
  //resp, err := http.Post( s.url + path, "application/json", bytes.NewReader( post ) )
  if err != nil {
    return false, err
  }

  defer resp.Body.Close()

  if resp.StatusCode == 404 {
    return false, nil
  }

  body, err := ioutil.ReadAll( resp.Body )
  if err != nil {
    return false, err
  }

  err = json.Unmarshal( body, result )
  if err != nil {
    return false, err
  }

  return true, nil
}
