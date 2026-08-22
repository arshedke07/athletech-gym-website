class Progress {
    constructor({
        userId,
        currentWeight,
        weightGoal,
        cardioType,
        currentCardio,
        cardioGoal,
        liftType,
        currentLift,
        liftGoal,
        currentBodyFat,
        bodyFatGoal,
    }) {
        this.userId = userId;
        this.currentWeight = currentWeight;
        this.weightGoal = weightGoal;
        this.cardioType = cardioType;
        this.currentCardio = currentCardio;
        this.cardioGoal = cardioGoal;
        this.liftType = liftType;
        this.currentLift = currentLift;
        this.liftGoal = liftGoal;
        this.currentBodyFat = currentBodyFat;
        this.bodyFatGoal = bodyFatGoal;
    }
}

module.exports = Progress; // ✅ CommonJS export