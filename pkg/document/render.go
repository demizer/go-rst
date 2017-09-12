package document

type Renderer interface {
	Bytes() ([]byte, error) // Render to bytes
}
