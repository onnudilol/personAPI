package main

import (
	"database/sql"
	"log"
	"sort"
)

type Person struct {
	Id          int    `json:"id"`
	DisplayName string `json:"displayname"`
	IsTeam      bool   `json:"is_team"`
}

func getTeams(db *sql.DB, person string) ([]Person, error) {
	var PersonID int
	err := db.QueryRow("SELECT id FROM person_person WHERE displayname=$1", person).Scan(&PersonID)

	if err != nil {
		log.Println(err)
	}

	rows, err := db.Query("SELECT id, displayname, is_team FROM person_person WHERE is_team=TRUE")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var teams []Person

	for rows.Next() {
		var p Person

		if err := rows.Scan(&p.Id, &p.DisplayName, &p.IsTeam); err != nil {
			return nil, err
		}

		members, _ := getMembers(db, p.DisplayName)

		for _, element := range members {

			if element.Id == PersonID {
				teams = append(teams, p)
			}
		}
	}

	return teams, nil
}

func getMembers(db *sql.DB, person string) ([]Person, error) {
	var PersonID int
	err := db.QueryRow("SELECT p.id FROM person_person p WHERE p.displayname=$1", person).Scan(&PersonID)

	if err != nil {
		log.Println(err)
	}

	var allMembers = make(map[Person]bool)
	hasTeams := true
	var teamCache = make(map[string]bool)

	people, err := queryMembers(db, PersonID)

	for hasTeams {

		for _, element := range people {

			if element.IsTeam == false {
				allMembers[element] = true

			} else {

				if _, value := teamCache[element.DisplayName]; !value {
					teamCache[element.DisplayName] = true
					subMembers, _ := queryMembers(db, element.Id)
					people = append(people, subMembers...)
				}
			}
		}

		hasTeams = checkHasNewTeams(people, &teamCache)
	}

	var memberSlice []Person

	for key := range allMembers {
		memberSlice = append(memberSlice, key)
	}

	sort.Slice(memberSlice, func(i int, j int) bool { return memberSlice[i].DisplayName < memberSlice[j].DisplayName })

	return memberSlice, nil
}

func checkHasNewTeams(members []Person, teamCache *map[string]bool) bool {
	for _, element := range members {

		if element.IsTeam {

			if _, value := (*teamCache)[element.DisplayName]; !value {
				return true
			}
		}
	}

	return false
}

func queryMembers(db *sql.DB, personID int) ([]Person, error) {
	rows, err := db.Query("SELECT m1.id, m1.displayname, m1.is_team FROM person_person m1 INNER JOIN person_person_members m2 on m1.id = m2.to_person_id WHERE m2.from_person_id=$1", personID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var persons []Person

	for rows.Next() {
		var p Person

		if err := rows.Scan(&p.Id, &p.DisplayName, &p.IsTeam); err != nil {
			return nil, err
		}

		persons = append(persons, p)
	}

	return persons, err
}
