package batch

type batch struct {
	Nodes []node `yaml:"nodes"`
}

func newBatch() *batch {
	_nodes := make([]node, 0)

	_batch := batch{
		Nodes: _nodes,
	}
	return &_batch
}

type batches struct {
	Batches []batch `yaml:"batches"`
}
