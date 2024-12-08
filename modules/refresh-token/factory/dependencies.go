package factory

type Encryption interface {
	GenerateSalt(int) string
}
