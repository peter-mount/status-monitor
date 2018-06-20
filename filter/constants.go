package filter

import (
  "strconv"
)

const (
  FALSE = int64(0)
  TRUE = int64(1)
)

// Create a new node with a value
func (p *Parser) newConstant( v int64 ) *Node {
  return &Node{ token: p.token(), value: v }
}

// tokenInt the current token as an int64
func (p *Parser) tokenInt() ( int64, error ) {
  switch p.token() {
    case "true":
      return TRUE, nil
    case "false":
      return FALSE, nil
    default:
      v, err := strconv.Atoi( p.token() )
      return int64(v), err
  }
}
