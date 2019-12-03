package norms

import "math"

type (
	NormResources struct {
		Resource int
		Count    int
	}
	BaseNorm struct {
		ID               int
		Norms            []NormResources
		RewardReputation int
		RewardCredits    int
	}
)

var (
	allNorms = []BaseNorm{
		{
			ID: 1,
			Norms: []NormResources{
				{
					Resource: 8,
					Count:    50,
				},
				{
					Resource: 10,
					Count:    51,
				},
			},
			RewardReputation: 2,
			RewardCredits:    20,
		},
		{
			ID: 2,
			Norms: []NormResources{
				{
					Resource: 8,
					Count:    55,
				},
				{
					Resource: 10,
					Count:    59,
				},
			},
			RewardReputation: 2,
			RewardCredits:    25,
		},
		{
			ID: 3,
			Norms: []NormResources{
				{
					Resource: 8,
					Count:    63,
				},
				{
					Resource: 10,
					Count:    66,
				},
			},
			RewardReputation: 2,
			RewardCredits:    30,
		},
		{
			ID: 4,
			Norms: []NormResources{
				{
					Resource: 8,
					Count:    71,
				},
				{
					Resource: 10,
					Count:    74,
				},
			},
			RewardReputation: 2,
			RewardCredits:    40,
		},
		{
			ID: 5,
			Norms: []NormResources{
				{
					Resource: 8,
					Count:    77,
				},
				{
					Resource: 10,
					Count:    82,
				},
			},
			RewardReputation: 2,
			RewardCredits:    40,
		},
		{
			ID: 6,
			Norms: []NormResources{
				{
					Resource: 8,
					Count:    80,
				},
				{
					Resource: 10,
					Count:    89,
				},
			},
			RewardReputation: 2,
			RewardCredits:    50,
		},
		{
			ID: 7,
			Norms: []NormResources{
				{
					Resource: 8,
					Count:    88,
				},
				{
					Resource: 10,
					Count:    91,
				},
			},
			RewardReputation: 2,
			RewardCredits:    100,
		},
	}
)

func GetAllNorms() []BaseNorm {
	return allNorms
}

func GetMaxNormID() (max int) {
	for _, it := range allNorms {
		max = int(math.Max(float64(max), float64(it.ID)))
	}
	return max
}

func GetNormByID(id int) *BaseNorm {
	for _, it := range allNorms {
		if it.ID != id {
			continue
		}
		return &it
	}
	var maxID int
	for _, it := range allNorms {
		maxID = int(math.Max(float64(maxID), float64(it.ID)))
	}
	for _, it := range allNorms {
		if it.ID != maxID {
			continue
		}
		return &it
	}
	return nil
}
