package simplate

// AddFuncMap register a func in the template.
func AddFuncMap(key string, fn interface{}) error {
	simplateTplFuncMap[key] = fn
	return nil
}
