package name

type Name interface {
	Name(string) (string, error)
	Key(string) (string, error)
	OtherNames(string) ([]string, error)
}
