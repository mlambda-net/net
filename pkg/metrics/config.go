package metrics


type Configuration struct {

	App struct {
		Name    string
		Version string
		Env 	string
		Path  string
	}
	Metric struct {
		Namespace string
		SubSystem string
	}
}
