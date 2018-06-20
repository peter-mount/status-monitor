package status

import (
  "fmt"
  "github.com/peter-mount/golib/kernel"
  "github.com/peter-mount/golib/kernel/bolt"
  "github.com/peter-mount/golib/kernel/cron"
  "github.com/peter-mount/golib/rest"
  "log"
  "os"
  "github.com/peter-mount/status-monitor/filter"
)

const (
  INCIDENT_BUCKET = "incidents"
)
type Status struct {
  restService  *rest.Server
  cron         *cron.CronService
  bolt         *bolt.BoltService

  // API url & key
  url     string
  key     string
}

func (s *Status) Name() string {
  return "status-monitor"
}

func (s *Status) Init( k *kernel.Kernel ) error {

  parser := &filter.Parser{}
  n, err := parser.Parse( "hour between 4 and 2")
  if err != nil {
    log.Println( err )
  }
  n.LogTree()
  return fmt.Errorf( "test")

  s.url = os.Getenv( "STATUS_URL" )
  if s.url == "" {
    return fmt.Errorf( "STATUS_URL is required" )
  }

  s.key = os.Getenv( "STATUS_KEY" )
  if s.key == "" {
    return fmt.Errorf( "STATUS_KEY is required" )
  }

  service, err := k.AddService( &cron.CronService{} )
  if err != nil {
    return err
  }
  s.cron = (service).(*cron.CronService)

  service, err = k.AddService( &bolt.BoltService{ FileName: "/database/status.db" } )
  if err != nil {
    return err
  }
  s.bolt = (service).(*bolt.BoltService)

  service, err = k.AddService( &rest.Server{} )
  if err != nil {
    return err
  }
  s.restService = (service).(*rest.Server)

  return nil
}

func (s *Status) PostInit() error {

  s.restService.Handle( "/grafana", s.grafanaHandler ).Methods( "POST" )

  return nil
}

func (s *Status) Start() error {

  s.bolt.Update( func( tx *bolt.Tx ) error {
    _, err := tx.CreateBucketIfNotExists( INCIDENT_BUCKET )
    if err != nil {
      return err
    }

    return nil
  })

  // Run berth cleanup function every 10 minutes
  //s.cron.AddFunc( "0 0/10 * * * *", s.cleanup )

  return nil
}
