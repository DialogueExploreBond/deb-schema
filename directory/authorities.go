package directory

// This enumeration holds all the known authorities for the current application.
const (
	// AuthAuthenticated indicates that the user is logged in
	// and configured on the remote directory
	AuthAuthenticated = "AUTHENTICATED"

	// AuthRoot indicates that user is a super admin in the current context
	AuthRoot = "ROOT"

	// AuthBotOperator indicates that the user belongs to the Bot Operators group
	AuthBotOperator = "BOT_OPERATOR"

	// AuthEnabledBotOperator indicates that the user is correctly configured to operate as
	// Bot Operator (has a Telegram ID and an active status)
	AuthEnabledBotOperator = "ENABLED_BOT_OPERATOR"

	// AuthBotServiceAdmin indicates that the user belongs to the Bot Service Administrators group
	AuthBotServiceAdmin = "BOT_SERVICE_ADMIN"

	// AuthEnrolledIn2Sv indicates that the user is enrolled in 2 step verification
	AuthEnrolledIn2Sv = "2FA_ENROLLED"
)
