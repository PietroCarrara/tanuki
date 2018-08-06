package parser

type Program struct {
	Statements []Statement
}

func parseProgram() Program {

	res := Program{}

	for {
		for isNewLine(next()) {
			// Skipping separators...
		}

		if isEOF(current()) {
			break
		}

		res.Statements = append(res.Statements, parseStatement())
	}

	return res
}

func (p Program) String() string {

	res := ""

	for _, s := range p.Statements {
		res += s.String() + "\n"
	}

	return res
}
