package queries

import "errors"

var InvalidNumberOfChunks = errors.New("invalid number of chunks. Minimum number of syntax chunks is 6")
var InvalidSelectChunk = errors.New("expected 'select', got something else")
var InvalidAsChunk = errors.New("expected 'as', got something else")
var InvalidFromChunk = errors.New("expected 'from', got something else")
var InvalidFilePathChunk = errors.New("expected 'path:path_to_file' but did not get the path: part")
var InvalidFilePath = errors.New("invalid file path")
var InvalidWhereClause = errors.New("invalid WHERE clause")
var InvalidValueChunk = errors.New("invalid value chunk")
