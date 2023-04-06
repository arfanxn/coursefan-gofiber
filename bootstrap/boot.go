package bootstrap

// Boot will bootstraps everything that needs to be bootstrapped
func Boot() error {
	errs := []error{
		ENV(),
		Logger(),
		FileSystem(),
		Request(),
	}
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
