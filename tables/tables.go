// Code generated by 'sqgen-postgres tables'; DO NOT EDIT.
package tables

import (
	sq "github.com/bokwoon95/go-structured-query/postgres"
)

// TABLE_ASSIGNMENTS references the public.assignments table.
type TABLE_ASSIGNMENTS struct {
	*sq.TableInfo
	ASSIGNMENT_ID sq.NumberField
	QUESTIONS     sq.JSONField
}

// ASSIGNMENTS creates an instance of the public.assignments table.
func ASSIGNMENTS() TABLE_ASSIGNMENTS {
	tbl := TABLE_ASSIGNMENTS{TableInfo: &sq.TableInfo{
		Schema: "public",
		Name:   "assignments",
	}}
	tbl.ASSIGNMENT_ID = sq.NewNumberField("assignment_id", tbl.TableInfo)
	tbl.QUESTIONS = sq.NewJSONField("questions", tbl.TableInfo)
	return tbl
}

// As modifies the alias of the underlying table.
func (tbl TABLE_ASSIGNMENTS) As(alias string) TABLE_ASSIGNMENTS {
	tbl.TableInfo.Alias = alias
	return tbl
}

// TABLE_SUBMISSIONS references the public.submissions table.
type TABLE_SUBMISSIONS struct {
	*sq.TableInfo
	ANSWERS       sq.JSONField
	ASSIGNMENT_ID sq.NumberField
	SUBMISSION_ID sq.NumberField
	SUBMITTED     sq.BooleanField
	TEAM_ID       sq.NumberField
}

// SUBMISSIONS creates an instance of the public.submissions table.
func SUBMISSIONS() TABLE_SUBMISSIONS {
	tbl := TABLE_SUBMISSIONS{TableInfo: &sq.TableInfo{
		Schema: "public",
		Name:   "submissions",
	}}
	tbl.ANSWERS = sq.NewJSONField("answers", tbl.TableInfo)
	tbl.ASSIGNMENT_ID = sq.NewNumberField("assignment_id", tbl.TableInfo)
	tbl.SUBMISSION_ID = sq.NewNumberField("submission_id", tbl.TableInfo)
	tbl.SUBMITTED = sq.NewBooleanField("submitted", tbl.TableInfo)
	tbl.TEAM_ID = sq.NewNumberField("team_id", tbl.TableInfo)
	return tbl
}

// As modifies the alias of the underlying table.
func (tbl TABLE_SUBMISSIONS) As(alias string) TABLE_SUBMISSIONS {
	tbl.TableInfo.Alias = alias
	return tbl
}

// TABLE_TEAMS references the public.teams table.
type TABLE_TEAMS struct {
	*sq.TableInfo
	TEAM_ID   sq.NumberField
	TEAM_NAME sq.StringField
}

// TEAMS creates an instance of the public.teams table.
func TEAMS() TABLE_TEAMS {
	tbl := TABLE_TEAMS{TableInfo: &sq.TableInfo{
		Schema: "public",
		Name:   "teams",
	}}
	tbl.TEAM_ID = sq.NewNumberField("team_id", tbl.TableInfo)
	tbl.TEAM_NAME = sq.NewStringField("team_name", tbl.TableInfo)
	return tbl
}

// As modifies the alias of the underlying table.
func (tbl TABLE_TEAMS) As(alias string) TABLE_TEAMS {
	tbl.TableInfo.Alias = alias
	return tbl
}

// TABLE_USERS references the public.users table.
type TABLE_USERS struct {
	*sq.TableInfo
	CREATED_AT sq.TimeField
	DELETED_AT sq.TimeField
	EMAIL      sq.StringField
	NAME       sq.StringField
	TEAM_ID    sq.NumberField
	USER_ID    sq.NumberField
}

// USERS creates an instance of the public.users table.
func USERS() TABLE_USERS {
	tbl := TABLE_USERS{TableInfo: &sq.TableInfo{
		Schema: "public",
		Name:   "users",
	}}
	tbl.CREATED_AT = sq.NewTimeField("created_at", tbl.TableInfo)
	tbl.DELETED_AT = sq.NewTimeField("deleted_at", tbl.TableInfo)
	tbl.EMAIL = sq.NewStringField("email", tbl.TableInfo)
	tbl.NAME = sq.NewStringField("name", tbl.TableInfo)
	tbl.TEAM_ID = sq.NewNumberField("team_id", tbl.TableInfo)
	tbl.USER_ID = sq.NewNumberField("user_id", tbl.TableInfo)
	return tbl
}

// As modifies the alias of the underlying table.
func (tbl TABLE_USERS) As(alias string) TABLE_USERS {
	tbl.TableInfo.Alias = alias
	return tbl
}

// VIEW_V_SUBMISSIONS references the public.v_submissions view.
type VIEW_V_SUBMISSIONS struct {
	*sq.TableInfo
	ANSWERS       sq.JSONField
	ASSIGNMENT_ID sq.NumberField
	QUESTIONS     sq.JSONField
	SUBMISSION_ID sq.NumberField
	SUBMITTED     sq.BooleanField
	TEAM_ID       sq.NumberField
	TEAM_NAME     sq.StringField
	U1_CREATED_AT sq.TimeField
	U1_EMAIL      sq.StringField
	U1_NAME       sq.StringField
	U1_USER_ID    sq.NumberField
	U2_CREATED_AT sq.TimeField
	U2_EMAIL      sq.StringField
	U2_NAME       sq.StringField
	U2_USER_ID    sq.NumberField
}

// V_SUBMISSIONS creates an instance of the public.v_submissions view.
func V_SUBMISSIONS() VIEW_V_SUBMISSIONS {
	tbl := VIEW_V_SUBMISSIONS{TableInfo: &sq.TableInfo{
		Schema: "public",
		Name:   "v_submissions",
	}}
	tbl.ANSWERS = sq.NewJSONField("answers", tbl.TableInfo)
	tbl.ASSIGNMENT_ID = sq.NewNumberField("assignment_id", tbl.TableInfo)
	tbl.QUESTIONS = sq.NewJSONField("questions", tbl.TableInfo)
	tbl.SUBMISSION_ID = sq.NewNumberField("submission_id", tbl.TableInfo)
	tbl.SUBMITTED = sq.NewBooleanField("submitted", tbl.TableInfo)
	tbl.TEAM_ID = sq.NewNumberField("team_id", tbl.TableInfo)
	tbl.TEAM_NAME = sq.NewStringField("team_name", tbl.TableInfo)
	tbl.U1_CREATED_AT = sq.NewTimeField("u1_created_at", tbl.TableInfo)
	tbl.U1_EMAIL = sq.NewStringField("u1_email", tbl.TableInfo)
	tbl.U1_NAME = sq.NewStringField("u1_name", tbl.TableInfo)
	tbl.U1_USER_ID = sq.NewNumberField("u1_user_id", tbl.TableInfo)
	tbl.U2_CREATED_AT = sq.NewTimeField("u2_created_at", tbl.TableInfo)
	tbl.U2_EMAIL = sq.NewStringField("u2_email", tbl.TableInfo)
	tbl.U2_NAME = sq.NewStringField("u2_name", tbl.TableInfo)
	tbl.U2_USER_ID = sq.NewNumberField("u2_user_id", tbl.TableInfo)
	return tbl
}

// As modifies the alias of the underlying view.
func (tbl VIEW_V_SUBMISSIONS) As(alias string) VIEW_V_SUBMISSIONS {
	tbl.TableInfo.Alias = alias
	return tbl
}