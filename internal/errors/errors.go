package errors

const (
	ErrInvalid  = "Invalid hostname"
	ErrNoNSRec  = "has no nameserver record"
	ErrNoTarget = "No input target provided"
	ErrNoValid  = "File has no valid hostnames"
	ErrPattern  = "Invalid regExp pattern"
	ErrResolve  = "unable to resolve target hostname"
	ErrUnknown  = "Something error"
	ErrNotVuln  = "target not known to be vulnerable"
	ErrVerify   = "cannot verifying target hostname"
	ErrFinger   = "DNS unknown when fingerprinting"
)
