package ex

type Friendly struct {
  msg string
  err error
}

func (f Friendly) Error() string {
  return f.msg
}
func (f Friendly) Fail () error  {
  return f.err
}

type Crashed struct {
  msg string
  error error
}

func (c Crashed) Error() string {
  return c.msg
}

func (c Crashed) Fail () error  {
  return c.error
}

func Friend (msg string, err error ) error  {
  return Friendly{
    msg: msg,
    err: err,
  }
}

func Crash(msg string, err error) error  {
  return Crashed{
    msg:   msg,
    error: err,
  }
}
