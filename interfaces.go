package healthcheckmodule

type ITestableConnection interface {
	Connect() error
}
