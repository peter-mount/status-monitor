package filter

func equalHandler( m *Value, n *Node ) int64 {
  return bool2int( n.invokeLhs(m) == n.invokeRhs(m) )
}

func notEqualHandler( m *Value, n *Node ) int64 {
  return bool2int( n.invokeLhs(m) != n.invokeRhs(m) )
}

func lessThanHandler( m *Value, n *Node ) int64 {
  return bool2int( n.invokeLhs(m) < n.invokeRhs(m) )
}

func lessThanEqualHandler( m *Value, n *Node ) int64 {
  return bool2int( n.invokeLhs(m) <= n.invokeRhs(m) )
}

func greaterThanEqualHandler( m *Value, n *Node ) int64 {
  return bool2int( n.invokeLhs(m) >= n.invokeRhs(m) )
}

func greaterThanHandler( m *Value, n *Node ) int64 {
  return bool2int( n.invokeLhs(m) > n.invokeRhs(m) )
}

func betweenHandler( m *Value, n *Node ) int64 {
  return bool2int( m.Value >= n.invokeLhs(m) && m.Value <= n.invokeRhs(m) )
}

func andHandler( m *Value, n *Node ) int64 {
  return bool2int( n.invokeLhsBool(m) && n.invokeRhsBool(m) )
}

func orHandler( m *Value, n *Node ) int64 {
  return bool2int( n.invokeLhsBool(m) || n.invokeRhsBool(m) )
}

func notHandler( m *Value, n *Node ) int64 {
  return bool2int( !int2bool( n.invokeLhsBool(m) ) )
}
