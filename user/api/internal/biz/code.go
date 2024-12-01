package biz

const OK = 200

var (
	DBError = NewError(10000, "database error")
)
