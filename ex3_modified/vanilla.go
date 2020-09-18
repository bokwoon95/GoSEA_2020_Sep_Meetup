package ex3_modified

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

	-- teams
	t.team_id,
	t.team_name,

	-- user
	u.user_id,
	u.name,
	u.email,
	u.created_at
FROM
	submissions AS s
	JOIN assignments AS a ON a.assignment_id = s.assignment_id
	JOIN teams AS t ON t.team_id = s.team_id
	JOIN users AS u ON u.team_id = t.team_id`,
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

			// User
			&submission.Team.User.UserID,
			&submission.Team.User.Name,
			&submission.Team.User.Email,
			&submission.Team.User.CreatedAt,
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
	// prettyprint(submissions)
}
