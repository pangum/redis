package redis

func mustInt(first int, secondary int) (final int) {
	if 0 != first {
		final = first
	} else {
		final = secondary
	}

	return
}

func mustString(first string, secondary string) (final string) {
	if "" != first {
		final = first
	} else {
		final = secondary
	}

	return
}
