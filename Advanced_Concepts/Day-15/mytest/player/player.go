package player

import "errors"

type Stats struct {
	 Name string
	 Minutes float32
	 Points int 
	 Assist int8
	 TurnOver int8
	 Rebounds int8
} 

func HadAGoodGame(stats Stats)(bool, error)  {
	if stats.Assist < 0 || stats.TurnOver < 0 || stats.Points < 0 || stats.Minutes < 0 || stats.Rebounds < 0 {
		return false, errors.New("Player Stats can't be negative")
	}

	if stats.Assist >= (stats.TurnOver * 2){
		return true, nil
	}

	if stats.Name ==""{
		return false, errors.New("Player Stats can't be empty")
	}

	if stats.Assist >= 10 && stats.Rebounds >= 10 && stats.Points >=10 {
		return true, nil
	} else if stats.Points < 10 && stats.Assist < 10 && stats.Minutes > 25.0 {
		return false, nil
	}
	return false, nil


}