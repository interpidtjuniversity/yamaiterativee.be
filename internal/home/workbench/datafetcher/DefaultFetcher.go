package datafetcher

type DefaultFetcher struct {
	Fetcher
}

func (m DefaultFetcher) Fetch(userName string, limit int) ([]byte, error) {
	return []byte("[]"), nil
}

