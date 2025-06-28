package utils

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

func PaginationParams(c echo.Context, defaultPageSize, maxPageSize int) (int, int, error) {
	var err error
	page, limit := 1, defaultPageSize
	if len(c.QueryParam("page")) != 0 {
		page, err = strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			return 0, 0, fmt.Errorf("invalid 'page' query param")
		}
		if page <= 0 {
			page = 1
		}
	}

	if len(c.QueryParam("page_size")) != 0 {
		limit, err = strconv.Atoi(c.QueryParam("page_size"))
		if err != nil {
			return 0, 0, fmt.Errorf("invalid 'page_size' query param")
		}
		if limit > maxPageSize {
			limit = maxPageSize
		}
	}

	return page, limit, nil
}

func RequestQueryParamToStruct(c echo.Context, keys []string, dest interface{}) error {
	data := map[string]string{}
	for _, key := range keys {
		val := c.QueryParam(key)
		if len(val) > 0 {
			data[key] = val
		}
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           dest,
	})
	if err != nil {
		return err
	}

	return decoder.Decode(data)
}
