
//
function sendRequest(url, method) {
    return new Promise((resolve, reject) => {
        fetch(url, { method: method }).then(res => {
            res.text().then(str => {
                resolve(str);
            }).catch(err => {
                reject(err);
            }
            ).catch(err => {
                reject(err);
            })
        });
    });
}

function testRoutes(url, method) {
    sendRequest(url, method).then(str => {
        alert(str);
    });
}

['get', 'post', 'put', 'delete', 'header', 'options'].forEach((metod) => {
    testRoutes('/allmethod',metod);
});
testRoutes('/', 'post');
testRoutes('/user', 'put');
testRoutes('/users/1', 'delete');