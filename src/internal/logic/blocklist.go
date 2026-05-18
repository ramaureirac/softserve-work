package logic

type BlockList struct {
	Items map[string]map[string]struct{}
}

func (b *BlockList) Search(host string, query string) (bool, error) {
	queries, ok := b.Items[host]
	if ok {
		_, ok = queries[query]
		if ok {
			return true, nil
		}
	}
	return false, nil
}

func (b *BlockList) Add(host string, query string) error {
	_, ok := b.Items[host]
	if !ok {
		b.Items[host] = make(map[string]struct{})
	}
	b.Items[host][query] = struct{}{}
	return nil
}
