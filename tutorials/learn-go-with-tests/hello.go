package learning_tests

func greeting(name, language string) string {
	if name == "" {
		name = "Tom"
	}
	if language == "" {
		language = "Chinese"
	}
	return "hello, " + name + " in " + language
}
