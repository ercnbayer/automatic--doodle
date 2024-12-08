package factory

type Encryption interface {
	GenerateSalt(int) string
	HashPassword(string, string) string
}
