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

func Compatible(compatibilities []Compatibility, version Application) error {
	for _, compatibility := range compatibilities {
		if compatibility.Compatible(version) == nil {
			return nil
		}
	}
	return compatibilities[0].Compatible(version)
}

func Before(compatibilities []Compatibility, version Application) bool {
	for _, compatibility := range compatibilities {
		if compatibility.Version().Before(version) {
			return true
		}
	}
	return false
}
