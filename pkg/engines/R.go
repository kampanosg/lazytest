package engines

type Runner interface {
	RunCmd(cmd string) (string, error)
}