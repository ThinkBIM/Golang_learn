package mysql

import "fmt"

type GameHero struct {
	Id       int    `db:"id"`
	Cname    string `db:"cname"`
	Title    string `db:"title"`
	NewType  int    `db:"new_type"`
	HeroType int    `db:"hero_type"`
	SkinName string `db:"skin_name"`
	HeroImg  string `db:"hero_img"`
}

func GetAll() []GameHero {
	defer DB.Close()

	var gameHero []GameHero

	err := DB.Select(&gameHero, "SELECT id,cname,title,new_type,skin_name,hero_img FROM game_hero limit 1")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil
	}
	return gameHero
}

func main() {
	user := GetAll()
	fmt.Println(user)
}
