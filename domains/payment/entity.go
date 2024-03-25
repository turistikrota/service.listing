package payment

type Channel string

const (
	ChannelBankTransfer Channel = "bank_transfer"
	ChannelDebitCard    Channel = "debit_card"
	ChannelPapara       Channel = "papara"
	ChannelAtTheDoor    Channel = "at_the_door"
)
