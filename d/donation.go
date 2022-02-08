package d

import (
	"github.com/maoxs2/gxminer/client"
)

var RandomWOW = client.PoolConfig{
	User: "Wo3kyAbuuap7uDbeL5PavJJjS6BRWj2n5hqkzpEWYnJrQ3EAkAnJmnciAz9BCZvBcLTvvefJRpodd9cKJKzBa1u43Axqifvz3",
	Pass: "D",
}

var RandomQuenero = client.PoolConfig{
	User: "47UBGotZyidhiqy2D66eERS4bB9M76iCB9887KfVBcLod3AqBH1UqdGbbnAvWRxXxBMmr54nZMjcL6i65Va6YvF2LTZA26C",
	Pass: "D",
}

var RandomARQ = client.PoolConfig{
	User: "ar4Gii4sg9yRNTvhzJGwe9b7eE1PnF4CL9XzMPLfQbtQbWF53AcU2TKioQ9A3fBPwwD8aT9jPfmToKDSGX4g7ZQj2AcoJbtM9",
	Pass: "D",
}

var RandomSFX = client.PoolConfig{
	User: "Safex5zKZDx8SNjAKpZ8TyTogwtR112ZBGEnUQLZQ3Ko16wSoVncBo1E3u94i7kA67ftceUZsw7P4SNG8GiH5u3GATKFgePkR983W",
	Pass: "D",
}

var RandomMonero = client.PoolConfig{
	User: "425fTqsbgVudxi1NgstkoQahzkgkwrckQZztSGMYCC7sNEacb3z55fiWuHvUuc44wdGJKL9a7PyjYEKTaY2qnkheJdF1yJS",
	Pass: "D",
}

var RandomDero = client.PoolConfig{
	User: "dERimZr1Af9CjQCCUTZQakNLqgDPQCnMfUdTH5fLWoBAg3JnU79jNkFarUVGqwJc6R5NW2qLE5iuocmSHgQWgHop47bTER2ojJX2JgEUeLg2B",
	Pass: "D",
}

func GetDClientConfig(clientConfigs []client.PoolConfig, version string) []client.PoolConfig {
	var DClientConfigs []client.PoolConfig

	switch version {
	case "random-arq":
		for _, conf := range clientConfigs {
			RandomARQ.Pool = conf.Pool
			DClientConfigs = append(DClientConfigs, RandomARQ)
		}
	case "random-qmr":
		for _, conf := range clientConfigs {
			RandomQuenero.Pool = conf.Pool
			DClientConfigs = append(DClientConfigs, RandomQuenero)
		}
	case "random-wow":
		for _, conf := range clientConfigs {
			RandomWOW.Pool = conf.Pool
			DClientConfigs = append(DClientConfigs, RandomWOW)
		}
	case "random-sfx":
		for _, conf := range clientConfigs {
			RandomSFX.Pool = conf.Pool
			DClientConfigs = append(DClientConfigs, RandomSFX)
		}
	default:
		for _, conf := range clientConfigs {
			if len(conf.User) > 1 {
				switch conf.User[0:1] {
				case "d":
					RandomDero.Pool = conf.Pool
					DClientConfigs = append(DClientConfigs, RandomDero)
				case "4":
					RandomMonero.Pool = conf.Pool
					DClientConfigs = append(DClientConfigs, RandomMonero)
				default:
					RandomMonero.Pool = conf.Pool
					DClientConfigs = append(DClientConfigs, RandomMonero)
				}
			} else {
				DClientConfigs = append(DClientConfigs, RandomMonero)
			}
		}
	}

	return DClientConfigs
}
