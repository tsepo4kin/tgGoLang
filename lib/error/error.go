package error

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func WrapIfErr(msg string, err error) {
	if err == nil {
		return nil
	}

	return Wrap(msg, err)
}