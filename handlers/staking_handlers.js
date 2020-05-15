const {setupApiAtHeight} = require('../utils/setup');
const stakingMappers = require('../mappers/staking/staking_mappers');

/**
 * Get staking information by height
 */
const getByHeight = (api) => async (call, callback) => {
  const height = call.request.height;

  const {blockHash} = await setupApiAtHeight(api, height);

  const timestamp = await api.query.timestamp.now.at(blockHash);
  console.log('timestamp', timestamp.toNumber());

  const sessionAt = await api.query.session.currentIndex.at(blockHash);
  console.log('current session #: ', sessionAt.toString());

  const eraAt = await api.query.staking.currentEra.at(blockHash);
  console.log('currentEra: ', eraAt.toString());

  // ERA QUERIES
  const erasRewardPoints = await api.query.staking.erasRewardPoints(eraAt.toString());
  console.log('erasRewardPoints: ', erasRewardPoints.toString());

  const erasTotalStake = await api.query.staking.erasTotalStake(eraAt.toString());
  console.log('erasTotalStake: ', erasTotalStake.toString());

  // Validator reward for era
  // The total validator era payout for the last HISTORY_DEPTH eras.
  // Eras that haven't finished yet or has been removed doesn't have reward.
  const erasValidatorReward = await api.query.staking.erasValidatorReward(eraAt.toString());
  console.log('erasValidatorReward: ', erasValidatorReward.toString());

  const validatorsAt = await api.query.session.validators.at(blockHash);
  const validatorsData = [];
  for (const rawValidator of validatorsAt) {
    const validatorControllerAccount = rawValidator.toString();
    const validator = {
      controllerAccount: validatorControllerAccount,
      rewardPoints: (erasRewardPoints.individual.toJSON()[validatorControllerAccount] || '0').toString(),
      stakers: [],
    };

    // Get validator stash account
    const validatorStashAccount = await api.query.staking.bonded.at(blockHash, validatorControllerAccount);
    if (!validatorStashAccount.isEmpty) {
      validator.stashAccount = validatorStashAccount.toString();
    }

    // Get stakers for validator
    const erasStakers = await api.query.staking.erasStakers(eraAt.toString(), validatorControllerAccount);
    validator.totalStake = erasStakers.total.toString();
    validator.ownStake = erasStakers.own.toString();

    for (const stake of erasStakers.others) {
      const nominatorControllerAccount = stake.who.toString();

      // Get nominator stash account
      const nominatorStashAccount = await api.query.staking.bonded.at(blockHash, nominatorControllerAccount);

      validator.stakers.push({
        stashAccount: nominatorStashAccount.toString(),
        controllerAccount: nominatorControllerAccount,
        stake: stake.value.toString(),
      })
    }

    // Get Validator prefs (commission)
    const erasValidatorPrefs = await api.query.staking.erasValidatorPrefs(eraAt.toString(), validatorControllerAccount);
    validator.commission = erasValidatorPrefs.commission.toString();

    validatorsData.push(validator);
    break;
  }

  const response = {
    staking: {
      session: sessionAt.toString(),
      era: eraAt.toString(),
      totalStake: erasTotalStake.toString(),
      totalRewardPayout: (erasValidatorReward.isEmpty ? '0' : erasValidatorReward).toString(),
      totalRewardPoints: erasRewardPoints.total.toString(),
      validators: validatorsData,
    },
  };

  console.log("Response", response.staking.validators);

  callback(null, response);
};

module.exports = {
  getByHeight,
};
