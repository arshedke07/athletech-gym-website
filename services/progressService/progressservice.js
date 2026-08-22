const { pool } = require('./db')
const Progress = require('./model')

const GetProgressService = async (req, res) => {
    try {
        const userId = req.params.id
        const UserName = req.user.username
        console.log(req.user)

        const rows1 = await pool.query("SELECT current_weight, weight_goal, cardio_type, current_cardio, cardio_goal, lift_type, current_lift, lift_goal, current_body_fat, body_fat_goal FROM user_goals WHERE user_id = $1", [userId])
        const rows2 = await pool.query("SELECT current_weight, cardio_type, current_cardio, lift_type, current_lift, current_body_fat FROM user_profile_history WHERE user_id = $1 ORDER BY updated_at ASC LIMIT 1", [userId])

        let goalRow = rows1.rows[0] || {}
        let userHistory = rows2.rows[0] || {}

        let userGoals = new Progress({
            userId: userId,
            currentWeight: goalRow.current_weight || null,
            weightGoal: goalRow.weight_goal || null,
            cardioType: goalRow.cardio_type || null,
            currentCardio: goalRow.current_cardio || null,
            cardioGoal: goalRow.cardio_goal || null,
            liftType: goalRow.lift_type || null,
            currentLift: goalRow.current_lift || null,
            liftGoal: goalRow.lift_goal || null,
            currentBodyFat: goalRow.current_body_fat || null,
            bodyFatGoal: goalRow.body_fat_goal || null,
        });

        res.render('progresslog', {
            'Goals': userGoals,
            'History': userHistory,
            'UserName': UserName,
            'UserId': userId,
        })
    } catch (err) {
        console.error('Error in GetProgressService:', err);
        return res.status(500).json({ error: 'Internal server error' });
    }
}

const upsertProgressService = async (req, res) => {
    try {
        const userId = req.params.id
        const data = req.body // this is the body of the form submitted by the user on the progress log page
        console.log(userId)

        const result = await pool.query("INSERT INTO user_goals (user_id, current_weight, weight_goal, cardio_type, current_cardio, cardio_goal, lift_type, current_lift, lift_goal, current_body_fat, body_fat_goal, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) ON CONFLICT (user_id) DO UPDATE SET current_weight = COALESCE(EXCLUDED.current_weight, user_goals.current_weight), weight_goal = COALESCE(EXCLUDED.weight_goal, user_goals.weight_goal), cardio_type = COALESCE(nullIF(EXCLUDED.cardio_type, ''), user_goals.cardio_type), current_cardio = COALESCE(EXCLUDED.current_cardio, user_goals.current_cardio), cardio_goal = COALESCE(EXCLUDED.cardio_goal, user_goals.cardio_goal), lift_type = COALESCE(nullIF(EXCLUDED.lift_type, ''), user_goals.lift_type), current_lift = COALESCE(EXCLUDED.current_lift, user_goals.current_lift), lift_goal = COALESCE(EXCLUDED.lift_goal, user_goals.lift_goal), current_body_fat = COALESCE(EXCLUDED.current_body_fat, user_goals.current_body_fat), body_fat_goal = COALESCE(EXCLUDED.body_fat_goal, user_goals.body_fat_goal), updated_at = CURRENT_TIMESTAMP", [userId,
            data.current_weight,
            data.weight_goal,
            data.cardio_type,
            data.current_cardio,
            data.cardio_goal,
            data.lift_type,
            data.current_lift,
            data.lift_goal,
            data.current_body_fat,
            data.body_fat_goal]
        )

        return res.json(result)
    } catch (error) {
        console.error(error.message)
    }
}

module.exports = { GetProgressService, upsertProgressService }