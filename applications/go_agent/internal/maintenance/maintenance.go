package maintenance

// Log contains a maintenance action recorded for a server.
type Log struct {
	Action      string
	Description string
	PerformedBy string
}
