package logger

type Dummy struct {
}

func (d *Dummy) Log(msg string) error {
	return nil
}
