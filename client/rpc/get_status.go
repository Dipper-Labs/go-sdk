package rpc

func (c *client) GetNodeStatus() (status NodeStatus, err error) {
	s, err := c.rpc.Status()
	if err != nil {
		return
	}

	status.NodeInfo = s.NodeInfo
	status.SyncInfo = s.SyncInfo

	return
}
