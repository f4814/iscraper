package models

import (
	"github.com/ahmdrz/goinsta"
)

type Item struct {
	ID string `bson:"_id"`

	Comments []Comment `bson:"comments"` // XXX

	TakenAt          float64  `bson:"taken_at"`
	Pk               int64    `bson:"pk"`
	CommentsDisabled bool     `bson:"comments_disabled"`
	DeviceTimestamp  int64    `bson:"device_timestamp"`
	MediaType        int      `bson:"media_type"`
	Code             string   `bson:"code"`
	ClientCacheKey   string   `bson:"client_cache_key"`
	FilterType       int      `bson:"filter_type"`
	CarouselParentID string   `bson:"carousel_parent_id"`
	CarouselMediaIDs []string `bson:"carousel_media_ids,omitempty"`
	UserID           int64    `bson:"user_id"`
	CanViewerReshare bool     `bson:"can_viewer_reshare"`
	Caption          Caption  `bson:"caption"`
	CaptionIsEdited  bool     `bson:"caption_is_edited"`
	Likes            int      `bson:"like_count"`
	HasLiked         bool     `bson:"has_liked"`

	// Toplikers can be `string` or `[]string`.
	// Use TopLikers function instead of getting it directly.
	Toplikers                    interface{} `bson:"top_likers"`
	Likers                       []int64     `bson:"likers"`
	CommentLikesEnabled          bool        `bson:"comment_likes_enabled"`
	CommentThreadingEnabled      bool        `bson:"comment_threading_enabled"`
	HasMoreComments              bool        `bson:"has_more_comments"`
	MaxNumVisiblePreviewComments int         `bson:"max_num_visible_preview_comments"`

	// Previewcomments can be `string` or `[]string` or `[]Comment`.
	// Use PreviewComments function instead of getting it directly.
	Previewcomments interface{} `bson:"preview_comments,omitempty"`
	CommentCount    int         `bson:"comment_count"`
	PhotoOfYou      bool        `bson:"photo_of_you"`

	// Tags are tagged people in photo
	Tags struct {
		In []Tag `bson:"in"`
	} `bson:"usertags,omitempty"`
	FbUserTags           Tag    `bson:"fb_user_tags"`
	CanViewerSave        bool   `bson:"can_viewer_save"`
	OrganicTrackingToken string `bson:"organic_tracking_token"`

	// Images contains URL images in different versions.
	Images          []Image  `bson:"image_versions,omitempty"`
	OriginalWidth   int      `bson:"original_width,omitempty"`
	OriginalHeight  int      `bson:"original_height,omitempty"`
	ImportedTakenAt int64    `bson:"imported_taken_at,omitempty"`
	Location        Location `bson:"location,omitempty"`
	Lat             float64  `bson:"lat,omitempty"`
	Lng             float64  `bson:"lng,omitempty"`

	// Videos
	Videos            []Video `bson:"video_versions,omitempty"`
	HasAudio          bool    `bson:"has_audio,omitempty"`
	VideoDuration     float64 `bson:"video_duration,omitempty"`
	ViewCount         float64 `bson:"view_count,omitempty"`
	IsDashEligible    int     `bson:"is_dash_eligible,omitempty"`
	VideoDashManifest string  `bson:"video_dash_manifest,omitempty"`
	NumberOfQualities int     `bson:"number_of_qualities,omitempty"`

	// Only for stories
	StoryEvents              []interface{} `bson:"story_events,omitempty"`
	StoryHashtags            []interface{} `bson:"story_hashtags,omitempty"`
	StoryPolls               []interface{} `bson:"story_polls,omitempty"`
	StoryFeedMedia           []interface{} `bson:"story_feed_media,omitempty"`
	StorySoundOn             []interface{} `bson:"story_sound_on,omitempty"`
	CreativeConfig           interface{}   `bson:"creative_config,omitempty"`
	StoryLocations           []interface{} `bson:"story_locations,omitempty"`
	StorySliders             []interface{} `bson:"story_sliders,omitempty"`
	StoryQuestions           []interface{} `bson:"story_questions,omitempty"`
	StoryProductItems        []interface{} `bson:"story_product_items,omitempty"`
	SupportsReelReactions    bool          `bson:"supports_reel_reactions,omitempty"`
	ShowOneTapFbShareTooltip bool          `bson:"show_one_tap_fb_share_tooltip,omitempty"`
	HasSharedToFb            int64         `bson:"has_shared_to_fb,omitempty"`
	Mentions                 []Mentions    `bson:"mentions,omitempty"`
}

