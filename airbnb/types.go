package airbnb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"strings"
	"time"
)

type ListingAndReviews struct {
	Id                   string               `bson:"_id,omitempty"`
	ListingUrl           string               `bson:"listing_url"`
	Name                 string               `bson:"name"`
	Summary              string               `bson:"summary"`
	Space                string               `bson:"space,omitempty"`
	Description          string               `bson:"description,omitempty"`
	NeighborhoodOverview string               `bson:"neighborhood_overview,omitempty"`
	Notes                string               `bson:"notes,omitempty"`
	Transit              string               `bson:"transit,omitempty"`
	Access               string               `bson:"access,omitempty"`
	Interaction          string               `bson:"interaction,omitempty"`
	HouseRules           string               `bson:"house_rules,omitempty"`
	PropertyType         string               `bson:"property_type,omitempty"`
	RoomType             string               `bson:"room_type,omitempty"`
	BedType              string               `bson:"bed_type,omitempty"`
	MinimumNights        string               `bson:"minimum_nights,omitempty"`
	MaximumNights        string               `bson:"maximum_nights,omitempty"`
	CancellationPolicy   string               `bson:"cancellation_policy,omitempty"`
	LastScraped          time.Time            `bson:"last_scraped,omitempty"`
	CalendarLastScraped  time.Time            `bson:"calendar_last_scraped,omitempty"`
	FirstReview          time.Time            `bson:"first_review,omitempty"`
	LastReview           time.Time            `bson:"last_review,omitempty"`
	Accommodates         int                  `bson:"accommodates,omitempty"`
	Bedrooms             int                  `bson:"bedrooms,omitempty"`
	Beds                 int                  `bson:"beds,omitempty"`
	NumberOfReviews      int                  `bson:"number_of_reviews,omitempty"`
	Bathrooms            primitive.Decimal128 `bson:"bathrooms,omitempty"`
	Amenities            []string             `bson:"amenities,omitempty"`
	Price                primitive.Decimal128 `bson:"price,omitempty"`
	CleaningFee          primitive.Decimal128 `bson:"cleaning_fee,omitempty"`
	ExtraPeople          primitive.Decimal128 `bson:"extra_people,omitempty"`
	GuestsIncluded       primitive.Decimal128 `bson:"guests_included,omitempty"`
	Images               struct {
		ThumbnailUrl string `bson:"thumbnail_url,omitempty"`
		MediumUrl    string `bson:"medium_url,omitempty"`
		PictureUrl   string `bson:"picture_url,omitempty"`
		XlPictureUrl string `bson:"xl_picture_url,omitempty"`
	} `bson:"images,omitempty"`
	Host struct {
		HostId                 string   `bson:"host_id,omitempty"`
		HostUrl                string   `bson:"host_url,omitempty"`
		HostName               string   `bson:"host_name,omitempty"`
		HostLocation           string   `bson:"host_location,omitempty"`
		HostAbout              string   `bson:"host_about,omitempty"`
		HostThumbnailUrl       string   `bson:"host_thumbnail_url,omitempty"`
		HostPictureUrl         string   `bson:"host_picture_url,omitempty"`
		HostNeighbourhood      string   `bson:"host_neighbourhood,omitempty"`
		HostIsSuperhost        bool     `bson:"host_is_superhost,omitempty"`
		HostHasProfilePic      bool     `bson:"host_has_profile_pic,omitempty"`
		HostIdentityVerified   bool     `bson:"host_identity_verified,omitempty"`
		HostListingsCount      int      `bson:"host_listings_count,omitempty"`
		HostTotalListingsCount int      `bson:"host_total_listings_count,omitempty"`
		HostVerifications      []string `bson:"host_verifications,omitempty"`
	} `bson:"host,omitempty"`
	Address struct {
		Street         string `bson:"street,omitempty"`
		Suburb         string `bson:"suburb,omitempty"`
		GovernmentArea string `bson:"government_area,omitempty"`
		Market         string `bson:"market,omitempty"`
		Country        string `bson:"country,omitempty"`
		CountryCode    string `bson:"country_code,omitempty"`
		Location       struct {
			Type            string    `bson:"type,omitempty"`
			Coordinates     []float64 `bson:"coordinates,omitempty"`
			IsLocationExact bool      `bson:"is_location_exact,omitempty"`
		} `bson:"location,omitempty"`
	} `bson:"address,omitempty"`
	Availability struct {
		Availability30  int `bson:"availability_30,omitempty"`
		Availability60  int `bson:"availability_60,omitempty"`
		Availability90  int `bson:"availability_90,omitempty"`
		Availability365 int `bson:"availability_365,omitempty"`
	} `bson:"availability,omitempty"`
	ReviewScores struct {
		ReviewScoresAccuracy      int `bson:"review_scores_accuracy,omitempty"`
		ReviewScoresCleanliness   int `bson:"review_scores_cleanliness,omitempty"`
		ReviewScoresCheckin       int `bson:"review_scores_checkin,omitempty"`
		ReviewScoresCommunication int `bson:"review_scores_communication,omitempty"`
		ReviewScoresLocation      int `bson:"review_scores_location,omitempty"`
		ReviewScoresValue         int `bson:"review_scores_value,omitempty"`
		ReviewScoresRating        int `bson:"review_scores_rating,omitempty"`
	} `bson:"review_scores,omitempty"`
	Reviews []struct {
		Id           string    `bson:"_id,omitempty"`
		Date         time.Time `bson:"date,omitempty"`
		ListingId    string    `bson:"listing_id,omitempty"`
		ReviewerId   string    `bson:"reviewer_id,omitempty"`
		ReviewerName string    `bson:"reviewer_name,omitempty"`
		Comments     string    `bson:"comments,omitempty"`
	} `bson:"reviews,omitempty"`
}

var nameElements = []string{
	"Gungy",
	"Dodge",
	"Dirty",
	"Damp",
	"Nasty",
	"Rotten",
	"Rough",
	"Filthy",
	"Flea-Ridden",
	"Mouldy",
	"Stinky",
	"Disgusting",
}

// randomName returns a random name from a given slice of name elements.
func randomName(nameElements []string) string {
	rand.Seed(time.Now().UnixNano())
	sb := new(strings.Builder)
	length := rand.Intn(len(nameElements))

	cpy := make([]string, len(nameElements))
	copy(cpy, nameElements)

	for i := 0; i < length-2; i++ {
		index := rand.Intn(len(cpy))
		sb.WriteString(cpy[index])
		//remove the element at index from the slice
		cpy = append(cpy[:index], cpy[index+1:]...)
		sb.WriteString(", ")
	}
	index := rand.Intn(len(cpy))
	sb.WriteString(cpy[index] + " and ")
	cpy = append(cpy[:index], cpy[index+1:]...)
	index = rand.Intn(len(cpy))
	sb.WriteString(cpy[index])

	return sb.String()
}
