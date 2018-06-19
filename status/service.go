package status

import (
  "fmt"
  "github.com/peter-mount/golib/kernel"
  "github.com/peter-mount/golib/kernel/cron"
  "github.com/peter-mount/golib/rest"
  "github.com/peter-mount/status-monitor/status/bolt"
  "os"
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

  //s.restService.Handle( "/areas", s.tdGetAreas ).Methods( "GET" )

  return nil
}

func (s *Status) Start() error {

  // Run berth cleanup function every 10 minutes
  //s.cron.AddFunc( "0 0/10 * * * *", s.cleanup )

  return nil
}
