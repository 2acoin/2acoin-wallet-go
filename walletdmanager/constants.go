// Copyright (c) 2018, The TurtleCoin Developers
// Copyright (c) 2018, 2ACoin Developers
//
// Please see the included LICENSE file for more information.
//

package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in ARMS
	DefaultTransferFee float64 = 0.0005

	logWalletdCurrentSessionFilename     = "2acoin-service-session.log"
	logWalletdAllSessionsFilename        = "2acoin-service.log"
	logxCoindCurrentSessionFilename      = "2ACoind-session.log"
	logxCoindAllSessionsFilename         = "2ACoind.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "2acoin-service"
	xcoindCommandName                    = "2ACoind"
)
