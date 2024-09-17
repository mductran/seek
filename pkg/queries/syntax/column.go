package syntax

const ColumnType = "column"

// column hold list of columns names
type column struct {
	columns []string
}

type Column interface {
	HasColumn(colName string) int
}

// HasColumn returns the column's index if the csv file has a column with name, and -1 otherwise
func (c *column) HasColumn(name string) int {
	for i, cl := range c.columns {
		if cl == name {
			return i
		}
	}

	return -1
}

// Type returns the type of data in column
func (c *column) Type() string {
	return ColumnType
}

func NewColumn(colNames []string) Column {
	return &column{
		columns: colNames,
	}
}
