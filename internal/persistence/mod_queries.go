package persistence

const (
	insertModQuery = `
insert OR REPLACE into mod_info (
    id
	, name
	, url
	, source_url
	, description
	, class_id
	, game_id
	, versions
	, categories
	, authors
) values (
    :id
	, :name
	, :url
	, :source_url
	, :description
	, :class_id
	, :game_id
	, json(:versions)
	, json(:categories)
	, json(:authors)
)
`
)