type Tag struct {
	In []TagIn `bson:"in"`
}

type TagIn struct {
	UserID                int64       `bson:"user_id"`
	Position              []float64   `bson:"position"`
	StartTimeInVideoInSec interface{} `bson:"start_time_in_video_in_sec"`
	DurationInVideoInSec  interface{} `bson:"duration_in_video_in_sec"`
}

type Image struct {
	Width  int    `bson:"width"`
	Height int    `bson:"height"`
	URL    string `bson:"url"`
}

type Location struct {
	Pk               int64   `bson:"pk"`
	Name             string  `bson:"name"`
	Address          string  `bson:"address"`
	City             string  `bson:"city"`
	ShortName        string  `bson:"short_name"`
	Lng              float64 `bson:"lng"`
	Lat              float64 `bson:"lat"`
	ExternalSource   string  `bson:"external_source"`
	FacebookPlacesID int64   `bson:"facebook_places_id"`
}

type Video struct {
	Type   int    `bson:"type"`
	Width  int    `bson:"width"`
	Height int    `bson:"height"`
	URL    string `bson:"url"`
	ID     string `bson:"id"`
}

type Mentions struct {
	X        float64 `bson:"x"`
	Y        float64 `bson:"y"`
	Z        int64   `bson:"z"`
	Width    float64 `bson:"width"`
	Height   float64 `bson:"height"`
	Rotation float64 `bson:"rotation"`
	IsPinned int     `bson:"is_pinned"`
	UserID   int64   `bson:"user_id"`
}

type Caption struct {
	ID              int64  `bson:"pk"`
	UserID          int64  `bson:"user_id"`
	Text            string `bson:"text"`
	Type            int    `bson:"type"`
	CreatedAt       int64  `bson:"created_at"`
	CreatedAtUtc    int64  `bson:"created_at_utc"`
	ContentType     string `bson:"content_type"`
	Status          string `bson:"status"`
	BitFlags        int    `bson:"bit_flags"`
	DidReportAsSpam bool   `bson:"did_report_as_spam"`
	MediaID         int64  `bson:"media_id"`
	HasTranslation  bool   `bson:"has_translation"`
}

func (m *Image) FromIG(i *goinsta.Candidate) {
	m.Width = i.Width
	m.Height = i.Height
	m.URL = i.URL
}

func (m *Tag) FromIG(t *goinsta.Tag) {
	for _, i := range t.In {
		var n TagIn
		n.UserID = i.User.ID
		n.Position = i.Position
		n.StartTimeInVideoInSec = i.StartTimeInVideoInSec
		n.DurationInVideoInSec = i.DurationInVideoInSec
		m.In = append(m.In, n)
	}
}

func (m *Location) FromIG(l *goinsta.Location) {
	m.Pk = l.Pk
	m.Name = l.Name
	m.Address = l.Address
	m.City = l.City
	m.ShortName = l.ShortName
	m.Lng = l.Lng
	m.Lat = l.Lat
	m.ExternalSource = l.ExternalSource
	m.FacebookPlacesID = l.FacebookPlacesID
}

func (m *Mentions) FromIG(i *goinsta.Mentions) {
	m.X = i.X
	m.Y = i.Y
	m.Z = i.Z
	m.Width = i.Width
	m.Height = i.Height
	m.Rotation = i.Rotation
	m.IsPinned = i.IsPinned
	m.UserID = i.User.ID
}

func (m *Caption) FromIG(c *goinsta.Caption) {
	m.ID = c.ID
	m.UserID = c.UserID
	m.Text = c.Text
	m.Type = c.Type
	m.CreatedAt = c.CreatedAt
	m.CreatedAtUtc = c.CreatedAtUtc
	m.ContentType = c.ContentType
	m.Status = c.Status
	m.BitFlags = c.BitFlags
	m.DidReportAsSpam = c.DidReportAsSpam
	m.MediaID = c.MediaID
	m.HasTranslation = c.HasTranslation
}

