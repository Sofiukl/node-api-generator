package apigen

//DoPostProcess - This function post process the response returned from the database
func DoPostProcess(apimodel API) []string {
	postProcess := apimodel.Methods.Detail.PostProcess
	if isPostProcessEnabled(postProcess) {
		return apimodel.Methods.Detail.PostProcess
	}
	return []string{}
}

func isPostProcessEnabled(postProcess []string) bool {
	if len(postProcess) != 0 {
		return true
	}
	return false
}
