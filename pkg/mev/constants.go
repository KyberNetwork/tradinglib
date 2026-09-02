package mev

// supported builders in flashbots
// ref: https://raw.githubusercontent.com/flashbots/dowg/main/builder-registrations.json
const (
	FlashbotBuilderRegistrationFlashbot      = "flashbots"
	FlashbotBuilderRegistrationF1b1          = "f1b.io"
	FlashbotBuilderRegistrationRsync         = "rsync"
	FlashbotBuilderRegistrationBeaverBuild   = "beaverbuild.org"
	FlashbotBuilderRegistrationBuilder0x69   = "builder0x69"
	FlashbotBuilderRegistrationTitan         = "Titan"
	FlashbotBuilderRegistrationEigenPhi      = "EigenPhi"
	FlashbotBuilderRegistrationBobaBuilder   = "boba-builder"
	FlashbotBuilderRegistrationGambitLabs    = "Gambit Labs"
	FlashbotBuilderRegistrationPayload       = "payload"
	FlashbotBuilderRegistrationLoki          = "Loki"
	FlashbotBuilderRegistrationBuildAI       = "BuildAI"
	FlashbotBuilderRegistrationJetBuilder    = "JetBuilder"
	FlashbotBuilderRegistrationTBuilder      = "tbuilder"
	FlashbotBuilderRegistrationPenguinBuild  = "penguinbuild"
	FlashbotBuilderRegistrationBobTheBuilder = "bobthebuilder"
	FlashbotBuilderRegistrationBTCS          = "BTCS"
	XFlashbotSignatureHeader                 = "X-Flashbots-Signature"
	XBlinkSignatureHeader                    = "X-Blink-Signature"
)

// BomboraMaxRefundPercent is the highest refund percent Bombora accepts.
const BomboraMaxRefundPercent = 99

// UltrasoundMaxRefundPercent is the highest refund percent Ultrasound accepts.
const UltrasoundMaxRefundPercent = 99

// UltrasoundSendBundleRPS and UltrasoundSendBundleBurst are the eth_sendBundle rate
// limit each Ultrasound builder endpoint accepts: 50 requests per second. The limit is
// per endpoint, so a caller that sends to EndpointUltrasoundEU, EndpointUltrasoundUS
// and EndpointUltrasoundJP builds one Client with its own limit for each.
const (
	UltrasoundSendBundleRPS   = 50
	UltrasoundSendBundleBurst = 50
)
