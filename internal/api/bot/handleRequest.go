package bot

// 	} else if commandAccountTotals.MatchString(command) {

// 		result := commandAccountTotals.FindAllStringSubmatch(command, -1)

// 		accountId := result[0][1]

// 		err = HandleAccountTotals(user, accountId)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle account totals: %s", err.Error()))
// 		}

// 	} else if commandAccountDetail.MatchString(command) {

// 		result := commandAccountDetail.FindAllStringSubmatch(command, -1)

// 		accountId := result[0][1]

// 		err = HandleAccountDetail(user, accountId)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle account detail: %s", err.Error()))
// 		}

// 	} else if commandAccountPositions.MatchString(command) {

// 		result := commandAccountPositions.FindAllStringSubmatch(command, -1)

// 		accountId := result[0][1]

// 		err = HandleAccountPositions(user, accountId)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle account positions: %s", err.Error()))
// 		}

// 	} else if commandAccountPosition.MatchString(command) {

// 		result := commandAccountPosition.FindAllStringSubmatch(command, -1)

// 		accountId := result[0][1]

// 		if isCommand(message.Text) {
// 			err = HandleAccountPosition(user, accountId, "")
// 			if err != nil {
// 				logger.Error(errors.Errorf("error on handle account position: %s", err.Error()))
// 			}
// 		} else {
// 			err = HandleAccountPosition(user, accountId, message.Text)
// 			if err != nil {
// 				logger.Error(errors.Errorf("error on handle account position: %s", err.Error()))
// 			}
// 		}

// 	} else if commandAccountPositionDetail.MatchString(command) {

// 		result := commandAccountPositionDetail.FindAllStringSubmatch(command, -1)

// 		accountId := result[0][1]
// 		ticker := result[0][2]

// 		err = HandleAccountPosition(user, accountId, ticker)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle account position: %s", err.Error()))
// 		}

