package errors

const (
	ErrInvalid  = "Invalid hostname"
	ErrNoNSRec  = "Has no nameserver record"
	ErrNoTarget = "No input target provided"
	ErrNoValid  = "File has no valid hostname"
	ErrPattern  = "Invalid RegExp pattern"
	ErrResolve  = "Unable to resolve target hostname"
	ErrUnknown  = "Unknown error"
	ErrNotVuln  = "Target is not known to be vulnerable"
	ErrVerify   = "Cannot verify the target hostname"
	ErrFinger   = "DNS unknown when fingerprinting"
)
