package datasource


type Datasource interface{
	Columns() []string
	Rows() (int, error)
}