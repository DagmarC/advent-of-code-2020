package adapterpkg

import (
	"fmt"
	"sort"
)

func DistributeAdapters(allAdapters *[]Adapter) (JoltAdaptersMap, error) {
	joltDifferences := make(JoltAdaptersMap, 0)
	joltDifferences.InitializeDistribution()

	// SORT ARRAY
	sort.SliceStable(*allAdapters, func(i, j int) bool {
		return (*allAdapters)[i] < (*allAdapters)[j]
	})
	// DISTRIBUTE
	createDistribution(&joltDifferences, allAdapters)
	fmt.Println("Distributions", joltDifferences)

	return joltDifferences, nil
}

// createDistribution goes through allAdapters and add it to the map of distribution - joltDifferences (One, Two, Three Jolt keys),
// where each adapter rating is added to one of these distributions.
func createDistribution(joltDifferences *JoltAdaptersMap, allAdapters *[]Adapter) {

	for i := 0; i < len(*allAdapters); i++ {

		adapterRating := (*allAdapters)[i]
		var distribution = Undefined

		if i == len(*allAdapters)-1 {
			// your device's built-in adapter is always 3 higher than the highest adapter
			deviceAdapterRating := adapterRating + 3
			distribution = adapterRating.Difference(deviceAdapterRating)
		} else {
			nextAdapterRating := (*allAdapters)[i+1]
			distribution = adapterRating.Difference(nextAdapterRating)
		}
		(*joltDifferences)[distribution] = append((*joltDifferences)[distribution], adapterRating)
	}
}
