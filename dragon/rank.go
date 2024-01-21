package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sort"
)

type Rank struct {
	DragonID        int
	Name            string
	Experience      int
	ExperienceStage int
	Attack          int
	Defense         int
	Life            int
}

// ByExperience implements sort.Interface for []*Rank based on Experience field
type ByExperience []*Rank

func (a ByExperience) Len() int           { return len(a) }
func (a ByExperience) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByExperience) Less(i, j int) bool { return a[i].Experience > a[j].Experience }

func createTableIfNotExists(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS ranks (
	    dragon_id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT,
	    experience INTEGER,
	    experience_stage INTEGER,
	    attack INTEGER,
	    defense INTEGER,
	    life INTEGER
	);
	`
	_, err := db.Exec(createTableSQL)
	return err
}

func (r *Rank) save() {
	// 打开数据库连接。如果数据库文件不存在，将会创建它。
	db, err := sql.Open("sqlite3", "rank.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查并创建表
	if err := createTableIfNotExists(db); err != nil {
		log.Fatal(err)
	}

	// 插入数据
	insertSQL := `
	INSERT INTO ranks (name, experience, experience_stage, attack, defense, life) VALUES (?, ?, ?, ?, ?, ?);
	`
	_, err = db.Exec(insertSQL, r.Name, r.Experience, r.ExperienceStage, r.Attack, r.Defense, r.Life)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Rank) equal(rank *Rank) bool {
	if rank == nil {
		return false
	}
	return r.Name == rank.Name && r.Experience == rank.Experience && r.ExperienceStage == rank.ExperienceStage && r.Attack == rank.Attack && r.Defense == rank.Defense && r.Life == rank.Life

}

func getRanks() []*Rank {
	// 打开数据库连接。如果数据库文件不存在，将会创建它。
	db, err := sql.Open("sqlite3", "rank.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查并创建表
	if err := createTableIfNotExists(db); err != nil {
		log.Fatal(err)
	}

	// 查询数据
	querySQL := `
	SELECT dragon_id, name, experience, experience_stage, attack, defense, life FROM ranks ORDER BY experience DESC LIMIT 10;
	`
	rows, err := db.Query(querySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ranks := make([]*Rank, 0)
	for rows.Next() {
		rank := &Rank{}
		err := rows.Scan(&rank.DragonID, &rank.Name, &rank.Experience, &rank.ExperienceStage, &rank.Attack, &rank.Defense, &rank.Life)
		if err != nil {
			log.Fatal(err)
		}
		ranks = append(ranks, rank)
	}

	// 对 ranks 切片按照 Experience 字段降序排序
	sort.Sort(ByExperience(ranks))

	// 取前十个元素
	if len(ranks) > ShowRankSize {
		return ranks[:ShowRankSize]
	}

	return ranks
}

func newRank(d *Dragon) *Rank {
	return &Rank{
		Name:            d.Name,
		Experience:      d.Experience,
		ExperienceStage: d.ExperienceStage,
		Attack:          d.basic.attack,
		Defense:         d.basic.defense,
		Life:            d.basic.maxLife,
	}
}

func showRank(ranks []*Rank, rank *Rank) {
	p.rankText.Reset()
	for i, r := range ranks {
		s := fmt.Sprintf("第%v名，龙的ID：%v，名称：%v，经验值：%v，攻击力：%v，防御力：%v，生命值：%v", i+1, r.DragonID, r.Name, r.Experience, r.Attack, r.Defense, r.Life)
		if r.equal(rank) {
			s = "👑" + s
		}
		if i > 0 {
			s = s + "\n"
		}
		p.addRankLn(newRankInfo(s))
	}
}
