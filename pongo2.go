package pongo2

import "fmt"

// Version string
const Version = "6.0.0"

// Must panics, if a Template couldn't successfully parsed. This is how you
// would use it:
//     var baseTemplate = pongo2.Must(pongo2.FromFile("templates/base.html"))
func Must(tpl *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return tpl
}

var (
	openBlockToken     string = "{%"
	closeBlockToken    string = "%}"
	openVariableToken  string = "{{"
	closeVariableToken string = "}}"
	openCommentToken   string = "{#"
	closeCommentToken  string = "#}"
)

// SetBlockTokens changes block tokens
func SetBlockTokens(open, close string) error {
	if err := checkTokens(open, close); err != nil {
		return err
	}
	openBlockToken = open
	closeBlockToken = close
	templateTagMapping["openblock"] = open
	templateTagMapping["closeblock"] = close
	TokenSymbols[11] = open
	TokenSymbols[12] = close
	return nil
}

// SetVariableTokens changes variable tokens
func SetVariableTokens(open, close string) error {
	if err := checkTokens(open, close); err != nil {
		return err
	}
	openVariableToken = open
	closeVariableToken = close
	templateTagMapping["openvariable"] = open
	templateTagMapping["closevariable"] = close
	TokenSymbols[9] = open
	TokenSymbols[10] = close
	return nil
}

// SetCommentTokens changes comment tokens
func SetCommentTokens(open, close string) error {
	if err := checkTokens(open, close); err != nil {
		return err
	}
	openCommentToken = open
	closeCommentToken = close
	templateTagMapping["opencomment"] = open
	templateTagMapping["closecomment"] = close
	return nil
}

func checkTokens(open, close string) error {
	if err := checkToken(open); err != nil {
		return err
	}
	if err := checkToken(close); err != nil {
		return err
	}
	return nil
}

func checkToken(token string) error {
	if len(token) != 2 {
		return fmt.Errorf("Invalid token len(%d) != 2: %s", len(token), token)
	}
	return nil
}
