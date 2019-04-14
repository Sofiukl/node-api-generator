package apigen

//API - This is api configuration
type API struct {
	ModelName string `json:"model_name,omitempty"`
	Methods   Method `json:"methods,omitempty"`
}

//Method - This is method
type Method struct {
	Detail Detail `json:"detail,omitempty"`
}

//Detail - This is detail
type Detail struct {
	Name          string        `json:"name,omitempty"`
	FileName      FileName      `json:"file_name,omitempty"`
	LbConfig      LbConfig      `json:"lb_config,omitempty"`
	DataAPIConfig DataAPIConfig `json:"data_api_config,omitempty"`
}

//FileName - This is file name
type FileName struct {
	JSONName  string `json:"json_name,omitempty"`
	ModelName string `json:"model_name,omitempty"`
	LibName   string `json:"lib_name,omitempty"`
}

//LbConfig - This is lb config
type LbConfig struct {
	Accepts []Accept   `json:"accepts,omitempty"`
	Returns Returns    `json:"returns,omitempty"`
	HTTP    HTTPConfig `json:"http,omitempty"`
}

//Returns - This is Return config
type Returns struct {
	Arg  string `json:"arg,omitempty"`
	Type string `json:"type,omitempty"`
	Root string `json:"root,omitempty"`
}

//HTTPConfig - This is HttpConfig
type HTTPConfig struct {
	Verb string `json:"verb,omitempty"`
	Path string `json:"path,omitempty"`
}

// Accept - This is accept
type Accept struct {
	Arg      string `json:"arg,omitempty"`
	Type     string `json:"type,omitempty"`
	Required bool   `json:"required,omitempty"`
}

// DataAPIConfig - This is data api cnfig
type DataAPIConfig struct {
	DataAPIName string   `json:"data_api_name,omitempty"`
	Accepts     []Accept `json:"accepts,omitempty"`
}
