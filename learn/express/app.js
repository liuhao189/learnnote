const express = require('express');
const path = require('path');
const app = express();
const birds = require('./birdRoute');
const myMiddleWareFactory = require('./configMiddleWare');

const myLogger = function (req, res, next) {
    console.log('Logged!');
    next();
}

const requestTime = (req, res, next) => {
    req.requestTime = Date.now();
    next();
}
app.use(myMiddleWareFactory({ name: 'learn', age: 99 }));
app.use(requestTime);
app.use('/birds', birds);
app.use(myLogger);

app.get('/reporttime', (req, res) => {
    let responseTxt = "Hello World!<br/>";
    responseTxt += `<samll>Requested at ${req.requestTime} !</small>`;
    res.send(responseTxt);
});

app.get('/reportoption', (req, res) => {
    let responseTxt = "Hello World!<br/>";
    responseTxt += `<samll>Requested options is  ${JSON.stringify(req.options)} !</small>`;
    res.send(responseTxt);
});

app.route('/book').get((req, res) => {
    res.send('Get a random book!')
}).post((req, res) => {
    res.send('Add a book!');
}).put((req, res) => {
    res.send('Update the book!')
});
//
app.get('/example/b', function (req, res, next) {
    console.log('The response will be sent by the next function...');
    setTimeout(() => {
        next();
    }, 100);
}, (req, res) => {
    res.send(new Date().valueOf() + ':Hello from B!')
});

const cb0 = function (req, res, next) {
    console.log('cb0');
    next();
}

const cb1 = function (req, res, next) {
    console.log('cb1');
    next();
}

const cb2 = function (req, res) {
    console.log('cb2');
    res.send('Hello from C!')
}
app.get('/example/c', [cb0, cb1, cb2]);

app.get('/example/d', [cb0, cb1], (req, res, next) => {
    console.log('the response will be sent by the next function...')
    next();
}, (req, res) => {
    res.send('Hello from D!')
})

app.get('/', (req, res) => res.send('Hello World!'));

app.post('/', (req, res) => { res.send('Got a post request!') });

app.put('/user', (req, res) => res.send('Got a put request at /user!'));

app.delete('/users/1', (req, res) => res.send('Got a delete request at /user!'));

app.all('/allmethod', (req, res) => {
    res.send(`Got a ${req.method} request at /all!`);
})

app.get('/ab?cd', (req, res) => {
    res.send('ab?cd')
});

app.get('/ab+cd', (req, res) => {
    res.send('ab+cd')
});

app.get('/ab*cd', (req, res) => {
    res.send('ab*cd');
});

app.get('/ab(cd)?e', (req, res) => {
    res.send('ab(cd)?e');
});

// app.get(/a/, (req, res) => {
//     res.send('/a/')
// })

app.get(/.*fly$/, (req, res) => {
    res.send('/.*fly$/');
});

app.get('/user/:userId(\\d+)', (req, res) => {
    res.send(req.params)
});

app.get('/users/:userId/books/:bookId', (req, res) => {
    res.send(req.params);
});

app.get('/flights/:form-:to', (req, res) => {
    res.send(req.params);
});

app.get('/plantae/:genus.:species', (req, res) => {
    res.send(req.params);
});

const staticBasePath = path.join(__dirname, 'statics');
app.use('/store', express.static(staticBasePath + '\\imgs'));
app.use('/store', express.static(staticBasePath + '\\scripts'));

app.listen(3000, () => { console.log('Example app listening on port 3000!') });

