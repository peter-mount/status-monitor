package status

import (
  "github.com/peter-mount/golib/kernel"
  "github.com/peter-mount/golib/kernel/cron"
  "github.com/peter-mount/golib/rest"
)

type Status struct {
  restService  *rest.Server
  cron         *cron.CronService
}

func (s *Status) Name() string {
  return "status-monitor"
}

func (s *Status) Init( k *kernel.Kernel ) error {

  service, err := k.AddService( &cron.CronService{} )
  if err != nil {
    return err
  }
  s.cron = (service).(*cron.CronService)

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
