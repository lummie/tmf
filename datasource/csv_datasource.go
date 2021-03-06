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
	reader  io.Reader
	columns []string
	rows    int
	counted bool
}

// NewCsvDataSource create a new data source for a csv file
func NewCsvDataSource(reader io.Reader) *CsvDatasource {
	c := CsvDatasource{
		reader:  reader,
		columns: make([]string, 0),
		counted: false,
		rows:    0,
	}
	return &c
}

// String returns a formatted information about the data source
func (d *CsvDatasource) String() string {
	return fmt.Sprintf("cols:%#v rows:%d", d.columns, d.rows)
}

// ReadHeadersFromRow allows the columns to be initialized by reading a specific row in the input
func (d *CsvDatasource) ReadHeadersFromRow(row uint) error {
	s := bufio.NewScanner(d.reader)
	for i := uint(0); i < row; i++ {
		if !s.Scan() {
			return ErrEOF
		}
	}
	d.columns = strings.Split(s.Text(), ",")
	return nil
}

// Columns returns the list of columns in the data source
func (d *CsvDatasource) Columns() []string {
	return d.columns
}

// UpdateColumns allows setting of the columns for the data source
func (d *CsvDatasource) UpdateColumns(cols []string) {
	d.columns = cols
}

// Rows returns the number of rows in the datasource seperated by CR
func (d *CsvDatasource) Rows() (int, error) {
	if !d.counted {
		c, err := ioext.CountLines(d.reader, []byte{'\n'})
		if err != nil {
			return 0, err
		}
		d.rows, d.counted = c, true
		return c, nil
	}
	return d.rows, nil
}
