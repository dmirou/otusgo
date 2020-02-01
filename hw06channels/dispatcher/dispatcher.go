package dispatcher

type Task func() error

func Run(task []Task, N int, M int) error {
	return nil
}
