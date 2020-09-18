package ex3

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/tables"
	sq "github.com/bokwoon95/go-structured-query/postgres"
)


type User struct {
	UserID    int
	Name      string
	Email     string
	CreatedAt time.Time
}

func (user *User) RowMapper(u tables.TABLE_USERS) func(*sq.Row) {
	return func(row *sq.Row) {
		user.UserID = row.Int(u.USER_ID)
		user.Name = row.String(u.NAME)
		user.Email = row.String(u.EMAIL)
		user.CreatedAt = row.Time(u.CREATED_AT)
	}
}

type Team struct {
	TeamID   int
	TeamName string
	User1    User
	User2    User
}

func (team *Team) RowMapper(t tables.TABLE_TEAMS, u1, u2 tables.TABLE_USERS) func(*sq.Row) {
	return func(row *sq.Row) {
		team.TeamID = row.Int(t.TEAM_ID)
		team.TeamName = row.String(t.TEAM_NAME)
		team.User1.RowMapper(u1)(row)
		team.User2.RowMapper(u2)(row)
	}
}

type Assignment struct {
	AssignmentID int
	Questions    json.RawMessage
}

func (assign *Assignment) RowMapper(a tables.TABLE_ASSIGNMENTS) func(*sq.Row) {
	return func(row *sq.Row) {
		assign.AssignmentID = row.Int(a.ASSIGNMENT_ID)
		row.ScanInto(&assign.Questions, a.QUESTIONS)
	}
}

type Submission struct {
	SubmissionID int
	Assignment   Assignment
	Answers      json.RawMessage
	Team         Team
	Submitted    bool
}

func (subm *Submission) RowMapper(
	s tables.TABLE_SUBMISSIONS,
	a tables.TABLE_ASSIGNMENTS,
	t tables.TABLE_TEAMS,
	u1, u2 tables.TABLE_USERS,
) func(*sq.Row) {
	return func(row *sq.Row) {
		subm.SubmissionID = row.Int(s.SUBMISSION_ID)
		subm.Assignment.RowMapper(a)(row)
		row.ScanInto(&subm.Answers, s.ANSWERS)
		subm.Team.RowMapper(t, u1, u2)(row)
		subm.Submitted = row.Bool(s.SUBMITTED)
	}
}

func prettyprint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}
