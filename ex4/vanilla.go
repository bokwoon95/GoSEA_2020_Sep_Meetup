package ex4

import (
	"database/sql"
	"log"
)

func VanillaExample(db *sql.DB) {
	var submissions []Submission

	rows, err := db.Query(`
SELECT
	-- submissions
	submission_id,
	answers,
	submitted,

	-- assignments
	assignment_id,
	questions,

	-- teams
	team_id,
	team_name,

	-- user1
	u1_user_id,
	u1_name,
	u1_email,
	u1_created_at,

	-- user2
	u2_user_id,
	u2_name,
	u2_email,
	u2_created_at
FROM
	v_submissions`,
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
			Valid:    teamID.Valid,
			TeamID:   int(teamID.Int64),
			TeamName: teamName.String,
		}
		submission.Team.User1 = User{
			Valid:  u1UserID.Valid,
			UserID: int(u1UserID.Int64),
			Name:   u1Name.String,
			Email:  u1Email.String,
		}
		submission.Team.User2 = User{
			Valid:  u2UserID.Valid,
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
