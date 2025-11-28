package l3

type emailStatus int

const (
	emailBounced emailStatus = iota
	emailInvalid
	emailDelivered
	emailOpened
)
