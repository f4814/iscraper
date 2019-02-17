package models

import (
	"time"

	"github.com/ahmdrz/goinsta/v2"
	driver "github.com/arangodb/go-driver"
)

// TODO Tags, Comments, FBTags, Mentions, Carousel Media, Toplikers
type Item struct {
	ScrapedAt time.Time `json:"scraped_at"`
	AddedAt   time.Time `json:"added_at"`
	meta      driver.DocumentMeta

	TakenAt                      float64          `json:"taken_at"`
	Pk                           int64            `json:"pk"`
	ID                           string           `json:"id"`
	CommentsDisabled             bool             `json:"comments_disabled"`
	DeviceTimestamp              int64            `json:"device_timestamp"`
	MediaType                    int              `json:"media_type"`
	Code                         string           `json:"code"`
	ClientCacheKey               string           `json:"client_cache_key"`
	FilterType                   int              `json:"filter_type"`
	CarouselParentID             string           `json:"carousel_parent_id"`
	CanViewerReshare             bool             `json:"can_viewer_reshare"`
	Caption                      goinsta.Caption  `json:"caption"`
	CaptionIsEdited              bool             `json:"caption_is_edited"`
	Likes                        int              `json:"like_count"`
	HasLiked                     bool             `json:"has_liked"`
	Toplikers                    interface{}      `json:"top_likers"`
	CommentLikesEnabled          bool             `json:"comment_likes_enabled"`
	CommentThreadingEnabled      bool             `json:"comment_threading_enabled"`
	HasMoreComments              bool             `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int              `json:"max_num_visible_preview_comments"`
	Previewcomments              interface{}      `json:"preview_comments,omitempty"`
	CommentCount                 int              `json:"comment_count"`
	PhotoOfYou                   bool             `json:"photo_of_you"`
	CanViewerSave                bool             `json:"can_viewer_save"`
	OrganicTrackingToken         string           `json:"organic_tracking_token"`
	Images                       goinsta.Images   `json:"image_versions2,omitempty"`
	OriginalWidth                int              `json:"original_width,omitempty"`
	OriginalHeight               int              `json:"original_height,omitempty"`
	ImportedTakenAt              int64            `json:"imported_taken_at,omitempty"`
	Location                     goinsta.Location `json:"location,omitempty"`
	Lat                          float64          `json:"lat,omitempty"`
	Lng                          float64          `json:"lng,omitempty"`

	// Videos
	Videos            []goinsta.Video `json:"video_versions,omitempty"`
	HasAudio          bool            `json:"has_audio,omitempty"`
	VideoDuration     float64         `json:"video_duration,omitempty"`
	ViewCount         float64         `json:"view_count,omitempty"`
	IsDashEligible    int             `json:"is_dash_eligible,omitempty"`
	VideoDashManifest string          `json:"video_dash_manifest,omitempty"`
	NumberOfQualities int             `json:"number_of_qualities,omitempty"`

	// Only for stories
	StoryEvents              []interface{} `json:"story_events"`
	StoryHashtags            []interface{} `json:"story_hashtags"`
	StoryPolls               []interface{} `json:"story_polls"`
	StoryFeedMedia           []interface{} `json:"story_feed_media"`
	StorySoundOn             []interface{} `json:"story_sound_on"`
	CreativeConfig           interface{}   `json:"creative_config"`
	StoryLocations           []interface{} `json:"story_locations"`
	StorySliders             []interface{} `json:"story_sliders"`
	StoryQuestions           []interface{} `json:"story_questions"`
	StoryProductItems        []interface{} `json:"story_product_items"`
	SupportsReelReactions    bool          `json:"supports_reel_reactions"`
	ShowOneTapFbShareTooltip bool          `json:"show_one_tap_fb_share_tooltip"`
	HasSharedToFb            int64         `json:"has_shared_to_fb"`
}

func NewItem(i goinsta.Item) *Item {
	return &Item{
		AddedAt:                      time.Now(),
		ScrapedAt:                    time.Unix(0, 0),
		TakenAt:                      i.TakenAt,
		Pk:                           i.Pk,
		ID:                           i.ID,
		CommentsDisabled:             i.CommentsDisabled,
		DeviceTimestamp:              i.DeviceTimestamp,
		MediaType:                    i.MediaType,
		Code:                         i.Code,
		ClientCacheKey:               i.ClientCacheKey,
		FilterType:                   i.FilterType,
		CarouselParentID:             i.CarouselParentID,
		CanViewerReshare:             i.CanViewerReshare,
		Caption:                      i.Caption,
		CaptionIsEdited:              i.CaptionIsEdited,
		Likes:                        i.Likes,
		HasLiked:                     i.HasLiked,
		CommentLikesEnabled:          i.CommentLikesEnabled,
		CommentThreadingEnabled:      i.CommentThreadingEnabled,
		HasMoreComments:              i.HasMoreComments,
		MaxNumVisiblePreviewComments: i.MaxNumVisiblePreviewComments,
		Previewcomments:              i.Previewcomments,
		CommentCount:                 i.CommentCount,
		PhotoOfYou:                   i.PhotoOfYou,
		CanViewerSave:                i.CanViewerSave,
		OrganicTrackingToken:         i.OrganicTrackingToken,
		Images:                       i.Images,
		OriginalWidth:                i.OriginalWidth,
		OriginalHeight:               i.OriginalHeight,
		ImportedTakenAt:              i.ImportedTakenAt,
		Location:                     i.Location,
		Lat:                          i.Lat,
		Lng:                          i.Lng,
		Videos:                       i.Videos,
		HasAudio:                     i.HasAudio,
		VideoDuration:                i.VideoDuration,
		ViewCount:                    i.ViewCount,
		IsDashEligible:               i.IsDashEligible,
		VideoDashManifest:            i.VideoDashManifest,
		NumberOfQualities:            i.NumberOfQualities,
		StoryEvents:                  i.StoryEvents,
		StoryHashtags:                i.StoryHashtags,
		StoryPolls:                   i.StoryPolls,
		StoryFeedMedia:               i.StoryFeedMedia,
		StorySoundOn:                 i.StorySoundOn,
		CreativeConfig:               i.CreativeConfig,
		StoryLocations:               i.StoryLocations,
		StorySliders:                 i.StorySliders,
		StoryQuestions:               i.StoryQuestions,
		StoryProductItems:            i.StoryProductItems,
		SupportsReelReactions:        i.SupportsReelReactions,
		ShowOneTapFbShareTooltip:     i.ShowOneTapFbShareTooltip,
		HasSharedToFb:                i.HasSharedToFb,
	}
}

func (i *Item) GetMeta() driver.DocumentMeta {
	return i.meta
}

func (i *Item) SetMeta(meta driver.DocumentMeta) {
	i.meta = meta
}
