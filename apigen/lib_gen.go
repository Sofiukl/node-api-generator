package apigen

import (
	"fmt"
)

var libFn string

// WriteToLibFile - This writes the api libary defination into lib file
func WriteToLibFile(apimodel API) {
	createLibFuncStart(apimodel)
	createLibFuncParams(apimodel)
	createLibFuncBody(apimodel)
	ReplaceFileContent(apimodel.Methods.Detail.FileName.LibName, "#Replace#", libFn)
}

func createLibFuncStart(apimodel API) {
	methodName := apimodel.Methods.Detail.Name
	libFn += "async " + methodName
}

func createLibFuncParams(apimodel API) {
	paramStr := "("
	Accepts := apimodel.Methods.Detail.LbConfig.Accepts
	for i, Accept := range Accepts {
		if i == 0 {
			paramStr += Accept.Arg
		} else {
			paramStr += ", " + Accept.Arg
		}
	}
	paramStr += ")"
	libFn += paramStr
}

func createLibFuncBody(apimodel API) {
	validationMsg := apimodel.Methods.Detail.Name + " - required fields are undefined"
	validationIfStr := `
	if (PINProjectID) {
		`
	validationElseStr := `
	} else {
		return util.buildResponse({error: true,
		message: '` + validationMsg + `'});
	}
	`

	bodyContent := prepareLibBodyContent(apimodel)
	bodyStr := ` { 
		`
	bodyStr += validationIfStr
	bodyStr += `
		try { 
		`
	bodyStr += bodyContent
	bodyStr += `
		} catch (error) {
			return error;
		}
		`
	bodyStr += validationElseStr
	bodyStr += `
};`
	fmt.Println("body string: " + bodyStr)
	libFn += bodyStr
}

func prepareLibBodyContent(apimodel API) string {
	bc := `  
			let command = [];
		`
	for _, Accept := range apimodel.Methods.Detail.LbConfig.Accepts {
		key := "APIVariable." + Accept.Arg
		value := Accept.Arg
		objStr := `{key: ` + key + `, value: ` + value + `}`
		commandPushStr := `
			command.push(` + objStr + `);
		`
		bc += commandPushStr
	}
	bc += `
			const response = await dataAPI.getDataAPIProjectDB(QueryID.` + apimodel.Methods.Detail.DataAPIConfig.DataAPIName + `, PINProjectID, command);
			return response;
		`
	return bc
}

// PrepareLibFuncCalling - This return the function calling string
func PrepareLibFuncCalling(apimodel API) string {
	fnCallingStr := ""
	libName := apimodel.ModelName
	methodName := apimodel.Methods.Detail.Name
	fnCallingStr = libName + "." + methodName

	//
	paramStr := "("
	Accepts := apimodel.Methods.Detail.LbConfig.Accepts
	for i, Accept := range Accepts {
		if i == 0 {
			paramStr += Accept.Arg
		} else {
			paramStr += ", " + Accept.Arg
		}
	}
	paramStr += ")"
	fnCallingStr += paramStr

	return fnCallingStr
}