func (m *Video) FromIG(v *goinsta.Video) {
	m.Type = v.Type
	m.Width = v.Width
	m.URL = v.URL
	m.ID = v.ID
}

func (m *Item) FromIG(i *goinsta.Item) {
	// Comments *Comments `bson:"comments"` // XXX

	m.TakenAt = i.TakenAt
	m.Pk = i.Pk
	m.ID = i.ID
	m.CommentsDisabled = i.CommentsDisabled
	m.DeviceTimestamp = i.DeviceTimestamp
	m.MediaType = i.MediaType
	m.Code = i.Code
	m.ClientCacheKey = i.ClientCacheKey
	m.FilterType = i.FilterType
	m.CarouselParentID = i.CarouselParentID
	m.UserID = i.User.ID
	m.CanViewerReshare = i.CanViewerReshare
	m.Caption.FromIG(&i.Caption)
	m.CaptionIsEdited = i.CaptionIsEdited
	m.Likes = i.Likes
	m.HasLiked = i.HasLiked

	// Toplikers can be `string` or `[]string`.
	// Use TopLikers function instead of getting it directly.
	m.Toplikers = i.Toplikers
	m.CommentLikesEnabled = i.CommentLikesEnabled
	m.CommentThreadingEnabled = i.CommentThreadingEnabled
	m.HasMoreComments = i.HasMoreComments
	m.MaxNumVisiblePreviewComments = i.MaxNumVisiblePreviewComments

	// Previewcomments can be `string` or `[]string` or `[]Comment`.
	// Use PreviewComments function instead of getting it directly.
	m.Previewcomments = i.Previewcomments
	m.CommentCount = i.CommentCount
	m.PhotoOfYou = i.PhotoOfYou

	// Tags are tagged people in photo
	m.CanViewerSave = i.CanViewerSave
	m.OrganicTrackingToken = i.OrganicTrackingToken

	// Images contains URL images in different versions.
	// Version = quality.
	m.OriginalWidth = i.OriginalWidth
	m.OriginalHeight = i.OriginalHeight
	m.ImportedTakenAt = i.ImportedTakenAt
	m.Location.FromIG(&i.Location)
	m.Lat = i.Lat
	m.Lng = i.Lng

	// Videos
	m.HasAudio = i.HasAudio
	m.VideoDuration = i.VideoDuration
	m.ViewCount = i.ViewCount
	m.IsDashEligible = i.IsDashEligible
	m.VideoDashManifest = i.VideoDashManifest
	m.NumberOfQualities = i.NumberOfQualities

	// Only for stories
	m.StoryEvents = i.StoryEvents
	m.StoryHashtags = i.StoryHashtags
	m.StoryPolls = i.StoryPolls
	m.StoryFeedMedia = i.StoryFeedMedia
	m.StorySoundOn = i.StorySoundOn
	m.CreativeConfig = i.CreativeConfig
	m.StoryLocations = i.StoryLocations
	m.StorySliders = i.StorySliders
	m.StoryQuestions = i.StoryQuestions
	m.StoryProductItems = i.StoryProductItems
	m.SupportsReelReactions = i.SupportsReelReactions
	m.ShowOneTapFbShareTooltip = i.ShowOneTapFbShareTooltip
	m.HasSharedToFb = i.HasSharedToFb
	// Mentions                 []Mentions // XXX

	for _, u := range i.Likers {
		m.Likers = append(m.Likers, u.ID)
	}

	for _, i := range i.CarouselMedia {
		m.CarouselMediaIDs = append(m.CarouselMediaIDs, i.ID)
	}

	for _, v := range i.Tags.In {
		var t Tag
		t.FromIG(&v)
		m.Tags.In = append(m.Tags.In, t)
	}

	var t Tag
	t.FromIG(&i.FbUserTags)
	m.FbUserTags = t

	for _, v := range i.Images.Versions {
		var i Image
		i.FromIG(&v)
		m.Images = append(m.Images, i)
	}

	for _, v := range i.Videos {
		var i Video
		i.FromIG(&v)
		m.Videos = append(m.Videos, i)
	}

}
