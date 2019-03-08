const inquirer = require('inquirer');

const types = ['input', 'confirm', 'list', 'rawlist', 'expand', 'checkbox', 'password', 'editor']

inquirer.prompt(types.map((item) => {
    return {
        type: item,
        name: 'name' + item,
        choices: item === 'expand' ? [
            {
                key: 'a',
                name: 'aaa',
                value: 1
            },
            {
                key: 'b',
                name: 'bbb',
                value: 2
            }
        ] : ['a', 'b'],
        message: item + ":What is your name?",

    }
})).then(answer => {
    console.log('Your name is ' + JSON.stringify(answer) + '!');
}).catch(err => {
    console.error(err);
});