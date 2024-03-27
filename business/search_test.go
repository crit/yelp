package business

import (
	"github.com/crit/yelp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch(t *testing.T) {
	client := yelp.New(yelp.ConfigFromEnv())

	res, err := Search(client, SearchRequest{
		Location: "Idaho Falls",
		Term:     "food",
		Limit:    10,
		Offset:   0,
	})

	require.Nil(t, err)
	assert.Equal(t, 10, len(res.Businesses), "%#v", res.Businesses)
}
