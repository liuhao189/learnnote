const inquirer = require('inquirer');

inquirer.prompt([{
    type: 'input',
    name: 'likeNum',
    message: 'Which number do you like?',
    choices: ['One', 'Two', 'Three'],
    validate: function (val, answers) {
        if (val === 'three') {
            return true;
        } else {
            return 'Sorry!';
        }
    },
    filter: function (val, answers) {
        return val.toLowerCase();
    }
}, {
    type: 'expand',
    name: 'likeGirl',
    message: 'Which girl do you love?',
    when: function (answers) {
        if (answers.likeNum === 'three') {
            return false;
        }
    },
    choices: [{
        key: "1",
        name: 'girl1',
        value: 1,
        short: '1'
    },
    {
        key: "2",
        name: 'girl2',
        value: 2,
        short: '2'
    },
    {
        key: "3",
        name: 'girl3',
        value: 3,
        short: '3'
    }]
}, {
    type: 'checkbox',
    name: "wantEat",
    message: 'Which one do you want to eat?',
    choices: ['Baozi', 'Jiaozi', 'Miantiao', new inquirer.Separator(), 'Dabing'],
    pageSize: 6,
    prefix: 'Prefix',
    suffix: 'Suffix'
}]).then(answers => {
    console.log(`Your name is ${JSON.stringify(answers)}`);
}).catch(err => {
    console.error(err);
});