package callmap

type Callmap struct{}

func (c *Callmap) AddFile(path string) error {
	_, err := newFile(path)
	return err
}
