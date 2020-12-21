package main

import (
	"testing"
)

func TestShortestDistance(t *testing.T) {
	var Neighborhood = []Street{
		Street{From: "Kruthika's abode", To: "Mark's crib", Distance: 9},
		Street{From: "Kruthika's abode", To: "Greg's casa", Distance: 4},
		Street{From: "Kruthika's abode", To: "Matt's pad", Distance: 18},
		Street{From: "Kruthika's abode", To: "Brian's apartment", Distance: 8},
		Street{From: "Brian's apartment", To: "Wesley's condo", Distance: 7},
		Street{From: "Brian's apartment", To: "Cam's dwelling", Distance: 17},
		Street{From: "Greg's casa", To: "Cam's dwelling", Distance: 13},
		Street{From: "Greg's casa", To: "Mike's digs", Distance: 19},
		Street{From: "Greg's casa", To: "Matt's pad", Distance: 14},
		Street{From: "Wesley's condo", To: "Kirk's farm", Distance: 10},
		Street{From: "Wesley's condo", To: "Nathan's flat", Distance: 11},
		Street{From: "Wesley's condo", To: "Bryce's den", Distance: 6},
		Street{From: "Matt's pad", To: "Mark's crib", Distance: 19},
		Street{From: "Matt's pad", To: "Nathan's flat", Distance: 15},
		Street{From: "Matt's pad", To: "Craig's haunt", Distance: 14},
		Street{From: "Mark's crib", To: "Kirk's farm", Distance: 9},
		Street{From: "Mark's crib", To: "Nathan's flat", Distance: 12},
		Street{From: "Bryce's den", To: "Craig's haunt", Distance: 10},
		Street{From: "Bryce's den", To: "Mike's digs", Distance: 9},
		Street{From: "Mike's digs", To: "Cam's dwelling", Distance: 20},
		Street{From: "Mike's digs", To: "Nathan's flat", Distance: 12},
		Street{From: "Cam's dwelling", To: "Craig's haunt", Distance: 18},
		Street{From: "Nathan's flat", To: "Kirk's farm", Distance: 3},
	}

	startLocation := "Kruthika's abode"
	targetLocation := "Craig's haunt"

	// formats the data into a map so the input array is looped over only once
	NodeMap := formatData(Neighborhood)

	// holds PathData for each location
	LocationPathData := initPathDataMap(NodeMap, &startLocation)

	// loops over the NodeMap and determines the shortest route from the startLocation to all other locations
	findPath(startLocation, LocationPathData, NodeMap)

	shortestDistance := LocationPathData[targetLocation].ShortestDistance

	if shortestDistance != 31 {
		t.Errorf("Shortest distance was incorrect, got: %v, want: %v.", shortestDistance, 31)
	}
}

func TestPathResultString(t *testing.T) {
	var Neighborhood = []Street{
		Street{From: "Kruthika's abode", To: "Mark's crib", Distance: 9},
		Street{From: "Kruthika's abode", To: "Greg's casa", Distance: 4},
		Street{From: "Kruthika's abode", To: "Matt's pad", Distance: 18},
		Street{From: "Kruthika's abode", To: "Brian's apartment", Distance: 8},
		Street{From: "Brian's apartment", To: "Wesley's condo", Distance: 7},
		Street{From: "Brian's apartment", To: "Cam's dwelling", Distance: 17},
		Street{From: "Greg's casa", To: "Cam's dwelling", Distance: 13},
		Street{From: "Greg's casa", To: "Mike's digs", Distance: 19},
		Street{From: "Greg's casa", To: "Matt's pad", Distance: 14},
		Street{From: "Wesley's condo", To: "Kirk's farm", Distance: 10},
		Street{From: "Wesley's condo", To: "Nathan's flat", Distance: 11},
		Street{From: "Wesley's condo", To: "Bryce's den", Distance: 6},
		Street{From: "Matt's pad", To: "Mark's crib", Distance: 19},
		Street{From: "Matt's pad", To: "Nathan's flat", Distance: 15},
		Street{From: "Matt's pad", To: "Craig's haunt", Distance: 14},
		Street{From: "Mark's crib", To: "Kirk's farm", Distance: 9},
		Street{From: "Mark's crib", To: "Nathan's flat", Distance: 12},
		Street{From: "Bryce's den", To: "Craig's haunt", Distance: 10},
		Street{From: "Bryce's den", To: "Mike's digs", Distance: 9},
		Street{From: "Mike's digs", To: "Cam's dwelling", Distance: 20},
		Street{From: "Mike's digs", To: "Nathan's flat", Distance: 12},
		Street{From: "Cam's dwelling", To: "Craig's haunt", Distance: 18},
		Street{From: "Nathan's flat", To: "Kirk's farm", Distance: 3},
	}

	startLocation := "Kruthika's abode"
	targetLocation := "Craig's haunt"

	// formats the data into a map so the input array is looped over only once
	NodeMap := formatData(Neighborhood)

	// holds PathData for each location
	LocationPathData := initPathDataMap(NodeMap, &startLocation)

	// loops over the NodeMap and determines the shortest route from the startLocation to all other locations
	findPath(startLocation, LocationPathData, NodeMap)

	path := ReturnPath(LocationPathData, targetLocation)
	correctPath := "[\"Kruthika's abode\", \"Brian's apartment\", \"Wesley's condo\", \"Bryce's den\", \"Craig's haunt\"]"

	if path != correctPath {
		t.Errorf("Shortest distance was incorrect, got: %v, want: %v.", path, correctPath)
	}

}