package models

// type Item struct {
//     Comments *Comments `bson:"-"`
//
//     TakenAt          float64 `bson:"taken_at"`
//     Pk               int64   `bson:"pk"`
//     ID               string  `bson:"id"`
//     CommentsDisabled bool    `bson:"comments_disabled"`
//     DeviceTimestamp  int64   `bson:"device_timestamp"`
//     MediaType        int     `bson:"media_type"`
//     Code             string  `bson:"code"`
//     ClientCacheKey   string  `bson:"client_cache_key"`
//     FilterType       int     `bson:"filter_type"`
//     CarouselParentID string  `bson:"carousel_parent_id"`
//     CarouselMedia    []Item  `bson:"carousel_media,omitempty"`
//     User             User    `bson:"user"`
//     CanViewerReshare bool    `bson:"can_viewer_reshare"`
//     Caption          Caption `bson:"caption"`
//     CaptionIsEdited  bool    `bson:"caption_is_edited"`
//     Likes            int     `bson:"like_count"`
//     HasLiked         bool    `bson:"has_liked"`
//     // Toplikers can be `string` or `[]string`.
//     // Use TopLikers function instead of getting it directly.
//     Toplikers                    interface{} `bson:"top_likers"`
//     Likers                       []User      `bson:"likers"`
//     CommentLikesEnabled          bool        `bson:"comment_likes_enabled"`
//     CommentThreadingEnabled      bool        `bson:"comment_threading_enabled"`
//     HasMoreComments              bool        `bson:"has_more_comments"`
//     MaxNumVisiblePreviewComments int         `bson:"max_num_visible_preview_comments"`
//     // Previewcomments can be `string` or `[]string` or `[]Comment`.
//     // Use PreviewComments function instead of getting it directly.
//     Previewcomments interface{} `bson:"preview_comments,omitempty"`
//     CommentCount    int         `bson:"comment_count"`
//     PhotoOfYou      bool        `bson:"photo_of_you"`
//     // Tags are tagged people in photo
//     Tags struct {
//         In []Tag `bson:"in"`
//     }   `bson:"usertags,omitempty"`
//     FbUserTags           Tag    `bson:"fb_user_tags"`
//     CanViewerSave        bool   `bson:"can_viewer_save"`
//     OrganicTrackingToken string `bson:"organic_tracking_token"`
//     // Images contains URL images in different versions.
//     // Version = quality.
//     Images          Images   `bson:"image_versions2,omitempty"`
//     OriginalWidth   int      `bson:"original_width,omitempty"`
//     OriginalHeight  int      `bson:"original_height,omitempty"`
//     ImportedTakenAt int64    `bson:"imported_taken_at,omitempty"`
//     Location        Location `bson:"location,omitempty"`
//     Lat             float64  `bson:"lat,omitempty"`
//     Lng             float64  `bson:"lng,omitempty"`
//
//     // Videos
//     Videos            []Video `bson:"video_versions,omitempty"`
//     HasAudio          bool    `bson:"has_audio,omitempty"`
//     VideoDuration     float64 `bson:"video_duration,omitempty"`
//     ViewCount         float64 `bson:"view_count,omitempty"`
//     IsDashEligible    int     `bson:"is_dash_eligible,omitempty"`
//     VideoDashManifest string  `bson:"video_dash_manifest,omitempty"`
//     NumberOfQualities int     `bson:"number_of_qualities,omitempty"`
//
//     // Only for stories
//     StoryEvents              []interface{} `bson:"story_events"`
//     StoryHashtags            []interface{} `bson:"story_hashtags"`
//     StoryPolls               []interface{} `bson:"story_polls"`
//     StoryFeedMedia           []interface{} `bson:"story_feed_media"`
//     StorySoundOn             []interface{} `bson:"story_sound_on"`
//     CreativeConfig           interface{}   `bson:"creative_config"`
//     StoryLocations           []interface{} `bson:"story_locations"`
//     StorySliders             []interface{} `bson:"story_sliders"`
//     StoryQuestions           []interface{} `bson:"story_questions"`
//     StoryProductItems        []interface{} `bson:"story_product_items"`
//     SupportsReelReactions    bool          `bson:"supports_reel_reactions"`
//     ShowOneTapFbShareTooltip bool          `bson:"show_one_tap_fb_share_tooltip"`
//     HasSharedToFb            int64         `bson:"has_shared_to_fb"`
//     Mentions                 []Mentions
// }
