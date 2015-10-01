package readline

import "io"

type Instance struct {
	t *Terminal
	o *Operation
}

type Config struct {
	Prompt       string
	HistoryFile  string
	AutoComplete AutoCompleter
	VimMode      bool
	Stdout       io.Writer
	Stderr       io.Writer

	inited bool
}

func (c *Config) Init() error {
	if c.inited {
		return nil
	}
	c.inited = true
	if c.Stdout == nil {
		c.Stdout = Stdout
	}
	if c.Stderr == nil {
		c.Stderr = Stderr
	}
	return nil
}

func NewEx(cfg *Config) (*Instance, error) {
	t, err := NewTerminal(cfg)
	if err != nil {
		return nil, err
	}
	rl := t.Readline()
	return &Instance{
		t: t,
		o: rl,
	}, nil
}

func New(prompt string) (*Instance, error) {
	return NewEx(&Config{Prompt: prompt})
}

func (i *Instance) SetPrompt(s string) {
	i.o.SetPrompt(s)
}

func (i *Instance) Stdout() io.Writer {
	return i.o.Stdout()
}

func (i *Instance) Stderr() io.Writer {
	return i.o.Stderr()
}

func (i *Instance) SetVimMode(on bool) {
	i.o.SetVimMode(on)
}

func (i *Instance) ReadPassword(prompt string) ([]byte, error) {
	return i.o.Password(prompt)
}

func (i *Instance) Readline() (string, error) {
	return i.o.String()
}

func (i *Instance) ReadSlice() ([]byte, error) {
	return i.o.Slice()
}

func (i *Instance) Close() error {
	if err := i.t.Close(); err != nil {
		return err
	}
	i.o.Close()
	return nil
}
