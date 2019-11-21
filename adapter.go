package sheru

type adapter interface {
	quote(key string) string
}
