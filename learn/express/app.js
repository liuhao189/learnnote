const express = require('express');
const path = require('path');
const app = express();

app.get('/', (req, res) => res.send('Hello World!'));

app.post('/', (req, res) => { res.send('Got a post request!') });

app.put('/user', (req, res) => res.send('Got a put request at /user!'));

app.delete('/users/1', (req, res) => res.send('Got a delete request at /user!'));

const staticBasePath = path.join(__dirname, 'statics')
console.log(staticBasePath);
app.use('/store', express.static(staticBasePath + '\\imgs'));
app.use('/store', express.static(staticBasePath + '\\scripts'));

app.listen(3000, () => { console.log('Example app listening on port 3000!') });