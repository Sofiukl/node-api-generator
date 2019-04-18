package apigen

import "fmt"

//GenerateAPI - This generates the API
func GenerateAPI(apimodel API) {
	apiType := apimodel.Methods.Detail.Type
	switch apiType {
	case "GET":
		generateGetAPI(apimodel)
		break
	case "POST":
		generatePostAPI(apimodel)
		break
	case "PATCH":
		generatePatchAPI(apimodel)
		break
	case "DELETE":
		generateDeleteAPI(apimodel)
		break
	}
}

func generateGetAPI(apimodel API) {
	WriteToJSONFile(apimodel)
	WriteToModelFile(apimodel)
	WriteToLibFile(apimodel)
	WriteToConstantFile(apimodel)
}

func generatePostAPI(apimodel API) {
	WriteToJSONFile(apimodel)
	WriteToModelFile(apimodel)
	WritePostAPIToLib(apimodel)
	WriteToConstantFile(apimodel)
}

func generatePatchAPI(apimodel API) {
	fmt.Println("inside patch creator")
	WriteToJSONFile(apimodel)
	WriteToModelFile(apimodel)
	WritePatchAPIToLib(apimodel)
	WriteToConstantFile(apimodel)
}

func generateDeleteAPI(apimodel API) {
	WriteToJSONFile(apimodel)
	WriteToModelFile(apimodel)
	WriteToConstantFile(apimodel)
}
