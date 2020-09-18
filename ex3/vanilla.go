package ex3

import (
	"database/sql"
	"log"
)

func VanillaExample(db *sql.DB) {
	var submissions []Submission

	rows, err := db.Query(`
SELECT DISTINCT ON (s.team_id)
	-- submissions
	s.submission_id,
	s.answers,
	s.submitted,

	-- assignments
	a.assignment_id,
	a.questions,

	-- team
	t.team_id,
	t.team_name,

	-- user1
	u1.user_id,
	u1.name,
	u1.email,
	u1.created_at,

	-- user2
	u2.user_id,
	u2.name,
	u2.email,
	u2.created_at
FROM
	submissions AS s
	JOIN assignments AS a ON a.assignment_id = s.assignment_id
	LEFT JOIN teams AS t ON t.team_id = s.team_id
	LEFT JOIN users AS u1 ON u1.team_id = t.team_id
	LEFT JOIN users AS u2 ON u2.team_id = t.team_id AND u2.user_id > u1.user_id
ORDER BY
	s.team_id
	,u1.user_id NULLS LAST
	,u2.user_id NULLS LAST`,
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var submission Submission
		var u1UserID, u2UserID, teamID sql.NullInt64
		var u1Name, u1Email, u2Name, u2Email, teamName sql.NullString
		var u1CreatedAt, u2CreatedAt sql.NullTime
		err := rows.Scan(
			// Submission
			&submission.SubmissionID,
			&submission.Answers,
			&submission.Submitted,

			// Assignment
			&submission.Assignment.AssignmentID,
			&submission.Assignment.Questions,

			// Team
			&teamID,
			&teamName,

			// User1
			&u1UserID,
			&u1Name,
			&u1Email,
			&u1CreatedAt,

			// User2
			&u2UserID,
			&u2Name,
			&u2Email,
			&u2CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		submission.Team = Team{
			TeamID:   int(teamID.Int64),
			TeamName: teamName.String,
		}
		submission.Team.User1 = User{
			UserID: int(u1UserID.Int64),
			Name:   u1Name.String,
			Email:  u1Email.String,
		}
		submission.Team.User2 = User{
			UserID: int(u2UserID.Int64),
			Name:   u2Name.String,
			Email:  u2Email.String,
		}
		submissions = append(submissions, submission)
	}
	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}

	// Print the results
	prettyprint(submissions)
}
