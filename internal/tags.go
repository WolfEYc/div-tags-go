package internal

import "github.com/meteran/gnext"

type GetTagsQuery struct {
	gnext.Query
	DivisionId int `form:"divisionId"`
}

type TagList struct {
	MetricNiceName []string `json:"metric_nice_name"`
}

type GetTagsRes struct {
	Tags TagList `json:"tags"`
}

func getAllTagsAndCategories(query *GetTagsQuery) *GetTagsRes {

	return &GetTagsRes{Tags: TagList{MetricNiceName: []string{"oil_tank", "wellbore_pressure"}}}
}
