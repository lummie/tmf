package datasource_test

import (
	"testing"
	"os"
	. "github.com/lummie/assert"
	"github.com/lummie/tmf/datasource"
)

func TestNewCsvDataSource(t *testing.T) {
	f, err := os.Open("testfiles/SalesJan2009.csv")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	ds := datasource.NewCsvDataSource(f)
	expected := []string{}
	Assert(t, ds.Columns(), EqualDeep, expected, "Columns should be empty")
}

func TestCsvDatasource_ReadHeadersFromRow(t *testing.T) {
	f, err := os.Open("testfiles/SalesJan2009.csv")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	ds := datasource.NewCsvDataSource(f)
	if err := ds.ReadHeadersFromRow(1); err != nil {
		t.Error(err)
	}
	expected := []string{"Transaction_date", "Product", "Price", "Payment_Type", "Name", "City", "State", "Country", "Account_Created", "Last_Login", "Latitude", "Longitude"}
	Assert(t, ds.Columns(), EqualDeep, expected, "Columns do not match")
}

func TestCsvDatasource_ReadHeadersFromRowInvalid(t *testing.T) {
	f, err := os.Open("testfiles/SalesJan2009.csv")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	ds := datasource.NewCsvDataSource(f)
	err = ds.ReadHeadersFromRow(1001)
	Assert(t, err, EqualDeep, datasource.ErrEOF, "Expected an ErrEOF error")
}

func TestCsvDatasource_Columns(t *testing.T) {
	f, err := os.Open("testfiles/SalesJan2009.csv")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	ds := datasource.NewCsvDataSource(f)
	expected := []string{}
	Assert(t, ds.Columns(), EqualDeep, expected, "Row Count does not match")
}

func TestCsvDatasource_Rows(t *testing.T) {
	f, err := os.Open("testfiles/SalesJan2009.csv")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	ds := datasource.NewCsvDataSource(f)
	expected := 999
	rows, err := ds.Rows()
	if err != nil {
		t.Error(err)
	}
	Assert(t, rows, EqualInt, expected, "Row Count does not match")
}