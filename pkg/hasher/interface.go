package hasher

type Hasher interface {
	GetHash(msg string) (string, error)
}
