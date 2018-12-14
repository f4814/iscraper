package models

import "github.com/ahmdrz/goinsta"

type Comment struct {
	ID                             int64   `bson:"pk"`
	Text                           string  `bson:"text"`
	Type                           int     `bson:"type"`
	UserID                         int64   `bson:"user_id"`
	BitFlags                       int     `bson:"bit_flags"`
	ChildCommentCount              int     `bson:"child_comment_count"`
	CommentIndex                   int     `bson:"comment_index"`
	CommentLikeCount               int     `bson:"comment_like_count"`
	ContentType                    string  `bson:"content_type"`
	CreatedAt                      int64   `bson:"created_at"`
	CreatedAtUtc                   int64   `bson:"created_at_utc"`
	DidReportAsSpam                bool    `bson:"did_report_as_spam"`
	HasLikedComment                bool    `bson:"has_liked_comment"`
	InlineComposerDisplayCondition string  `bson:"inline_composer_display_condition"`
	OtherPreviewUsersID            []int64 `bson:"other_preview_users"`
	PreviewChildCommentsID         []int64 `bson:"preview_child_comments"`
	NextMaxChildCursor             string  `bson:"next_max_child_cursor,omitempty"`
	HasMoreTailChildComments       bool    `bson:"has_more_tail_child_comments,omitempty"`
	NextMinChildCursor             string  `bson:"next_min_child_cursor,omitempty"`
	HasMoreHeadChildComments       bool    `bson:"has_more_head_child_comments,omitempty"`
	NumTailChildComments           int     `bson:"num_tail_child_comments,omitempty"`
	NumHeadChildComments           int     `bson:"num_head_child_comments,omitempty"`
	Status                         string  `bson:"status"`
	// contains filtered or unexported fields
}

func (m *Comment) FromIG(c *goinsta.Comment) {
	m.ID = c.ID
	m.Text = c.Text
	m.Type = c.Type
	m.UserID = c.UserID
	m.BitFlags = c.BitFlags
	m.ChildCommentCount = c.ChildCommentCount
	m.CommentIndex = c.CommentIndex
	m.CommentLikeCount = c.CommentLikeCount
	m.ContentType = c.ContentType
	m.CreatedAt = c.CreatedAt
	m.CreatedAtUtc = c.CreatedAtUtc
	m.DidReportAsSpam = c.DidReportAsSpam
	m.HasLikedComment = c.HasLikedComment
	m.InlineComposerDisplayCondition = c.InlineComposerDisplayCondition
	m.NextMaxChildCursor = c.NextMaxChildCursor
	m.HasMoreTailChildComments = c.HasMoreTailChildComments
	m.NextMinChildCursor = c.NextMinChildCursor
	m.HasMoreHeadChildComments = c.HasMoreHeadChildComments
	m.NumTailChildComments = c.NumTailChildComments
	m.NumHeadChildComments = c.NumHeadChildComments
	m.Status = c.Status

	for _, v := range c.OtherPreviewUsers {
		m.OtherPreviewUsersID = append(m.OtherPreviewUsersID, v.ID)
	}

	for _, v := range c.PreviewChildComments {
		m.PreviewChildCommentsID = append(m.PreviewChildCommentsID, v.ID)
	}
}
