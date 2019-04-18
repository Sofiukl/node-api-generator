package apigen

var qIDStr string
var paramStr string

// WriteToConstantFile - This writes the constants
func WriteToConstantFile(apimodel API) {
	prepareQueryID(apimodel)
	prepareConstantStr(apimodel)
	fileContent := qIDStr + paramStr
	ReplaceFileContent(apimodel.Methods.Detail.FileName.ConstName, "#Replace#", fileContent)
}

func prepareQueryID(apimodel API) {
	qIDStr = ""
	dataAPIName := apimodel.Methods.Detail.DataAPIConfig.DataAPIName
	qIDStr = `QueryID.` + dataAPIName + `= '` + dataAPIName + `';`
}
func prepareConstantStr(apimodel API) {
	paramStr = ""
	Accepts := apimodel.Methods.Detail.LbConfig.Accepts
	for _, Accept := range Accepts {
		paramStr +=
			`
APIVariable.` + Accept.Arg + " = '" + Accept.Arg + `';
`
	}
}
