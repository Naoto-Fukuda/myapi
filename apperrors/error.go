package apperrors

type MyAppError struct {
	ErrCode	// (フィールド名を省略した場合、型名がそのままフィールド名になる)
	Message string
	Err error `json:"-"`
}

func (myErr *MyAppError) Error() string {
	// myErr.Err.Error()を呼び出すと、エラーメッセージが文字列として返される。
	// ラップしたエラーの内容を表示するために、myErr.Err.Error()という書き方をしている
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
