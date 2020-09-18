package ex5

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/tables"
	sq "github.com/bokwoon95/go-structured-query/postgres"
)

type User struct {
	Valid     bool
	UserID    int
	Name      string
	Email     string
	CreatedAt time.Time
}

type Team struct {
	Valid    bool
	TeamID   int
	TeamName string
	User1    User
	User2    User
}

type Assignment struct {
	AssignmentID int
	Questions    json.RawMessage
}

type Submission struct {
	SubmissionID int
	Assignment   Assignment
	Answers      json.RawMessage
	Team         Team
	Submitted    bool
}

func (subm *Submission) RowMapper(vs tables.VIEW_V_SUBMISSIONS) func(*sq.Row) {
	return func(row *sq.Row) {
		*subm = Submission{
			SubmissionID: row.Int(vs.SUBMISSION_ID),
			Submitted:    row.Bool(vs.SUBMITTED),
			Assignment: Assignment{
				AssignmentID: row.Int(vs.ASSIGNMENT_ID),
			},
			Team: Team{
				Valid:    row.IntValid(vs.TEAM_ID),
				TeamID:   row.Int(vs.TEAM_ID),
				TeamName: row.String(vs.TEAM_NAME),
				User1: User{
					Valid:     row.IntValid(vs.U1_USER_ID),
					UserID:    row.Int(vs.U1_USER_ID),
					Name:      row.String(vs.U1_NAME),
					Email:     row.String(vs.U1_EMAIL),
					CreatedAt: row.Time(vs.U1_CREATED_AT),
				},
				User2: User{
					Valid:     row.IntValid(vs.U2_USER_ID),
					UserID:    row.Int(vs.U2_USER_ID),
					Name:      row.String(vs.U2_NAME),
					Email:     row.String(vs.U2_EMAIL),
					CreatedAt: row.Time(vs.U2_CREATED_AT),
				},
			},
		}
		row.ScanInto(subm.Assignment.Questions, vs.QUESTIONS)
		row.ScanInto(subm.Answers, vs.ANSWERS)
	}
}

type SqlxUser struct {
	UserID    int       `db:"user_id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	TeamID    int       `db:"team_id"`
}

func prettyprint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}
