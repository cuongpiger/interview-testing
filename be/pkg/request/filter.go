package request

import (
	"errors"
	"strconv"
	"strings"
)

type Filter struct {
	Field    string
	Operator string
	Value    any
}

func ToFilter(cmd string) (*Filter, error) {
	commands := strings.Split(cmd, ":")
	if len(commands) != 3 {
		return nil, errors.New("invalid filter command")
	}

	filter := new(Filter)
	switch commands[1] {
	case "eq":
		return filter.eq(commands[0], commands[2])
	case "in":
		return filter.in(commands[0], commands[2])
	case "like":
		return filter.like(commands[0], commands[2])
	case "gte":
		return filter.gte(commands[0], commands[2])
	case "lte":
		return filter.lte(commands[0], commands[2])
	}

	return nil, errors.New("invalid filter operator")
}

func (s *Filter) GetQuery(opt string) string {
	switch opt {
	case "postgres":
		return s.postgresQuery()
	}

	return ""
}

func (s *Filter) postgresQuery() string {
	switch s.Operator {
	case "eq":
		return s.Field + " = ?"
	case "in":
		return s.Field + " IN ?"
	case "like":
		return s.Field + " ILIKE ?"
	case "gte":
		return s.Field + " >= ?"
	case "lte":
		return s.Field + " <= ?"
	}

	return ""
}

func (s *Filter) eq(field string, data any) (*Filter, error) {
	s.Field = field
	s.Operator = "eq"
	ds := data.(string)

	if nmb, err := strconv.ParseInt(ds, 10, 32); err == nil {
		s.Value = int(nmb)
		return s, nil
	}

	if nmb, err := strconv.ParseFloat(ds, 32); err == nil {
		s.Value = float32(nmb)
		return s, nil
	}

	s.Value = ds
	return s, nil
}

func (s *Filter) in(field string, data string) (*Filter, error) {
	splitData := strings.Split(data, ",")
	if len(splitData) < 1 {
		return nil, errors.New("invalid filter value")
	}

	s.Field = field
	s.Operator = "in"
	if _, err := strconv.ParseInt(splitData[0], 10, 64); err == nil {
		s.Value = func(data []string) []int {
			res := make([]int, 0)
			for _, v := range data {
				if nmb, e := strconv.ParseInt(v, 10, 32); e == nil {
					res = append(res, int(nmb))
				}
			}

			return res
		}(splitData)

		return s, nil
	}

	s.Value = splitData
	return s, nil
}

func (s *Filter) like(field string, data string) (*Filter, error) {
	s.Field = field
	s.Operator = "like"
	s.Value = data
	return s, nil
}

func (s *Filter) gte(field string, data string) (*Filter, error) {
	s.Field = field
	s.Operator = "gte"
	s.Value = data
	return s, nil
}

func (s *Filter) lte(field string, data string) (*Filter, error) {
	s.Field = field
	s.Operator = "lte"
	s.Value = data
	return s, nil
}
