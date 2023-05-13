package clusterizator

func CountLimit(clz *Clusterizator, count int) bool {
	if count < 1 {
		count = 1
	}
	return clz.Clusters().Len() <= count
}

func Limit10(clz *Clusterizator) bool {
	return CountLimit(clz, 10)
}

func MustBeOne(clz *Clusterizator) bool {
	return CountLimit(clz, 1)
}
