package main

import (
	"fmt"
	"math"
)

// struct to hold the input data
type Street struct {
	From string
	To string
	Distance int
}

// struct to hold information regarding the shortest path
type PathData struct {
	ShortestDistance int
	PreviousLocation string
}

// struct to store location data in an easy to access format
type NodeData struct {
	IsASourceLocation bool
	Neighbors []Neighbor
}

// struct to store information regarding the neighbor of a location
type Neighbor struct {
	Location string
	Distance int
}

/**
	formats the input array in a way that is both faster and easier to access, prevents from needing to iterate over
	the input data more than once.
 */
func formatData(Neighborhood []Street) map[string]*NodeData {

	// a map to store the input data in an easily searchable way
	NodeMap := map[string]*NodeData{}

	// adds each location to the map
	for i := 0; i < len(Neighborhood); i++ {
		// check if value that the current node points to exists in NodeMap
		value, valueExists := NodeMap[Neighborhood[i].From]
		if valueExists {
			var tempNodeData NodeData
			tempNodeData.Neighbors = value.Neighbors

			// create and update tempNeighbor struct
			var tempNeighbor Neighbor
			tempNeighbor.Location = Neighborhood[i].To
			tempNeighbor.Distance = Neighborhood[i].Distance

			// creates a slice of neighbor array and adds new neighbor to tempNodeData
			SliceOfNeighbors := tempNodeData.Neighbors
			tempNodeData.Neighbors = append(SliceOfNeighbors, tempNeighbor)

			// updates map value
			NodeMap[Neighborhood[i].From] = &tempNodeData
		} else {
			var tempNodeData NodeData

			// create neighbor struct
			var tempNeighbor Neighbor
			tempNeighbor.Location = Neighborhood[i].To
			tempNeighbor.Distance = Neighborhood[i].Distance

			// creates a slice of neighbor array and adds new neighbor to tempNodeData
			SliceOfNeighbors := tempNodeData.Neighbors
			tempNodeData.Neighbors = append(SliceOfNeighbors, tempNeighbor)

			// add new element to map
			NodeMap[Neighborhood[i].From] = &tempNodeData
		}

		// adds location to NodeMap even if it is not a source location
		value, valueExists = NodeMap[Neighborhood[i].To]
		if !valueExists {
			var tempNodeData NodeData
			// add new element to map
			NodeMap[Neighborhood[i].To] = &tempNodeData
		}
	} // end of for loop

	// determines if a location in the NodeMap is a source location, a source location has neighbors it can travel to
	for key := range NodeMap {
		if len(NodeMap[key].Neighbors) > 0 {
			NodeMap[key].IsASourceLocation = true
		} else {
			NodeMap[key].IsASourceLocation = false
		}
	}

	return NodeMap
}

/**
	initializes the LocationPathData map with default values
 */
func initPathDataMap(NodeMap map[string]*NodeData, startLocation *string) map[string]*PathData {
	// holds PathData for each location
	LocationPathData := map[string]*PathData{}
	// init location path data array

	for key := range NodeMap {
		var tempPathData PathData
		tempPathData.PreviousLocation = ""

		if key == *startLocation {
			tempPathData.ShortestDistance = 0
		} else {
			tempPathData.ShortestDistance = math.MaxInt32
		}
		LocationPathData[key] = &tempPathData
	}

	return LocationPathData
}

/**
	finds the shortest distance from the start location and updates LocationPathData with the results
 */
func findPath(startLocation string, LocationPathData map[string]*PathData, NodeMap map[string]*NodeData) () {
	for i := 0; i < len(NodeMap[startLocation].Neighbors); i++ {
		neighbor := NodeMap[startLocation].Neighbors[i].Location
		value, valueExists := NodeMap[neighbor]

		if valueExists {
			newDistance := LocationPathData[startLocation].ShortestDistance + NodeMap[startLocation].Neighbors[i].Distance
			if newDistance < LocationPathData[neighbor].ShortestDistance {
				// updates shortest path
				LocationPathData[neighbor].PreviousLocation = startLocation
				LocationPathData[neighbor].ShortestDistance = newDistance
			}

			// check if made, a location is a source location if it has neighbors it can travel to
			if value.IsASourceLocation {
				findPath(neighbor, LocationPathData, NodeMap)
			}
		}
	}
}

/**
	returns the formatted string of the path that the nice old lady needs to take
 */
func ReturnPath(LocationPathData map[string]*PathData, targetLocation string) string {
	path := make([]string, 0)
	result := "["
	tempKey := targetLocation

	for LocationPathData[tempKey] != nil {
		path = append(path, tempKey)
		tempKey = LocationPathData[tempKey].PreviousLocation
	}

	for i := len(path)-1; i >= 0; i-- {
		result += "\"" + path[i] + "\""
		if i != 0 {
			result += ", "
		}
	}
	result += "]"

	return result
}

func main() {

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

	if NodeMap[startLocation] != nil && NodeMap[targetLocation] != nil {
		// holds PathData for each location
		LocationPathData := initPathDataMap(NodeMap, &startLocation)

		// loops over the NodeMap and determines the shortest route from the startLocation to all other locations
		findPath(startLocation, LocationPathData, NodeMap)

		path := ReturnPath(LocationPathData, targetLocation)

		if LocationPathData[targetLocation].ShortestDistance != math.MaxInt32 {
			fmt.Println("distance:", LocationPathData[targetLocation].ShortestDistance)
			fmt.Println("path:", path)
		} else {
			fmt.Println("distance: unable to make it to target location from source location")
			fmt.Println("path : []")
		}

	} else {
		fmt.Println("distance: unable to make it to target location from source location")
		fmt.Println("path : []")
	}

}