package atsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPlayers = []Player{
	{
		ID:         "111",
		FullName:   "Alpha (Alpha)",
		Name:       "Alpha",
		NativeName: "阿尔法",
	},
	{
		ID:         "222",
		FullName:   "Beta (Beta)",
		Name:       "Beta",
		NativeName: "贝塔",
	},
	{
		ID:         "112",
		FullName:   "Alpha (Alpha)",
		Name:       "Alpha",
		NativeName: "阿尔法2号",
	},
}

func TestGetByID(t *testing.T) {
	db := NewPlayerDB(testPlayers)
	p := db.FindPlayer("111")
	assert.NotNil(t, p)
	assert.Equal(t, "111", p.ID)
	assert.Equal(t, "Alpha", p.Name)

	q := db.FindPlayer("333")
	assert.Nil(t, q)
}

func TestGetByFullName(t *testing.T) {
	db := NewPlayerDB(testPlayers)
	p := db.FindPlayersByFullName(" Alpha (Alpha)")
	assert.Len(t, p, 2)

	q := db.FindPlayersByFullName(" Beta (Beta) ")
	assert.Len(t, q, 1)

	r := db.FindPlayersByFullName("Gamma (Gamma)")
	assert.Empty(t, r)
}

func TestGetByNames(t *testing.T) {
	db := NewPlayerDB(testPlayers)
	p := db.FindPlayers("Alpha")
	assert.Len(t, p, 2)

	q := db.FindPlayers("贝塔")
	assert.Len(t, q, 1)

	r := db.FindPlayers("Gamma")
	assert.Empty(t, r)
}
