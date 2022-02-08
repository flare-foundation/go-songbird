package version

func Unmaskable(compatibilities []Compatibility, version Application) error {
	for _, compatibility := range compatibilities {
		if compatibility.Unmaskable(version) == nil {
			return nil
		}
	}
	return compatibilities[0].Unmaskable(version)
}

func WontMask(compatibilities []Compatibility, version Application) error {
	for _, compatibility := range compatibilities {
		if compatibility.WontMask(version) == nil {
			return nil
		}
	}
	return compatibilities[0].WontMask(version)
}
