package filter

// This is the main parser function - in a separate file for maintainability
func (p *Parser) parse( n *Node ) (*Node,error) {

  err := p.scan()
  if err != nil {
    return nil, err
  }

  switch p.token() {
    case "between":
      n, err = p.parseBinaryOpSep( n, betweenHandler, "and" )

    case "hour":
      n, err = p.parseUnaryOp( n, hourHandler )

    default:
      // Check for constant & if not then fail
      v, err := p.tokenInt()
      if err == nil {
        err = n.append( p.newConstant( v ) )
        if err != nil {
          return nil, err
        }
      } else {
        return nil, p.unknownToken()
      }
  }

  return n, err
}
