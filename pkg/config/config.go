package sandyconfig

type Config struct {
	Name         string
	Run          []Command
	BuildDebug   []Command
	BuildRelease []Command
	InstallDeps  []Command
	OnChange     []Command
	Command      []Command
}

type Command struct {
	Name string
	Cmd  string
	Args []string
}
