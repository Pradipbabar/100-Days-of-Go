package player

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHadAGoodGame(t *testing.T) {
	tests := []struct {
		name string
		stats Stats
		goodGame bool
		wantErr string
	}
	{
	{
		"Bad game error testing",
		Stats{
			Name: "Player 1",
			Minutes: 34.1,
			Points: 11,
			Assist: 10,
			TurnOver: -2,
			Rebounds: 26,
		},
		false,
		"Player Stats can't be negative",
	}, { "good game", Stats{
		Name: "Player 1",
		Minutes: 34.1,
		Points: 11,
		Assist: 10,
		TurnOver: 11,
		Rebounds: 26,
	}, true, ""

	}
}

	for _, tt := range tests {
		if tt.wantErr != "" {
			assert.Contains(t, err.Error(), t.wantErr)
		}else{
			assert.Equal(t,goodGame, isGooGame)
		}
	}

	// t.Run("bad game error testing", func(t *testing.T) {
	// 	s := Stats{
	// 		Name: "Player 1",
	// 		Minutes: 34.1,
	// 		Points: 11,
	// 		Assist: 10,
	// 		TurnOver: -2,
	// 		Rebounds: 26,
	// 	}

	// 	_, err := HadAGoodGame(s)
	// 	require.NotNil(t,err)
	// })

	// t.Run("good game testing", func(t *testing.T) {
	// 	s := Stats{
	// 		Name: "Player 1",
	// 		Minutes: 34.1,
	// 		Points: 11,
	// 		Assist: 10,
	// 		TurnOver: 11,
	// 		Rebounds: 26,
	// 	}

	// 	goodGame, err := HadAGoodGame(s)
	// 	require.Nil(t,err)
	// 	assert.True(t,goodGame)
	// })
}