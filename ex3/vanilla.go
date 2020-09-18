package ex3

import (
	"database/sql"
	"log"
)

func VanillaExample(db *sql.DB) {
	var submissions []Submission

	rows, err := db.Query(`
SELECT
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
	JOIN teams AS t ON t.team_id = s.team_id
	JOIN users AS u1 ON u1.team_id = t.team_id
	JOIN users AS u2 ON u2.team_id = t.team_id AND u2.user_id > u1.user_id
LIMIT 10`,
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var submission Submission
		err := rows.Scan(
			// Submission
			&submission.SubmissionID,
			&submission.Answers,
			&submission.Submitted,

			// Assignment
			&submission.Assignment.AssignmentID,
			&submission.Assignment.Questions,

			// Team
			&submission.Team.TeamID,
			&submission.Team.TeamName,

			// User1
			&submission.Team.User1.UserID,
			&submission.Team.User1.Name,
			&submission.Team.User1.Email,
			&submission.Team.User1.CreatedAt,

			// User2
			&submission.Team.User2.UserID,
			&submission.Team.User2.Name,
			&submission.Team.User2.Email,
			&submission.Team.User2.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		submissions = append(submissions, submission)
	}
	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}

	// Print the results
	prettyprint(submissions)
}
