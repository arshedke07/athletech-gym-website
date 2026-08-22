const jwt = require('jsonwebtoken')

const verify = (req, res, next) => {
    try {
        const token = req.cookies.jwt
        // console.log(req.cookies)

        if (!token) {
            return res.status(401).json({ error: 'No auth token provided' });
        }

        const decoded = jwt.verify(token, process.env.JWT_SECRET)
        req.user = decoded
        next()
    } catch (error) {
        console.error('JWT verification failed:', error.message);
        return res.status(401).json({ error: 'Unauthorized' });
    }
}

module.exports = verify