package business

import (
	"fmt"
	"github.com/crit/yelp"
)

// SearchRequest is all the options listed at https://docs.developer.yelp.com/reference/v3_business_search
type SearchRequest struct {
	Location   string   `url:"location,omitempty"` // "NYC" "New York City"
	Latitude   int      `url:"latitude,omitempty"`
	Longitude  int      `url:"longitude,omitempty"`
	Term       string   `url:"term,omitempty"` // "food" "restaurant"
	Radius     int      `url:"radius,omitempty"`
	Categories []string `url:"categories,omitempty"` // "bars,french" will filter by Bars OR French
	Locale     string   `url:"locale,omitempty"`     // {language}_{country code}
	Price      string   `url:"price,omitempty"`      // options 1, 2, 3, 4. "1,2,3" will show $, $$, or $$$ options
	OpenNow    bool     `url:"open_now,omitempty"`
	OpenAt     int      `url:"open_at,omitempty"` // unix timestamp

	// hot_and_new - popular businesses which recently joined Yelp
	// request_a_quote - businesses which actively reply to Request a Quote inquiries
	// reservation - businesses with Yelp Reservations bookings enabled on their profile page
	// waitlist_reservation - businesses with Yelp Wait List bookings enabled on their profile screen (iOS/Android)
	// deals - businesses offering Yelp Deals on their profile page
	// gender_neutral_restrooms - businesses which provide gender neutral restrooms
	// open_to_all - businesses which are Open To All
	// wheelchair_accessible - businesses which are Wheelchair Accessible
	Attributes []string `url:"attributes,omitempty"`

	// Suggestion to the search algorithm that the results be sorted by one of the these modes: best_match, rating, review_count or distance.
	// The default is best_match. Note that specifying the sort_by is a suggestion (not strictly enforced) to Yelp's search, which considers multiple input parameters to return the most relevant results.
	//
	// e.g., the rating sort is not strictly sorted by the rating value, but by an adjusted rating value that takes into account the number of ratings,
	// similar to a Bayesian average. This is to prevent skewing results to businesses with a single review.
	SortBy string `url:"sort_by,omitempty"`

	DevicePlatform string `url:"device_platform,omitempty"` // "android", "ios", "mobile-generic"

	ReservationDate       string `url:"reservation_date,omitempty"`         // YYYY-mm-dd
	ReservationTime       string `url:"reservation_time,omitempty"`         // HH:MM
	ReservationCovers     int    `url:"reservation_covers,omitempty"`       // how many attending
	MatchesPartySizeParam bool   `url:"matches_party_size_param,omitempty"` // filter out results that don't have an opening?

	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type SearchResult struct {
	Businesses []Business `json:"businesses"`
	Region     Region     `json:"region"`
	Total      int        `json:"total"`
}

func Search(client *yelp.Client, req SearchRequest) (*SearchResult, error) {
	var res SearchResult

	err := client.Get("/businesses/search", req, &res)
	if err != nil {
		return nil, fmt.Errorf("unable to search for businesses; %w", err)
	}

	return &res, nil
}
