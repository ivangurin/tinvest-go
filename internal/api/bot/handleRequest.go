package bot

// func HandleRequest(update *tgbotapi.Update) {

// var message *tgbotapi.Message
// if update.Message != nil {
// 	message = update.Message
// } else if update.CallbackQuery != nil {
// 	message = update.CallbackQuery.Message
// 	message.From = update.CallbackQuery.From
// 	message.Text = update.CallbackQuery.Data
// 	defer sendCallback(update.CallbackQuery.ID)
// } else {
// 	return
// }

// user, err := updateUser(message)
// if err != nil {
// 	logger.Error(err)
// 	return
// }

// lock, exists := locks[message.From.ID]
// if !exists {
// 	lock = make(chan interface{}, 1)
// 	locks[message.From.ID] = lock
// }

// if len(lock) == 0 {
// 	lock <- true
// 	defer func() {
// 		<-lock
// 	}()
// } else {
// 	_, err := sendMessageWithText(message.Chat.ID, texts.StillProcessing)
// 	if err != nil {
// 		logger.Error(err)
// 	}
// 	return
// }

// Сохранение записи истории запросов
// err = repo.CreateHistory(user.Id, message.Text)
// if err != nil {
// 	logger.Error(errors.Errorf("error on create hostory record: %s", err.Error()))
// }

// var command string
// if isCommand(message.Text) {
// 	command = message.Text
// 	lastCommands[user.ID] = command
// } else {
// 	command = lastCommands[user.ID]
// }

// 	if command == commandStart {

// 		err = HandleStart(user)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle start command: %s", err.Error()))
// 		}

// 	} else if command == commandTest {

// 		err = handleTest(user)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle rsid command: %s", err.Error()))
// 		}

// 	} else if command == commandToken {

// 		if isCommand(message.Text) {
// 			err = HandleTokenRequest(user)
// 			if err != nil {
// 				logger.Error(errors.Errorf("error on handle token request: %s", err.Error()))
// 			}
// 		} else {
// 			err = handleTokenResponse(user, update.Message)
// 			if err != nil {
// 				logger.Error(errors.Errorf("error on handle token response: %s", err.Error()))
// 			}
// 		}

// 	} else if command == commandRsiD {

// 		err = HandleRSIDaily(user)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle rsid command: %s", err.Error()))
// 		}

// 	} else if command == commandAccounts {

// 		err = HandleAccounts(user)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle token command: %s", err.Error()))
// 		}

// 	} else if commandAccount.MatchString(command) {

// 		result := commandAccount.FindAllStringSubmatch(command, -1)

// 		accountId := result[0][1]

// 		err = HandleAccount(user, accountId)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle account: %s", err.Error()))
// 		}

// 	} else if commandAccountTotals.MatchString(command) {

// 		result := commandAccountTotals.FindAllStringSubmatch(command, -1)

// 		accountId := result[0][1]

// 		err = HandleAccountTotals(user, accountId)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle account totals: %s", err.Error()))
// 		}

// 	} else if commandAccountPortfolio.MatchString(command) {

// 		result := commandAccountPortfolio.FindAllStringSubmatch(command, -1)

// 		accountId := result[0][1]

// 		err = HandleAccountPortfolio(user, accountId)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle account portfolio: %s", err.Error()))
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

// 	} else if commandAccountTrades.MatchString(command) {

// 		result := commandAccountTrades.FindAllStringSubmatch(command, -1)

// 		accountId := result[0][1]

// 		err = HandleAccountTrades(user, accountId)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle account trades menu: %s", err.Error()))
// 		}

// 	} else if commandAccountTradesFor.MatchString(command) {

// 		result := commandAccountTradesFor.FindAllStringSubmatch(command, -1)

// 		accountId := result[0][1]
// 		period := result[0][2]

// 		err = HandleAccountTradesFor(user, accountId, period)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle account trades for %s: %s", period, err.Error()))
// 		}

// 	} else if commandAdmin.MatchString(command) {

// 		err = HandleAdmin(user)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle admin: %s", err.Error()))
// 		}

// 	} else if commandAdminUsers.MatchString(command) {

// 		err = HandleAdminUsers(user)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle admin users: %s", err.Error()))
// 		}

// 	} else {

// 		err = handleDefault(user)
// 		if err != nil {
// 			logger.Error(errors.Errorf("error on handle unknown command: %s", err.Error()))
// 		}

// 	}

// }

// func checkToken(user *model.User) (bool, error) {

// 	if user.Token == "" {

// 		message := tgbotapi.NewMessage(user.ChatID, texts.TokenIsEmpty)

// 		_, err = a.botClient.SendMessage(ctx, &message)
// 		if err != nil {
// 			return false, err
// 		}

// 		return false, err

// 	}

// 	return true, nil

// }
