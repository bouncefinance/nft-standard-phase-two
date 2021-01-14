package msgs

var MsgFlags = map[int]string{
	INVALID_PARAMS: "request params error: ",
}

var MsgReturn = map[int]string{
	SUCCESS:        "ok",
	ERROR_DB_ERROR: "There is an error with db",
}
