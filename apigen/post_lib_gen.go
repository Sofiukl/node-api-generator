package apigen

import (
	"fmt"
)

var postLibFn string

// WritePostAPIToLib - This writes the post api libary defination into lib file
func WritePostAPIToLib(apimodel API) {
	postLibFn = ""
	createPostAPILibFuncStart(apimodel)
	createPostAPILibFuncParams(apimodel)
	createPostAPILibFuncBody(apimodel)
	ReplaceFileContent(apimodel.Methods.Detail.FileName.LibName, "#Replace#", postLibFn)
}

func createPostAPILibFuncStart(apimodel API) {
	methodName := apimodel.Methods.Detail.Name
	postLibFn += "async " + methodName
}

func createPostAPILibFuncParams(apimodel API) {
	paramStr := "(input)"
	postLibFn += paramStr
}

func createPostAPILibFuncBody(apimodel API) {
	tableName := apimodel.Methods.Detail.DataAPIConfig.TableName
	primaryKey := apimodel.Methods.Detail.DataAPIConfig.PrimaryKey
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
	sequenceGenStr := `	const ` + primaryKey + ` = await dataAPI.executeSequenceGenerator(TableID.` + tableName + `,1);`
	logMsg := apimodel.ModelName + "." + apimodel.Methods.Detail.Name + "- Error - ${error.message}"
	bodyContent := preparePostAPILibBodyContent(apimodel)
	bodyStr := ` { 
		`
	bodyStr += validationIfStr
	bodyStr += `
		try { 
		`
	bodyStr += sequenceGenStr
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
	postLibFn += bodyStr
}

func preparePostAPILibBodyContent(apimodel API) string {
	preProcessStrs := DoPreProcess(apimodel)
	postProcessStrs := DoPostProcess(apimodel)
	preProcessStr := ""
	for _, preStr := range preProcessStrs {
		preProcessStr += `
			` + preStr + `
		`
	}
	primaryKey := apimodel.Methods.Detail.DataAPIConfig.PrimaryKey
	keyStr := "APIVariable." + primaryKey
	valueStr := primaryKey + ".result"
	primaryKeyObjStr := `{key: ` + keyStr + `, value: ` + valueStr + `}`
	primaryKeyCommandPushStr := `
			command.push(` + primaryKeyObjStr + `);
		`
	bc := preProcessStr +
		`  
			let command = [];
		`
	for _, Accept := range apimodel.Methods.Detail.DataAPIConfig.Accepts {
		key := "APIVariable." + Accept.Arg
		value := "input." + Accept.Arg
		objStr := `{key: ` + key + `, value: ` + value + `}`
		commandPushStr := `
			command.push(` + objStr + `);
		`
		bc += commandPushStr
	}
	bc += primaryKeyCommandPushStr
	bc += `
			const response = await dataAPI.getDataAPIProjectDB(QueryID.` + apimodel.Methods.Detail.DataAPIConfig.DataAPIName + `, PINProjectID, command);
			if(response && response.result[0].affectedRows != 1) {
				return Promise.resolve(result);
			}`
	for _, postStr := range postProcessStrs {
		bc += `
			` + postStr + `
			`
	}
	bc += `
			return Promise.resolve({
				result: [{
					error: false,
					`
	bc += apimodel.Methods.Detail.DataAPIConfig.PrimaryKey + `: ` + valueStr
	bc += `
				}]
			});
		`
	return bc
}

// PreparePostAPILibFuncCalling - This return the function calling string
func PreparePostAPILibFuncCalling(apimodel API) string {
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
