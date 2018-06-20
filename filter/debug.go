package filter

import (
  "log"
  "strings"
)

func (r *Node) LogTree() {
  r.logTree( 0 )
}

func (r *Node) logTree( depth int ) {
  s := strings.Repeat( " ", depth ) + "+ %s"

  if r.test {
    log.Printf( s + " INFINITE LOOP?", r.token )
    return
  }
  r.test = true
  defer func() {r.test=false} ()

  if r.handler== nil {
    log.Printf( s + " (%d)", r.token, r.value )
  } else {
    log.Printf( s, r.token )
  }

  if r.lhs != nil {
    r.lhs.logTree( depth+1 )
  }
  if r.rhs != nil {
    r.rhs.logTree( depth+1 )
  }
}
