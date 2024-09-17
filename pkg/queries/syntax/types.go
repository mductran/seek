package syntax

const ColumnType = "column"
const FileDbType = "fileDb"

type SyntaxType interface {
	Type() string
}
