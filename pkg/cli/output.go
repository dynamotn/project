package cli

import (
	"io"

	"github.com/gosuri/uitable"
)

func EncodeTable(out io.Writer, table *uitable.Table) error {
	raw := table.Bytes()
	raw = append(raw, []byte("\n")...)
	_, err := out.Write(raw)
	if err != nil {
		Error("Unable to write table output", err)
		return err
	}
	return nil
}
