package schema

type QueryParams struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Author string `query:"author"`
}
