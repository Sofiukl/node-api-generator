package apigen

//DoPreProcess - This function pre process the input data before calling data APIs
func DoPreProcess(apimodel API) []string {
	preProcess := apimodel.Methods.Detail.PreProcess
	if isPostProcessEnabled(preProcess) {
		return apimodel.Methods.Detail.PreProcess
	}
	return []string{}
}

func isPreProcessEnabled(preProcess []string) bool {
	if len(preProcess) != 0 {
		return true
	}
	return false
}
