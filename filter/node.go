package filter

import (
  "fmt"
  "log"
)

type Value struct {
  Name  string
  Value int64
}

type NodeHandler func( *Value, *Node ) int64

// A node in the filter tree
type Node struct {
  token       string
  // parent node
  parent     *Node
  // left hand side
  lhs        *Node
  // right hand side
  rhs        *Node
  // handler for this node
  handler     NodeHandler
  // The value of this node
  value       int64
  test  bool
}

// Create a new node for a handler
func (p *Parser) new( f NodeHandler ) *Node {
  return &Node{ token: p.token(), handler: f }
}

// Invoke the handler of this node
func (n *Node) invoke( m *Value ) int64 {
  r := n.value
  log.Printf( "invoke:%s:%v", n.token,m)
  if n.handler != nil {
    r = n.handler( m, n )
  }
  log.Printf( "invoke:%s:%d:%v", n.token, r, m)
  return r
}

// set lhs or rhs if lhs is occupied
func (n *Node) append( next *Node ) error {
  if n.lhs == nil {
    log.Printf( "append lhs %s %s", n.token, next.token )
    n.lhs = next
  } else if n.rhs == nil {
    log.Printf( "append rhs %s %s", n.token, next.token )
    n.rhs = next
  } else {
    log.Printf( "%v\nlhs: %v\nrhs: %v", n, n.lhs, n.rhs )
    return fmt.Errorf( "Node full: %v %v", n, next )
  }
  next.parent = n
  return nil
}

// Replace this node in the tree with a new node and make this one the new node's
// lhs. Used when parsing a AND b where this is a.
func (n *Node) replace( next *Node ) error {
  if n.parent == nil {
    return fmt.Errorf( "No lhs for %s", next.token )
  }

  // add next to the parent instead of this node
  p := n.parent
  if p.lhs == n {
    p.lhs = next
  } else {
    p.rhs = n
  }

  // Add ourselves as the new lhs
  next.lhs = n

  // Fix the parentage
  next.parent = p
  n.parent = next

  return nil
}

// Invokes the left hand side node or returns false if none
func (n *Node) invokeLhs( m *Value ) int64 {
  if n.lhs != nil {
    return n.lhs.invoke(m)
  }
  return FALSE
}

// Invokes the right hand side node or returns false if none
func (n *Node) invokeRhs( m *Value ) int64 {
  if n.rhs != nil {
    return n.rhs.invoke(m)
  }
  return FALSE
}

func (n *Node) invokeLhsBool( m *Value ) bool {
  return int2bool( n.invokeLhs(m) )
}

func (n *Node) invokeRhsBool( m *Value ) bool {
  return int2bool( n.invokeRhs(m) )
}

// utility if int64 is 0 then false else true
func int2bool( v int64 ) bool {
  if v == 0 {
    return false
  }
  return true
}

func bool2int( b bool ) int64 {
  if b {
    return TRUE
  }
  return FALSE
}
