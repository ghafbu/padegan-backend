package ftppkg

import (
	"fmt"
	"io"
	"time"

	"github.com/jlaffaye/ftp"
)

type Ftp struct {
	serverConn *ftp.ServerConn
	timeout    time.Duration
}

func New() *Ftp {
	return &Ftp{}
}

func (f *Ftp) GetConfig() {

}

func (f *Ftp) SetConfig() {

}

func (f *Ftp) SetDialWithTimeout(timeout time.Duration) *Ftp {
	f.timeout = timeout

	return f
}

func (f *Ftp) Dial(host string, port string) (*Ftp, error) {
	var err error

	if f.timeout == 0 {
		f.timeout = ftp.DefaultDialTimeout
	}

	f.serverConn, err = ftp.Dial(
		fmt.Sprintf("%s:%s", host, port),
		ftp.DialWithTimeout(f.timeout),
	)
	if err != nil {
		return f, err
	}

	return f, nil
}

func (f *Ftp) Login(username string, password string) error {
	err := f.serverConn.Login(username, password)
	if err != nil {
		return err
	}

	return nil
}

func (f *Ftp) Store(path string, data io.Reader) error {
	if err := f.serverConn.Stor(path, data); err != nil {
		return err
	}

	return nil
}

func (f *Ftp) Read(path string) ([]byte, error) {
	var err error
	r, err := f.serverConn.Retr(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (f *Ftp) Quit() error {
	if err := f.serverConn.Quit(); err != nil {
		return err
	}

	return nil
}

func (f *Ftp) Logout() error {
	if err := f.serverConn.Logout(); err != nil {
		return err
	}

	return nil
}

func (f *Ftp) MakeDir(path string) error {
	if err := f.serverConn.MakeDir(path); err != nil {
		//return f, err
		switch err.Error() {
		case fmt.Sprintf("550 %s: File exists", path):
			return nil
		default:
			return err
		}
	}

	return nil
}

func (f *Ftp) ChangeDir(path string) error {
	if err := f.serverConn.ChangeDir(path); err != nil {
		return err
	}

	return nil
}

func (f *Ftp) FileSize(path string) (int64, error) {
	size, err := f.serverConn.FileSize(path)
	if err != nil {
		return 0, err
	}

	return size, nil
}

func (f *Ftp) Delete(path string) error {
	if err := f.serverConn.Delete(path); err != nil {
		return err
	}

	return nil
}

func (f *Ftp) ChangeToParentDir() error {
	if err := f.serverConn.ChangeDirToParent(); err != nil {
		return err
	}

	return nil
}

func (f *Ftp) CurrentDir() (string, error) {
	currentDir, err := f.serverConn.CurrentDir()
	if err != nil {
		return "", err
	}

	return currentDir, nil
}
