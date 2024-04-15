package service

func Deposite() {
	// check user exist
	// check amount is valid
	// insert activity
	// insert or update balance
	// maybe send email
	// return balance
}

func Withdraw() {
	// check user exist
	// check balance
	// check amount is valid
	// insert activity
	// update balance
	// maybe send email
	// return balance
}

func Transfer() {
	// check from and to user.id is exist
	// check from.user.balance is valid
	// check transfer amount is valid
	// insert activity(from and to)
	// insert transfer
	// insert or update balance(from and to)
	// maybe send email to to_user
	// return from_user.balance
}

func ListActivities() {
	// get a user's activities
}

func ListTransfers() {
	// get a user's transfer history list
}
