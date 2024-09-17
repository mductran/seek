package queries

import (
	"fmt"
	"os"
	"seeker/queries/operators"
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

	whereClause := chunks[6:]
	if len(whereClause) != 0 {
		if len(whereClause) < 4 {
			errs = append(errs, fmt.Errorf("Expected at least a single condition for WHERE clause but got something else: %w", InvalidWhereClause))
			return errs
		}

		where := whereClause[0]
		operator := whereClause[2]
		value := whereClause[3]

		if strings.ToLower(where) != "where" {
			errs = append(errs, fmt.Errorf("expected WHERE, got %s: %w", where, InvalidWhereClause))
		}
		if err := checkOperator(operator); err != nil {
			errs = append(errs, err)
		}
		if err := checkValue(value); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func checkOperator(op string) error {
	var found bool
	for _, v := range operators.Operators {
		if v == op {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("expected one of valid operators: %s, got %s: %w",
			strings.Join(operators.Operators, ","), op, InvalidWhereClause)
	}

	return nil
}

func checkValue(v string) error {
	if v[0] != '\'' || v[len(v)-1] != '\'' {
		return fmt.Errorf("invalid string comparison value. Comparison values should be in enclosed in single quotes: %w", InvalidValueChunk)
	}

	return nil
}
