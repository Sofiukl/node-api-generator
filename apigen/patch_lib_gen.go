package apigen

import (
	"fmt"
)

var patchLibFn string

// WritePatchAPIToLib - This writes the patch api libary defination into lib file
func WritePatchAPIToLib(apimodel API) {
	patchLibFn = ""
	createPatchAPILibFuncStart(apimodel)
	createPatchAPILibFuncParams(apimodel)
	createPatchAPILibFuncBody(apimodel)
	ReplaceFileContent(apimodel.Methods.Detail.FileName.LibName, "#Replace#", patchLibFn)
}

func createPatchAPILibFuncStart(apimodel API) {
	methodName := apimodel.Methods.Detail.Name
	patchLibFn += "async " + methodName
}

func createPatchAPILibFuncParams(apimodel API) {
	paramStr := "(input)"
	patchLibFn += paramStr
}

func createPatchAPILibFuncBody(apimodel API) {
	validationMsg := apimodel.Methods.Detail.Name + " - required fields are undefined"
	validationIfStr := `
	if (input && input.PINProjectID) {
		`
	validationElseStr := `
	} else {
		return util.buildResponse({error: true,
		message: '` + validationMsg + `'});
	}
	`
	logMsg := apimodel.ModelName + "." + apimodel.Methods.Detail.Name + "- Error - ${error.message}"
	bodyContent := preparePatchAPILibBodyContent(apimodel)
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
	patchLibFn += bodyStr
}

func preparePatchAPILibBodyContent(apimodel API) string {
	preProcessStrs := DoPreProcess(apimodel)
	postProcessStrs := DoPostProcess(apimodel)
	preProcessStr := ""
	for _, preStr := range preProcessStrs {
		preProcessStr += `
			` + preStr + `
		`
	}

	bc := preProcessStr +
		`  
			let fields = input.UpdateField.split(',');
			let values = input.UpdateFieldValue.split(',');
			const currentUser = util.getLoggedInUser();
			let arrayUpdateFields = [];
			fields.forEach((field, index) => {
				const f = ` + "`${field}`;" + `
				const v = ` + "`${values[index]}`;" + `
				const r = f+v;
				arrayUpdateFields.push(r);
			});
			let command = [];
		`

	for _, Accept := range apimodel.Methods.Detail.DataAPIConfig.Accepts {
		key := "APIVariable." + Accept.Arg
		value := Accept.Arg
		objStr := `{key: ` + key + `, value: input.` + value + `}`
		commandPushStr := `
			command.push(` + objStr + `);
		`
		bc += commandPushStr
	}
	bc += `
			command.push({key: APIVariable.strUpdateColumns, value: arrayUpdateFields.join()});
			command.push({key: APIVariable.ModifyBy, value: currentUser.PWUserID});
			
			const response = await dataAPI.getDataAPIProjectDB(QueryID.` + apimodel.Methods.Detail.DataAPIConfig.DataAPIName + `, input.PINProjectID, command);
			if(response && response.result[0].affectedRows != 1) {
				`
	bc += `
				return Promise.resolve(response);
			}`
	for _, postStr := range postProcessStrs {
		bc += `
			` + postStr + `
		`
	}
	bc += `
			return Promise.resolve({
				result: [{
				  error: false,`
	bc += apimodel.Methods.Detail.DataAPIConfig.PrimaryKey + `: response.result
				}]
			});
		`
	return bc
}

// PreparePatchPILibFuncCalling - This return the function calling string
func PreparePatchPILibFuncCalling(apimodel API) string {
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
