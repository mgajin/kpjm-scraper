package parser

import "github.com/valyala/fastjson"

// Parser is wrapper around fastjson's parser
type Parser struct {
	*fastjson.Parser
}

// New returns new parser
func New() *Parser {
	return &Parser{&fastjson.Parser{}}
}

// Pool is wrapper around fastjson's parser pool
type Pool struct {
	fastjson.ParserPool
}

// NewPool returns new ParserPool
func NewPool() *Pool {
	return &Pool{
		fastjson.ParserPool{},
	}
}
