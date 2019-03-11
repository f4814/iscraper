package models

import (
	"time"

	"github.com/ahmdrz/goinsta/v2"
	driver "github.com/arangodb/go-driver"
)

type User struct {
	ScrapedAt time.Time `json:"scraped_at"`
	AddedAt   time.Time `json:"added_at"`
	meta      driver.DocumentMeta

	ID                         int64   `json:"id"`
	Username                   string  `json:"username"`
	FullName                   string  `json:"full_name"`
	Biography                  string  `json:"biography"`
	ProfilePicURL              string  `json:"profile_pic_url"`
	Email                      string  `json:"email"`
	PhoneNumber                string  `json:"phone_number"`
	IsBusiness                 bool    `json:"is_business"`
	Gender                     int     `json:"gender"`
	ProfilePicID               string  `json:"profile_pic_id"`
	HasAnonymousProfilePicture bool    `json:"has_anonymous_profile_picture"`
	IsPrivate                  bool    `json:"is_private"`
	IsUnpublished              bool    `json:"is_unpublished"`
	AllowedCommenterType       string  `json:"allowed_commenter_type"`
	IsVerified                 bool    `json:"is_verified"`
	MediaCount                 int     `json:"media_count"`
	FollowerCount              int     `json:"follower_count"`
	FollowingCount             int     `json:"following_count"`
	FollowingTagCount          int     `json:"following_tag_count"`
	MutualFollowersID          []int64 `json:"profile_context_mutual_follow_ids"`
	ProfileContext             string  `json:"profile_context"`
	GeoMediaCount              int     `json:"geo_media_count"`
	ExternalURL                string  `json:"external_url"`
	HasBiographyTranslation    bool    `json:"has_biography_translation"`
	ExternalLynxURL            string  `json:"external_lynx_url"`
	BiographyWithEntities      struct {
		RawText  string        `json:"raw_text"`
		Entities []interface{} `json:"entities"`
	} `json:"biography_with_entities"`
	UsertagsCount                int                  `json:"usertags_count"`
	HasChaining                  bool                 `json:"has_chaining"`
	IsFavorite                   bool                 `json:"is_favorite"`
	IsFavoriteForStories         bool                 `json:"is_favorite_for_stories"`
	IsFavoriteForHighlights      bool                 `json:"is_favorite_for_highlights"`
	CanBeReportedAsFraud         bool                 `json:"can_be_reported_as_fraud"`
	ShowShoppableFeed            bool                 `json:"show_shoppable_feed"`
	ShoppablePostsCount          int                  `json:"shoppable_posts_count"`
	ReelAutoArchive              string               `json:"reel_auto_archive"`
	HasHighlightReels            bool                 `json:"has_highlight_reels"`
	PublicEmail                  string               `json:"public_email"`
	PublicPhoneNumber            string               `json:"public_phone_number"`
	PublicPhoneCountryCode       string               `json:"public_phone_country_code"`
	ContactPhoneNumber           string               `json:"contact_phone_number"`
	CityID                       int64                `json:"city_id"`
	CityName                     string               `json:"city_name"`
	AddressStreet                string               `json:"address_street"`
	DirectMessaging              string               `json:"direct_messaging"`
	Latitude                     float64              `json:"latitude"`
	Longitude                    float64              `json:"longitude"`
	Category                     string               `json:"category"`
	BusinessContactMethod        string               `json:"business_contact_method"`
	IncludeDirectBlacklistStatus bool                 `json:"include_direct_blacklist_status"`
	HdProfilePicURLInfo          goinsta.PicURLInfo   `json:"hd_profile_pic_url_info"`
	HdProfilePicVersions         []goinsta.PicURLInfo `json:"hd_profile_pic_versions"`
	School                       goinsta.School       `json:"school"`
	Byline                       string               `json:"byline"`
	SocialContext                string               `json:"social_context,omitempty"`
	SearchSocialContext          string               `json:"search_social_context,omitempty"`
	MutualFollowersCount         float64              `json:"mutual_followers_count"`
	LatestReelMedia              int64                `json:"latest_reel_media,omitempty"`
	IsCallToActionEnabled        bool                 `json:"is_call_to_action_enabled"`
	FbPageCallToActionID         string               `json:"fb_page_call_to_action_id"`
	Zip                          string               `json:"zip"`
	Friendship                   goinsta.Friendship   `json:"friendship_status"`
}

