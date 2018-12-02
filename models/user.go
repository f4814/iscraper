package models

import "github.com/ahmdrz/goinsta"

type User struct {
	ID                         int64   `bson:"id"`
	Username                   string  `bson:"username"`
	FullName                   string  `bson:"full_name"`
	Biography                  string  `bson:"biography"`
	ProfilePicURL              string  `bson:"profile_pic_url"`
	Email                      string  `bson:"email"`
	PhoneNumber                string  `bson:"phone_number"`
	IsBusiness                 bool    `bson:"is_business"`
	Gender                     int     `bson:"gender"` // XXX WHAT
	ProfilePicID               string  `bson:"profile_pic_id"`
	HasAnonymousProfilePicture bool    `bson:"has_anonymous_profile_picture"`
	IsPrivate                  bool    `bson:"is_private"`
	IsUnpublished              bool    `bson:"is_unpublished"`
	AllowedCommenterType       string  `bson:"allowed_commenter_type"`
	IsVerified                 bool    `bson:"is_verified"`
	MediaCount                 int     `bson:"media_count"`
	FollowerCount              int     `bson:"follower_count"`
	Followers                  []int64 `bson:"followers"`
	FollowingCount             int     `bson:"following_count"`
	Following                  []int64 `bson:"following"`
	FollowingTagCount          int     `bson:"following_tag_count"`
	ProfileContext             string  `bson:"profile_context"`
	GeoMediaCount              int     `bson:"geo_media_count"`
	ExternalURL                string  `bson:"external_url"`
	HasBiographyTranslation    bool    `bson:"has_biography_translation"`
	ExternalLynxURL            string  `bson:"external_lynx_url"`
	BiographyWithEntities      struct {
		RawText  string        `bson:"raw_text"`
		Entities []interface{} `bson:"entities"`
	} `bson:"biography_with_entities"`
	UsertagsCount                int      `bson:"usertags_count"`
	HasChaining                  bool     `bson:"has_chaining"`
	CanBeReportedAsFraud         bool     `bson:"can_be_reported_as_fraud"`
	ShowShoppableFeed            bool     `bson:"show_shoppable_feed"`
	ShoppablePostsCount          int      `bson:"shoppable_posts_count"`
	ReelAutoArchive              string   `bson:"reel_auto_archive"`
	HasHighlightReels            bool     `bson:"has_highlight_reels"`
	PublicEmail                  string   `bson:"public_email"`
	PublicPhoneNumber            string   `bson:"public_phone_number"`
	PublicPhoneCountryCode       string   `bson:"public_phone_country_code"`
	ContactPhoneNumber           string   `bson:"contact_phone_number"`
	CityID                       int64    `bson:"city_id"`
	CityName                     string   `bson:"city_name"`
	AddressStreet                string   `bson:"address_street"`
	DirectMessaging              string   `bson:"direct_messaging"`
	Latitude                     float64  `bson:"latitude"`
	Longitude                    float64  `bson:"longitude"`
	Category                     string   `bson:"category"`
	BusinessContactMethod        string   `bson:"business_contact_method"`
	IncludeDirectBlacklistStatus bool     `bson:"include_direct_blacklist_status"`
	HdProfilePicVersions         []string `bson:"hd_profile_pic_versions"`
	Byline                       string   `bson:"byline"`
	SocialContext                string   `bson:"social_context,omitempty"`
	SearchSocialContext          string   `bson:"search_social_context,omitempty"`
	MutualFollowersCount         float64  `bson:"mutual_followers_count"`
	LatestReelMedia              int64    `bson:"latest_reel_media,omitempty"`
	IsCallToActionEnabled        bool     `bson:"is_call_to_action_enabled"`
	FbPageCallToActionID         string   `bson:"fb_page_call_to_action_id"`
	Zip                          string   `bson:"zip"`
}

// Transforms an gointsta.User into an models.User
// Expects the user to be updated
// Does not scrape followers or following
func (m *User) FromIG(u *goinsta.User) {
	m.ID = u.ID
	m.Username = u.Username
	m.FullName = u.FullName
	m.Biography = u.Biography
	m.ProfilePicURL = u.ProfilePicURL
	m.Email = u.Email
	m.PhoneNumber = u.PhoneNumber
	m.IsBusiness = u.IsBusiness
	m.Gender = u.Gender
	m.ProfilePicID = u.ProfilePicID
	m.HasAnonymousProfilePicture = u.HasAnonymousProfilePicture
	m.IsPrivate = u.IsPrivate
	m.IsUnpublished = u.IsUnpublished
	m.AllowedCommenterType = u.AllowedCommenterType
	m.IsVerified = u.IsVerified
	m.MediaCount = u.MediaCount
	m.FollowerCount = u.FollowerCount
	m.FollowingCount = u.FollowingCount
	m.FollowingTagCount = u.FollowingTagCount
	m.ProfileContext = u.ProfileContext
	m.GeoMediaCount = u.GeoMediaCount
	m.ExternalURL = u.ExternalURL
	m.HasBiographyTranslation = u.HasBiographyTranslation
	m.ExternalLynxURL = u.ExternalLynxURL
	m.UsertagsCount = u.UsertagsCount
	m.HasChaining = u.HasChaining
	m.CanBeReportedAsFraud = u.CanBeReportedAsFraud
	m.ShowShoppableFeed = u.ShowShoppableFeed
	m.ShoppablePostsCount = u.ShoppablePostsCount
	m.ReelAutoArchive = u.ReelAutoArchive
	m.HasHighlightReels = u.HasHighlightReels
	m.PublicEmail = u.PublicEmail
	m.PublicPhoneNumber = u.PublicPhoneNumber
	m.PublicPhoneCountryCode = u.PublicPhoneCountryCode
	m.ContactPhoneNumber = u.ContactPhoneNumber
	m.CityID = u.CityID
	m.CityName = u.CityName
	m.AddressStreet = u.AddressStreet
	m.DirectMessaging = u.DirectMessaging
	m.Latitude = u.Latitude
	m.Longitude = u.Longitude
	m.Category = u.Category
	m.BusinessContactMethod = u.BusinessContactMethod
	m.IncludeDirectBlacklistStatus = u.IncludeDirectBlacklistStatus
	m.Byline = u.Byline
	m.SocialContext = u.SocialContext
	m.SearchSocialContext = u.SearchSocialContext
	m.MutualFollowersCount = u.MutualFollowersCount
	m.LatestReelMedia = u.LatestReelMedia
	m.IsCallToActionEnabled = u.IsCallToActionEnabled
	m.FbPageCallToActionID = u.FbPageCallToActionID
	m.Zip = u.Zip

	m.BiographyWithEntities.RawText = u.BiographyWithEntities.RawText
	m.BiographyWithEntities.Entities = u.BiographyWithEntities.Entities

	for _, u := range u.HdProfilePicVersions {
		m.HdProfilePicVersions = append(m.HdProfilePicVersions, u.URL)
	}
}
