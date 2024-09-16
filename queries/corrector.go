package queries

import (
	"fmt"
	"os"
	"strings"
)

func IsShallowSyntaxCorrect(s Splitter) []error {
	chunks := s.Chunks()
	errs := make([]error, 0)

	// minimum chunks should be 6. join not supported yet.
	//		SELECT: The keyword used to specify the columns to be retrieved.
	//		FROM: The keyword used to specify the table(s) from which data will be retrieved.
	//		WHERE: The keyword used to filter the results based on conditions.
	//		ORDER BY: The keyword used to sort the results.
	//		LIMIT: The keyword used to limit the number of rows returned.
	if len(chunks) < 6 {
		errs = append(errs, InvalidNumberOfChunks)
	}

	if strings.ToLower(chunks[0]) != "select" {
		errs = append(errs, InvalidSelectChunk)
	}

	if strings.ToLower(chunks[2]) != "from" {
		errs = append(errs, InvalidFromChunk)
	}

	splitPath := strings.Split(chunks[3], ":")
	if len(splitPath) != 2 {
		errs = append(errs, InvalidFilePathChunk)
		return errs
	}

	if splitPath[0] != "path" {
		errs = append(errs, InvalidFilePathChunk)
		return errs
	}

	path := splitPath[1]
	stat, err := os.Stat(path)
	if err != nil {
		errs = append(errs, fmt.Errorf("file path %s does not exist: %w", path, InvalidFilePath))
	}

	nameSplit := strings.Split(stat.Name(), ".")
	// only .csv supported for now
	if nameSplit[1] != "csv" {
		errs = append(errs, fmt.Errorf("file %s is not a csv file or it does not have .csv extension:  %w",
			path, InvalidFilePath))
	}

	if strings.ToLower(chunks[4]) != "as" {
		errs = append(errs, InvalidAsChunk)
	}

	return errs
}
