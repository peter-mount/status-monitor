package filter

import (
  "time"
)

// invokes the left hand side with the current hour (UTC) as the value
func hourHandler( m *Value, n *Node ) int64 {
  v := &Value{
    Name: "hour",
    Value: int64( time.Now().UTC().Hour() ),
  }
  return n.invokeLhs(v)
}
