const env = require('dotenv').config();
const express = require('express')
const app = express()
const { GetProgressService, upsertProgressService } = require('./progressservice');
const cookieParser = require('cookie-parser');
const verify = require('../../utils/authMiddleware')
const exphbs = require('express-handlebars');
const cors = require('cors')

// Handlebars setup — no layouts
app.engine('hbs', exphbs.engine({
    extname: 'hbs',
    defaultLayout: false  // 👈 disable layout system completely
}));
// set handlebars as rendering engine
app.set('view engine', 'hbs');
app.set('views', './templates');
//converted my progresslog.html to .hbs as a template rendering engine to use variables inside the .hbs file

app.use(cors())
app.use(express.json())
app.use(cookieParser())

app.get('/progressService/:id', verify, GetProgressService)
app.post('/progressService/:id', upsertProgressService)

app.listen(3001, () => {
    console.log("service running on port 3001...")
})