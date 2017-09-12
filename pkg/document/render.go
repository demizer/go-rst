package document

type Renderer interface {
	Bytes() (error, []byte) // Render to bytes
}
