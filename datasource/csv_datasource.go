package datasource

import (
	"fmt"
	"io"
	"bufio"
	"errors"
	"strings"
	"github.com/lummie/tmf/ioext"
)

var ErrEOF = errors.New("EOF reached when reading")






type CsvDatasource struct {
	reader io.Reader
	columns []string
	rows int
	counted bool
}


// NewCsvDataSource create a new data source for a csv file
func NewCsvDataSource(reader io.Reader) *CsvDatasource {
	c := CsvDatasource{
		reader:reader,
		columns : make([]string,0),
		counted : false,
		rows : 0,
	}
	return &c
}

// String returns a formatted information about the data source
func(d *CsvDatasource) String() string {
	return fmt.Sprintf("cols:%#v rows:%d", d.columns, d.rows)
}

func(d *CsvDatasource) ReadHeadersFromRow(row uint) error {
	s := bufio.NewScanner(d.reader)
	for i := uint(0); i < row; i++ {
		if !s.Scan() {
			return ErrEOF
		}
	}

	d.columns = strings.Split(s.Text(),",")

	return nil
}

func(d *CsvDatasource) Columns() []string {
	return d.columns
}

func(d *CsvDatasource) Rows() (int, error) {
	return ioext.CountLines(d.reader, []byte{'\n'})
}