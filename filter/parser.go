package filter

import (
  "fmt"
  "strings"
  "text/scanner"
)

type Parser struct {
  scanner   scanner.Scanner
  root     *Node
}

func (p *Parser) Parse( rule string ) ( *Node, error ) {
  // Root node is special
  p.root = &Node{ token: "ROOT", handler: rootHandler }

  p.scanner.Init( strings.NewReader( rule ) )
  p.scanner.Filename = "filter"

  // Parse the root but always return the root here
  _, err := p.parse( p.root )
  return p.root, err
}

// Just invokes the left hand side
func rootHandler( m *Value, n *Node ) int64 {
  return n.invokeLhs(m)
}

func (p *Parser) unknownToken() error {
  return fmt.Errorf( "Unknown token: \"%s\"", p.scanner.TokenText() )
}

func (p *Parser) token() string {
  return p.scanner.TokenText()
}

func (p *Parser) scan() error {
  if p.scanner.Scan() == scanner.EOF {
    return fmt.Errorf( "EOF" )
  }

  fmt.Printf( "%s: %s\n", p.scanner.Position, p.scanner.TokenText() )

  return nil
}

// Scan and parse the next value.
// This is the same as calling scan() then tokenInt()
func (p *Parser) scanInt64() (int64, error) {
  err := p.scan()
  if err != nil {
    return 0, err
  }
  v, err := p.tokenInt()
  return v, err
}

// Fail if the next token is not one thats expected
func (p *Parser) expect( s string ) error {
  err := p.scan()
  if err == nil && p.token() != s {
    err = fmt.Errorf( "Unexpected token \"%s\" - expected \"%s\"", p.token(), s )
  }
  return err
}

// Parse a unary operation, e.g. NOT v
func (p *Parser) parseUnaryOp( n *Node, h NodeHandler ) (*Node,error) {
  bn := p.new( h )

  err := n.append( bn )
  if err != nil {
    return nil, err
  }

  // parse nextNode
  _, err = p.parse( bn )
  if err != nil {
    return nil, err
  }

  // return the original node
  return n, err
}

// Parse a binary operation, e.g. n AND nextNode
func (p *Parser) parseBinaryOp( n *Node, h NodeHandler ) (*Node,error) {
  bn := p.new( h )

  // existing node is lhs
  n.replace( bn )

  // parse nextNode
  _, err := p.parse( bn )
  if err != nil {
    return nil, err
  }

  // return this new op node
  return bn, err
}

// Parse a binary operation with separator, e.g. BETWEEN a AND b
// For an operator requiring 2 params use "" for s. e.g. ATAN2 a b
func (p *Parser) parseBinaryOpSep( n *Node, h NodeHandler, s string ) (*Node,error) {
  bn := p.new( h )

  // this is just appended to
  err := n.append( bn )
  if err != nil {
    return nil, err
  }

  // lhs
  _, err = p.parse( bn )
  if err != nil {
    return nil, err
  }

  // Required separator, "" for not required
  if s != "" {
    err = p.expect( s )
    if err != nil {
      return nil, err
    }
  }

  // rhs
  _, err = p.parse( bn )
  if err != nil {
    return nil, err
  }

  return bn, nil
}