func NewUser(i goinsta.User) *User {
	u := User{
		AddedAt:                      time.Now(),
		ScrapedAt:                    time.Unix(0, 0),
		ID:                           i.ID,
		Username:                     i.Username,
		FullName:                     i.FullName,
		Biography:                    i.Biography,
		ProfilePicURL:                i.ProfilePicURL,
		Email:                        i.Email,
		PhoneNumber:                  i.PhoneNumber,
		IsBusiness:                   i.IsBusiness,
		Gender:                       i.Gender,
		ProfilePicID:                 i.ProfilePicID,
		HasAnonymousProfilePicture:   i.HasAnonymousProfilePicture,
		IsPrivate:                    i.IsPrivate,
		IsUnpublished:                i.IsUnpublished,
		AllowedCommenterType:         i.AllowedCommenterType,
		IsVerified:                   i.IsVerified,
		MediaCount:                   i.MediaCount,
		FollowerCount:                i.FollowerCount,
		FollowingCount:               i.FollowingCount,
		FollowingTagCount:            i.FollowingTagCount,
		MutualFollowersID:            i.MutualFollowersID,
		ProfileContext:               i.ProfileContext,
		GeoMediaCount:                i.GeoMediaCount,
		ExternalURL:                  i.ExternalURL,
		HasBiographyTranslation:      i.HasBiographyTranslation,
		ExternalLynxURL:              i.ExternalLynxURL,
		UsertagsCount:                i.UsertagsCount,
		HasChaining:                  i.HasChaining,
		IsFavorite:                   i.IsFavorite,
		IsFavoriteForStories:         i.IsFavoriteForStories,
		IsFavoriteForHighlights:      i.IsFavoriteForHighlights,
		CanBeReportedAsFraud:         i.CanBeReportedAsFraud,
		ShowShoppableFeed:            i.ShowShoppableFeed,
		ShoppablePostsCount:          i.ShoppablePostsCount,
		ReelAutoArchive:              i.ReelAutoArchive,
		HasHighlightReels:            i.HasHighlightReels,
		PublicEmail:                  i.PublicEmail,
		PublicPhoneNumber:            i.PublicPhoneNumber,
		PublicPhoneCountryCode:       i.PublicPhoneCountryCode,
		ContactPhoneNumber:           i.ContactPhoneNumber,
		CityID:                       i.CityID,
		CityName:                     i.CityName,
		AddressStreet:                i.AddressStreet,
		DirectMessaging:              i.DirectMessaging,
		Latitude:                     i.Latitude,
		Longitude:                    i.Longitude,
		Category:                     i.Category,
		BusinessContactMethod:        i.BusinessContactMethod,
		IncludeDirectBlacklistStatus: i.IncludeDirectBlacklistStatus,
		HdProfilePicURLInfo:          i.HdProfilePicURLInfo,
		HdProfilePicVersions:         i.HdProfilePicVersions,
		School:                       i.School,
		Byline:                       i.Byline,
		SocialContext:                i.SocialContext,
		SearchSocialContext:          i.SearchSocialContext,
		MutualFollowersCount:         i.MutualFollowersCount,
		LatestReelMedia:              i.LatestReelMedia,
		IsCallToActionEnabled:        i.IsCallToActionEnabled,
		FbPageCallToActionID:         i.FbPageCallToActionID,
		Zip:                          i.Zip,
		Friendship:                   i.Friendship,
	}

	u.BiographyWithEntities.RawText = i.BiographyWithEntities.RawText
	u.BiographyWithEntities.Entities = i.BiographyWithEntities.Entities

	return &u
}

func NewGoinstaUser(u *User) *goinsta.User {
	i := goinsta.User{
		ID:       u.ID,
		Username: u.Username,
	}

	return &i
}

func (u *User) GetMeta() driver.DocumentMeta {
	return u.meta
}

func (u *User) SetMeta(meta driver.DocumentMeta) {
	u.meta = meta
}
