package redis

type (
	membersOption interface {
		applyMembers(options *membersOptions)
	}

	membersOptions struct {
		*options

		members []*member
	}
)

func defaultMembersOptions() *membersOptions {
	return &membersOptions{
		options: defaultOptions(),

		members: make([]*member, 0, 0),
	}
}
