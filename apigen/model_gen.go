package apigen

import (
	"fmt"
)

var modelFn string

// WriteToModelFile - This writes the api model defination into model file
func WriteToModelFile(apimodel API) {
	Accepts := apimodel.Methods.Detail.LbConfig.Accepts
	fmt.Printf("params:  %v/n", Accepts)
	createModelFuncStart(apimodel)
	createModelFuncParams(apimodel)
	createModelFuncBody(apimodel)
	ReplaceFileContent(apimodel.Methods.Detail.FileName.ModelName, "#Replace#", modelFn)
}

func createModelFuncStart(apimodel API) {
	modelName := apimodel.ModelName
	methodName := apimodel.Methods.Detail.Name
	modelFn += modelName + "." + methodName + " = async "
}

func createModelFuncParams(apimodel API) {
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
	modelFn += paramStr
}

func createModelFuncBody(apimodel API) {

	Accepts := apimodel.Methods.Detail.LbConfig.Accepts
	for _, Accept := range Accepts {

	}

	bodyContent := `const libRes = ` + PrepareLibFuncCalling(apimodel) + `;`
	bodyContent += ` if (libRes.error) return libRes;`

	responseStr := `
		error : libRes.result[0].error,
		result : [{

		}]
	`

	bodyStr := ` => { 
	try { 
		`
	bodyStr += bodyContent
	bodyStr += responseStr
	bodyStr += `
		
	} catch (error) {
		return error;
	}
};`
	modelFn += bodyStr
}
