package chain

const RetargetInterval = 3

const TargetBlockTimeSeconds = 2

const MaxDifficultyStep = 1

func CalculateNextDifficulty(bc *Blockchain) int {

	current := bc.GetLatestBlock().Difficulty

	// Not enough blocks yet to measure an interval — keep current difficulty.
	if len(bc.Blocks) <= RetargetInterval {
		return current
	}

	// Only retarget every RetargetInterval blocks.
	if (len(bc.Blocks)-1)%RetargetInterval != 0 {
		return current
	}

	newest := bc.Blocks[len(bc.Blocks)-1]
	oldest := bc.Blocks[len(bc.Blocks)-1-RetargetInterval]

	actualSeconds := newest.Timestamp - oldest.Timestamp
	targetSeconds := int64(RetargetInterval * TargetBlockTimeSeconds)

	next := current

	if actualSeconds < targetSeconds {
		// Mined faster than target: increase difficulty.
		next = current + MaxDifficultyStep
	} else if actualSeconds > targetSeconds {
		// Mined slower than target: decrease difficulty.
		next = current - MaxDifficultyStep
	}

	if next < MinDifficulty {
		next = MinDifficulty
	}

	return next
}


func NextDifficultyFor(bc *Blockchain, requestedDifficulty int) int {

	if len(bc.Blocks) > RetargetInterval {
		return CalculateNextDifficulty(bc)
	}

	return requestedDifficulty
}