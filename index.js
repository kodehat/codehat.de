var minifier = require('minifier')

var appJs = 'src/js/app.js'
var appCss = 'src/css/app.css'

minifier.minify(appJs)//, { output: 'public/js/app.min.js'})
minifier.minify(appCss)//, { output: 'public/css/app.min.css'})
