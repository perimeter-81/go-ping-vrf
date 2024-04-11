package ping

// Returns the length of an ICMP message.
func (p *Pinger) getMessageLength() int {
	return 0
}

// Attempts to match the ID of an ICMP packet.
func (p *Pinger) matchID(ID int) bool {
	return true
}
