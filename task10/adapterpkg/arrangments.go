package adapterpkg

import "github.com/DagmarC/advent-of-code-2020/task9/dataqueue"

// createAdapterArrangements will create map where key is the Adapter Rating in Jolt and value is the slice of all
// possible adapters to which it can be joined (which means max 3 numbers with difference of 1, 2 or 3 from the key).
// if we have input slice [0, 1, 3, 5] then the map at key 0  would have [1, 3] values -->
// Result: 0:[1, 3], 1:[3], 3:[5], 5:[]
func createAdapterArrangements(allAdapters *[]Adapter) (JoltAdaptersMap, error) {
	
	adaptersGraph := make(JoltAdaptersMap, 0)
	err := adaptersGraph.InitializeAdaptersMap(allAdapters)
	if err != nil {
		return nil, err
	}

	err = createAllArrangements(&adaptersGraph, allAdapters)
	if err != nil {
		return nil, err
	}

	return adaptersGraph, nil
}

// createAllArrangements will fill the structure adaptersGraph on input allAdapters [0, 1, 3, 5]
// Result adaptersGraph: 0:[1, 3], 1:[3], 3:[5], 5:[]
func createAllArrangements(adaptersGraph *JoltAdaptersMap, allAdapters *[]Adapter) error {
	vertices := dataqueue.Queue{}

	for _, adapter := range *allAdapters {

		for _, vertex := range vertices.Elements() {
			diff := int(adapter) - vertex
			if diff <= 3 && diff > 0 {
				// 0 -> 1, where adapter is 1 and vertex is 0, so it will be g[0] = [1]
				(*adaptersGraph)[Jolt(vertex)] = append((*adaptersGraph)[Jolt(vertex)], adapter)

				// If the diff is 3, then the next vertex will never have the diff in (0,3> range.
				if diff == 3 {
					_, err := vertices.Dequeue()
					if err != nil {
						return err
					}
				}
			} else {
				// Deque current vertex if the difference is different. It should be the 0th element in the queue.
				_, err := vertices.Dequeue()
				if err != nil {
					return err
				}
			}
		}
		vertices.Enqueue(int(adapter))
	}
	return nil
}
