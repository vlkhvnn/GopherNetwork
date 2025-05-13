package storage

import (
	"net/http"
	"strconv"
	"strings"
)

type PaginatedFeedQuery struct {
	Limit  int      `json:"limit" validate:"gte=1,lte=20"`
	Offset int      `json:"offset" validate:"gte=0"`
	Sort   string   `json:"sort" validate:"oneof=asc desc"`
	Tags   []string `json:"tags" validate:"max=5"`
	Search string   `json:"search" validate:"max=100"`
}

func (fq PaginatedFeedQuery) Parse(r *http.Request) (PaginatedFeedQuery, error) {
	qs := r.URL.Query()

	// Parse the limit, offset, and sort query parameters
	limit := qs.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return fq, err
		}
		fq.Limit = l
	}

	offset := qs.Get("offset")
	if offset != "" {
		off, err := strconv.Atoi(offset)
		if err != nil {
			return fq, err
		}
		fq.Offset = off
	}

	sort := qs.Get("sort")
	if sort != "" {
		fq.Sort = sort
	}

	// Parse tags and search query parameters
	tags := qs.Get("tags")
	if tags != "" {
		fq.Tags = strings.Split(tags, ",")
	}

	search := qs.Get("search")
	if search != "" {
		fq.Search = search
	}

	return fq, nil
}
