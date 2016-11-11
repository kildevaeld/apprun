package apprun

type Kind int

const (
	NodeKind Kind = iota + 1
)

type Config struct {
	Workspace string `mapstructure:"workspace"`
	Remote    string
	Branch    string
	Kind      string
	Command   string
	Args      []string
	Environ   []string
}
