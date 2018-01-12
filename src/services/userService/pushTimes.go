package userService

func AddPushTimes(userId string, timerId string) error {
	oldUi, err := ReadUserInfo(userId)
	if err != nil {
		return err
	}

	pt := oldUi.PushTimes
	pt = append(pt, timerId)
	ui := NewUserInfo(
		userId,
		oldUi.ConsecutiveActionDays,
		oldUi.LastActionTime,
		pt,
	)
	if err := UpdateUserInfo(userId, ui); err != nil {
		return err
	}

	return nil
}
