package clusterizator

func CountLimit(clz *Clusterizator, count int) bool {
	if count < 1 {
		count = 1
	}
	return len(clz.Clusters()) <= count
}

func Limit10(clz *Clusterizator) bool {
	return CountLimit(clz, 10)
}

func MustBeOne(clz *Clusterizator) bool {
	return CountLimit(clz, 1)
}
