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
	Accepts := apimodel.Methods.Detail.DataAPIConfig.Accepts
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
	logMsg := apimodel.ModelName + "." + apimodel.Methods.Detail.Name + "- Error - ${error.message}"
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
			logger.log('error',` + "`" + logMsg + "`);" +
		`	
			return Promise.resolve(error);
		}`
	bodyStr += validationElseStr
	bodyStr += `
};`
	fmt.Println("body string: " + bodyStr)
	libFn += bodyStr
}

func prepareLibBodyContent(apimodel API) string {
	logMsg := apimodel.ModelName + "." + apimodel.Methods.Detail.Name + "- Error - Unable to process"
	bc := `  
			let command = [];
		`
	for _, Accept := range apimodel.Methods.Detail.DataAPIConfig.Accepts {
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
			if(response && response.result.length > 0) {
				return response;
			} else {
				logger.log('error','` + logMsg + `');
			}
			
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
	Accepts := apimodel.Methods.Detail.DataAPIConfig.Accepts
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
