package pongo2

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
	openCommentToken = open
	closeCommentToken = close
	templateTagMapping["opencomment"] = open
	templateTagMapping["closecomment"] = close
	return nil
}
